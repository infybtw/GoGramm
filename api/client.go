// Package api implements a typed client for the Telegram Bot API.
//
// Every Bot API method is exposed as a method on Client; requests funnel
// through Invoke, the single network path.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

// Client is a Telegram Bot API client. It is safe for concurrent use.
type Client struct {
	token   string
	baseURL string
	http    *http.Client
}

// Option configures a Client returned by New.
type Option func(*Client)

// WithHTTPClient sets the underlying *http.Client. A nil client is ignored
// and the default (&http.Client{}, no timeout: pass a context to cancel
// calls) is kept.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) {
		if h != nil {
			c.http = h
		}
	}
}

// WithServerURL overrides the Bot API server base URL (default
// "https://api.telegram.org"); useful for testing against an httptest
// server or a local Bot API server.
func WithServerURL(url string) Option {
	return func(c *Client) {
		if url != "" {
			c.baseURL = url
		}
	}
}

// New returns a Client that authenticates with the given bot token.
func New(token string, opts ...Option) *Client {
	c := &Client{
		token:   token,
		baseURL: "https://api.telegram.org",
		http:    &http.Client{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// apiResponse is the envelope every Bot API reply is wrapped in.
type apiResponse struct {
	OK          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// Invoke performs a raw Bot API call: it POSTs params (JSON-serialized, or
// multipart/form-data when params contains InputFile uploads) to
// baseURL/bot<token>/<method> and decodes the "result" field into result.
//
// params may be nil (empty body). result may be nil (result discarded).
// A Telegram-side failure is returned as *Error.
func (c *Client) Invoke(ctx context.Context, method string, params, result any) error {
	url := strings.TrimSuffix(c.baseURL, "/") + "/bot" + c.token + "/" + method

	var body io.Reader
	contentType := "application/json"
	if params != nil {
		uploads := collectUploads(params)
		if len(uploads) > 0 {
			buf, ct, err := multipartBody(params, uploads)
			if err != nil {
				return fmt.Errorf("telegram: %s: %w", method, err)
			}
			body, contentType = buf, ct
		} else {
			b, err := json.Marshal(params)
			if err != nil {
				return fmt.Errorf("telegram: %s: %w", method, err)
			}
			body = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("telegram: %s: %w", method, err)
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("telegram: %s: %w", method, err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("telegram: %s: %w", method, err)
	}

	var env apiResponse
	if err := json.Unmarshal(raw, &env); err != nil {
		return fmt.Errorf("telegram: %s: decode response: %w", method, err)
	}

	if !env.OK || resp.StatusCode < 200 || resp.StatusCode > 299 {
		e := &Error{
			Code:        env.ErrorCode,
			Description: env.Description,
			Parameters:  env.Parameters,
		}
		if e.Code == 0 {
			e.Code = resp.StatusCode
		}
		return e
	}

	if result != nil {
		if handled, err := decodeUnionResult(result, env.Result); handled {
			if err != nil {
				return fmt.Errorf("telegram: decode %s result: %w", method, err)
			}
			return nil
		}
		if err := json.Unmarshal(env.Result, result); err != nil {
			return fmt.Errorf("telegram: decode %s result: %w", method, err)
		}
	}
	return nil
}

// Error is a failed Bot API response (ok == false or non-2xx HTTP status).
type Error struct {
	Code        int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("telegram: %d: %s", e.Code, e.Description)
}

// ResponseParameters contains information about why a request was
// unsuccessful and, sometimes, what to do about it.
type ResponseParameters struct {
	// MigrateToChatID is the target chat the bot should migrate to
	// (the group was upgraded to a supergroup).
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	// RetryAfter is the number of seconds after which the request can be
	// retried.
	RetryAfter int `json:"retry_after"`
}

// ChatID identifies a target chat in request parameters, either by its
// numeric id or by a @username. Construct it with NewChatID or NewUsername;
// Telegram validates the value server-side.
type ChatID struct {
	id       int64
	username string
}

// NewChatID returns a ChatID referencing a chat by its numeric identifier.
func NewChatID(id int64) ChatID { return ChatID{id: id} }

// NewUsername returns a ChatID referencing a channel or supergroup by its
// @username.
func NewUsername(name string) ChatID { return ChatID{username: name} }

// MarshalJSON encodes the chat id as a JSON number, or the username as a
// JSON string.
func (c ChatID) MarshalJSON() ([]byte, error) {
	if c.username != "" {
		return json.Marshal(c.username)
	}
	return json.Marshal(c.id)
}

type inputFileKind int

const (
	inputFileID inputFileKind = iota
	inputFileURL
	inputFileUpload
)

// InputFile is a file reference in request parameters: an existing Telegram
// file_id, an HTTPS URL for Telegram to fetch, or a new file to upload
// (multipart/form-data). Construct it with FileID, FileURL, or FileUpload.
//
// Upload files are consumed once: when a request carrying uploads is
// encoded, each upload's InputFile field is rewritten in place to the
// "attach://N" file_id form and its reader is spent. Reusing a params
// struct after such a call is a programming error.
type InputFile struct {
	kind inputFileKind
	val  string // file_id, URL, or upload filename
	r    io.Reader
}

// FileID returns an InputFile referencing a file already stored on the
// Telegram servers.
func FileID(id string) InputFile { return InputFile{kind: inputFileID, val: id} }

// FileURL returns an InputFile referencing a file for Telegram to fetch
// over HTTPS.
func FileURL(url string) InputFile { return InputFile{kind: inputFileURL, val: url} }

// FileUpload returns an InputFile uploading a new file (consumed once; the
// request is sent as multipart/form-data). name is the reported filename.
func FileUpload(name string, r io.Reader) InputFile {
	return InputFile{kind: inputFileUpload, val: name, r: r}
}

// MarshalJSON encodes file_id and URL references as their string value.
// Uploading an InputFile outside a multipart request is a programming error.
func (f InputFile) MarshalJSON() ([]byte, error) {
	switch f.kind {
	case inputFileID, inputFileURL:
		return json.Marshal(f.val)
	default:
		return nil, errors.New("telegram: FileUpload requires multipart request")
	}
}

// upload is one file part discovered while scanning request parameters.
type upload struct {
	partName string
	fileName string
	reader   io.Reader
}

var inputFileType = reflect.TypeOf(InputFile{})

// collectUploads walks params (for typed calls always a non-nil pointer to
// a params struct), finds every InputFile of upload kind, rewrites it in
// place to the "attach://N" file_id form, and returns the uploads in
// discovery order. The walk descends exported struct fields, pointers,
// interfaces, slices, and maps; InputFile is the leaf.
func collectUploads(params any) []*upload {
	var uploads []*upload
	walkForUploads(reflect.ValueOf(params), &uploads)
	return uploads
}

func walkForUploads(v reflect.Value, uploads *[]*upload) {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !v.IsNil() {
			walkForUploads(v.Elem(), uploads)
		}
	case reflect.Struct:
		if v.Type() == inputFileType {
			if !v.CanAddr() {
				return // unreachable (map value); uploads must be addressable to rewrite
			}
			f := v.Addr().Interface().(*InputFile)
			if f.kind == inputFileUpload {
				part := strconv.Itoa(len(*uploads))
				name, r := f.val, f.r
				*f = FileID("attach://" + part)
				*uploads = append(*uploads, &upload{partName: part, fileName: name, reader: r})
			}
			return
		}
		t := v.Type()
		for i := range t.NumField() {
			if t.Field(i).PkgPath == "" { // exported only
				walkForUploads(v.Field(i), uploads)
			}
		}
	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			walkForUploads(v.Index(i), uploads)
		}
	case reflect.Map:
		iter := v.MapRange()
		for iter.Next() {
			walkForUploads(iter.Value(), uploads)
		}
	}
}

// multipartBody encodes params as multipart/form-data: one form field per
// exported top-level params field (name = json tag, value = the field's
// JSON encoding, honoring omitempty), plus one file part per upload (field
// name = part name, filename = upload name).
func multipartBody(params any, uploads []*upload) (io.Reader, string, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	v := reflect.ValueOf(params)
	for v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)
			if sf.PkgPath != "" {
				continue
			}
			name, ok := jsonFieldName(sf)
			if !ok {
				continue
			}
			fv := v.Field(i)
			if hasOmitEmpty(sf) && isEmptyValue(fv) {
				continue
			}
			b, err := json.Marshal(fv.Interface())
			if err != nil {
				return nil, "", err
			}
			if err := w.WriteField(name, string(b)); err != nil {
				return nil, "", err
			}
		}
	}

	for _, u := range uploads {
		fw, err := w.CreateFormFile(u.partName, u.fileName)
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(fw, u.reader); err != nil {
			return nil, "", err
		}
	}
	if err := w.Close(); err != nil {
		return nil, "", err
	}
	return &buf, w.FormDataContentType(), nil
}

// jsonFieldName returns the wire name for a struct field per its json tag.
func jsonFieldName(sf reflect.StructField) (string, bool) {
	tag := sf.Tag.Get("json")
	if tag == "-" {
		return "", false
	}
	name, _, _ := strings.Cut(tag, ",")
	if name == "" {
		name = sf.Name
	}
	return name, true
}

func hasOmitEmpty(sf reflect.StructField) bool {
	_, opts, _ := strings.Cut(sf.Tag.Get("json"), ",")
	return strings.Contains(opts, "omitempty")
}

// isEmptyValue mirrors encoding/json omitempty semantics.
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		return v.IsNil()
	case reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	}
	return false
}

// WebAppInfo describes a Web App.
type WebAppInfo struct {
	// URL is an HTTPS URL of a Web App to open.
	URL string `json:"url"`
}

// Text parse modes for sendMessage-family caption/text fields.
const (
	ParseModeMarkdown   = "Markdown"
	ParseModeMarkdownV2 = "MarkdownV2"
	ParseModeHTML       = "HTML"
)

// Values for sendChatAction's action parameter.
const (
	ChatActionTyping          = "typing"
	ChatActionUploadPhoto     = "upload_photo"
	ChatActionRecordVideo     = "record_video"
	ChatActionUploadVideo     = "upload_video"
	ChatActionRecordVoice     = "record_voice"
	ChatActionUploadVoice     = "upload_voice"
	ChatActionUploadDocument  = "upload_document"
	ChatActionChooseSticker   = "choose_sticker"
	ChatActionFindLocation    = "find_location"
	ChatActionRecordVideoNote = "record_video_note"
	ChatActionUploadVideoNote = "upload_video_note"
)

// Emoji that can be passed to sendDice.
const (
	DiceEmojiDice        = "🎲"
	DiceEmojiDarts       = "🎯"
	DiceEmojiBasketball  = "🏀"
	DiceEmojiFootball    = "⚽"
	DiceEmojiBowling     = "🎳"
	DiceEmojiSlotMachine = "🎰"
)

func decodeChatMember(data []byte) (ChatMember, error) {
	var probe struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	var member ChatMember
	switch probe.Status {
	case "creator":
		member = &ChatMemberOwner{}
	case "administrator":
		member = &ChatMemberAdministrator{}
	case "member":
		member = &ChatMemberMember{}
	case "restricted":
		member = &ChatMemberRestricted{}
	case "left":
		member = &ChatMemberLeft{}
	case "kicked":
		member = &ChatMemberBanned{}
	default:
		return nil, fmt.Errorf("telegram: unknown ChatMember status %q", probe.Status)
	}
	if err := json.Unmarshal(data, member); err != nil {
		return nil, err
	}
	return member, nil
}

func decodeMenuButton(data []byte) (MenuButton, error) {
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	var button MenuButton
	switch probe.Type {
	case "commands":
		button = &MenuButtonCommands{}
	case "web_app":
		button = &MenuButtonWebApp{}
	case "default":
		button = &MenuButtonDefault{}
	default:
		return nil, fmt.Errorf("telegram: unknown MenuButton type %q", probe.Type)
	}
	if err := json.Unmarshal(data, button); err != nil {
		return nil, err
	}
	return button, nil
}

func decodeUnionResult(result any, data []byte) (bool, error) {
	switch target := result.(type) {
	case *ChatMember:
		member, err := decodeChatMember(data)
		if err != nil {
			return true, err
		}
		*target = member
		return true, nil
	case *MenuButton:
		button, err := decodeMenuButton(data)
		if err != nil {
			return true, err
		}
		*target = button
		return true, nil
	case *[]ChatMember:
		var raw []json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			return true, err
		}
		members := make([]ChatMember, len(raw))
		for i := range raw {
			member, err := decodeChatMember(raw[i])
			if err != nil {
				return true, err
			}
			members[i] = member
		}
		*target = members
		return true, nil
	default:
		return false, nil
	}
}

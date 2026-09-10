package api

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// RichMessage represents a rich API object.
type RichMessage struct {
	Blocks []RichBlock `json:"blocks"`
	IsRTL  *bool       `json:"is_rtl,omitempty"`
}

func (m *RichMessage) UnmarshalJSON(data []byte) error {
	var raw struct {
		Blocks []json.RawMessage `json:"blocks"`
		IsRTL  *bool             `json:"is_rtl,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	blocks, err := unmarshalRichBlocks(raw.Blocks)
	if err != nil {
		return err
	}
	m.Blocks, m.IsRTL = blocks, raw.IsRTL
	return nil
}

// InputRichMessage represents a rich API object.
type InputRichMessage struct {
	Blocks              []InputRichBlock        `json:"blocks,omitempty"`
	HTML                *string                 `json:"html,omitempty"`
	Markdown            *string                 `json:"markdown,omitempty"`
	Media               []InputRichMessageMedia `json:"media,omitempty"`
	IsRTL               *bool                   `json:"is_rtl,omitempty"`
	SkipEntityDetection *bool                   `json:"skip_entity_detection,omitempty"`
}

// InputRichMessageMedia represents a rich API object.
type InputRichMessageMedia struct {
	ID    string     `json:"id"`
	Media InputMedia `json:"media"`
}

// RichMessageButton represents a rich API object.
type RichMessageButton struct {
	Text                         RichText                     `json:"text"`
	Style                        *string                      `json:"style,omitempty"`
	URL                          *string                      `json:"url,omitempty"`
	CallbackData                 *string                      `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	Disabled                     *DisabledButton              `json:"disabled,omitempty"`
}

// RichText is either plain text, an array of rich text nodes, or a typed rich text node.
type RichText interface{ richText() }

type RichTextPlain string

func (*RichTextPlain) richText()                     {}
func (v RichTextPlain) MarshalJSON() ([]byte, error) { return json.Marshal(string(v)) }

type RichTextArray []RichText

func (*RichTextArray) richText()                     {}
func (v RichTextArray) MarshalJSON() ([]byte, error) { return json.Marshal([]RichText(v)) }

// RichTextBold represents a rich API object.
type RichTextBold struct {
	Text RichText `json:"text"`
}

func (*RichTextBold) richText() {}

func (v RichTextBold) MarshalJSON() ([]byte, error) {
	type alias RichTextBold
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "bold"})
}

// RichTextItalic represents a rich API object.
type RichTextItalic struct {
	Text RichText `json:"text"`
}

func (*RichTextItalic) richText() {}

func (v RichTextItalic) MarshalJSON() ([]byte, error) {
	type alias RichTextItalic
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "italic"})
}

// RichTextUnderline represents a rich API object.
type RichTextUnderline struct {
	Text RichText `json:"text"`
}

func (*RichTextUnderline) richText() {}

func (v RichTextUnderline) MarshalJSON() ([]byte, error) {
	type alias RichTextUnderline
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "underline"})
}

// RichTextStrikethrough represents a rich API object.
type RichTextStrikethrough struct {
	Text RichText `json:"text"`
}

func (*RichTextStrikethrough) richText() {}

func (v RichTextStrikethrough) MarshalJSON() ([]byte, error) {
	type alias RichTextStrikethrough
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "strikethrough"})
}

// RichTextSpoiler represents a rich API object.
type RichTextSpoiler struct {
	Text RichText `json:"text"`
}

func (*RichTextSpoiler) richText() {}

func (v RichTextSpoiler) MarshalJSON() ([]byte, error) {
	type alias RichTextSpoiler
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "spoiler"})
}

// RichTextDateTime represents a rich API object.
type RichTextDateTime struct {
	Text           RichText `json:"text"`
	UnixTime       int64    `json:"unix_time"`
	DateTimeFormat string   `json:"date_time_format"`
}

func (*RichTextDateTime) richText() {}

func (v RichTextDateTime) MarshalJSON() ([]byte, error) {
	type alias RichTextDateTime
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "date_time"})
}

// RichTextTextMention represents a rich API object.
type RichTextTextMention struct {
	Text RichText `json:"text"`
	User User     `json:"user"`
}

func (*RichTextTextMention) richText() {}

func (v RichTextTextMention) MarshalJSON() ([]byte, error) {
	type alias RichTextTextMention
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "text_mention"})
}

// RichTextSubscript represents a rich API object.
type RichTextSubscript struct {
	Text RichText `json:"text"`
}

func (*RichTextSubscript) richText() {}

func (v RichTextSubscript) MarshalJSON() ([]byte, error) {
	type alias RichTextSubscript
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "subscript"})
}

// RichTextSuperscript represents a rich API object.
type RichTextSuperscript struct {
	Text RichText `json:"text"`
}

func (*RichTextSuperscript) richText() {}

func (v RichTextSuperscript) MarshalJSON() ([]byte, error) {
	type alias RichTextSuperscript
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "superscript"})
}

// RichTextMarked represents a rich API object.
type RichTextMarked struct {
	Text RichText `json:"text"`
}

func (*RichTextMarked) richText() {}

func (v RichTextMarked) MarshalJSON() ([]byte, error) {
	type alias RichTextMarked
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "marked"})
}

// RichTextCode represents a rich API object.
type RichTextCode struct {
	Text RichText `json:"text"`
}

func (*RichTextCode) richText() {}

func (v RichTextCode) MarshalJSON() ([]byte, error) {
	type alias RichTextCode
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "code"})
}

// RichTextCustomEmoji represents a rich API object.
type RichTextCustomEmoji struct {
	CustomEmojiID   string `json:"custom_emoji_id"`
	AlternativeText string `json:"alternative_text"`
}

func (*RichTextCustomEmoji) richText() {}

func (v RichTextCustomEmoji) MarshalJSON() ([]byte, error) {
	type alias RichTextCustomEmoji
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "custom_emoji"})
}

// RichTextMathematicalExpression represents a rich API object.
type RichTextMathematicalExpression struct {
	Expression string `json:"expression"`
}

func (*RichTextMathematicalExpression) richText() {}

func (v RichTextMathematicalExpression) MarshalJSON() ([]byte, error) {
	type alias RichTextMathematicalExpression
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "mathematical_expression"})
}

// RichTextUrl represents a rich API object.
type RichTextUrl struct {
	Text RichText `json:"text"`
	URL  string   `json:"url"`
}

func (*RichTextUrl) richText() {}

func (v RichTextUrl) MarshalJSON() ([]byte, error) {
	type alias RichTextUrl
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "url"})
}

// RichTextEmailAddress represents a rich API object.
type RichTextEmailAddress struct {
	Text         RichText `json:"text"`
	EmailAddress string   `json:"email_address"`
}

func (*RichTextEmailAddress) richText() {}

func (v RichTextEmailAddress) MarshalJSON() ([]byte, error) {
	type alias RichTextEmailAddress
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "email_address"})
}

// RichTextPhoneNumber represents a rich API object.
type RichTextPhoneNumber struct {
	Text        RichText `json:"text"`
	PhoneNumber string   `json:"phone_number"`
}

func (*RichTextPhoneNumber) richText() {}

func (v RichTextPhoneNumber) MarshalJSON() ([]byte, error) {
	type alias RichTextPhoneNumber
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "phone_number"})
}

// RichTextBankCardNumber represents a rich API object.
type RichTextBankCardNumber struct {
	Text           RichText `json:"text"`
	BankCardNumber string   `json:"bank_card_number"`
}

func (*RichTextBankCardNumber) richText() {}

func (v RichTextBankCardNumber) MarshalJSON() ([]byte, error) {
	type alias RichTextBankCardNumber
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "bank_card_number"})
}

// RichTextMention represents a rich API object.
type RichTextMention struct {
	Text     RichText `json:"text"`
	Username string   `json:"username"`
}

func (*RichTextMention) richText() {}

func (v RichTextMention) MarshalJSON() ([]byte, error) {
	type alias RichTextMention
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "mention"})
}

// RichTextHashtag represents a rich API object.
type RichTextHashtag struct {
	Text    RichText `json:"text"`
	Hashtag string   `json:"hashtag"`
}

func (*RichTextHashtag) richText() {}

func (v RichTextHashtag) MarshalJSON() ([]byte, error) {
	type alias RichTextHashtag
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "hashtag"})
}

// RichTextCashtag represents a rich API object.
type RichTextCashtag struct {
	Text    RichText `json:"text"`
	Cashtag string   `json:"cashtag"`
}

func (*RichTextCashtag) richText() {}

func (v RichTextCashtag) MarshalJSON() ([]byte, error) {
	type alias RichTextCashtag
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "cashtag"})
}

// RichTextBotCommand represents a rich API object.
type RichTextBotCommand struct {
	Text       RichText `json:"text"`
	BotCommand string   `json:"bot_command"`
}

func (*RichTextBotCommand) richText() {}

func (v RichTextBotCommand) MarshalJSON() ([]byte, error) {
	type alias RichTextBotCommand
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "bot_command"})
}

// RichTextButton represents a rich API object.
type RichTextButton struct {
	Button RichMessageButton `json:"button"`
}

func (*RichTextButton) richText() {}

func (v RichTextButton) MarshalJSON() ([]byte, error) {
	type alias RichTextButton
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "button"})
}

// RichTextAnchor represents a rich API object.
type RichTextAnchor struct {
	Name string `json:"name"`
}

func (*RichTextAnchor) richText() {}

func (v RichTextAnchor) MarshalJSON() ([]byte, error) {
	type alias RichTextAnchor
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "anchor"})
}

// RichTextAnchorLink represents a rich API object.
type RichTextAnchorLink struct {
	Text       RichText `json:"text"`
	AnchorName string   `json:"anchor_name"`
}

func (*RichTextAnchorLink) richText() {}

func (v RichTextAnchorLink) MarshalJSON() ([]byte, error) {
	type alias RichTextAnchorLink
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "anchor_link"})
}

// RichTextReference represents a rich API object.
type RichTextReference struct {
	Text RichText `json:"text"`
	Name string   `json:"name"`
}

func (*RichTextReference) richText() {}

func (v RichTextReference) MarshalJSON() ([]byte, error) {
	type alias RichTextReference
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "reference"})
}

// RichTextReferenceLink represents a rich API object.
type RichTextReferenceLink struct {
	Text          RichText `json:"text"`
	ReferenceName string   `json:"reference_name"`
}

func (*RichTextReferenceLink) richText() {}

func (v RichTextReferenceLink) MarshalJSON() ([]byte, error) {
	type alias RichTextReferenceLink
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "reference_link"})
}

func unmarshalRichText(data []byte) (RichText, error) {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return nil, fmt.Errorf("telegram: empty RichText")
	}
	switch data[0] {
	case '"':
		var text string
		if err := json.Unmarshal(data, &text); err != nil {
			return nil, err
		}
		v := RichTextPlain(text)
		return &v, nil
	case '[':
		var raw []json.RawMessage
		if err := json.Unmarshal(data, &raw); err != nil {
			return nil, err
		}
		values := make(RichTextArray, len(raw))
		for i := range raw {
			v, err := unmarshalRichText(raw[i])
			if err != nil {
				return nil, err
			}
			values[i] = v
		}
		return &values, nil
	case '{':
		var probe struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(data, &probe); err != nil {
			return nil, err
		}
		switch probe.Type {
		case "bold":
			var v RichTextBold
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "italic":
			var v RichTextItalic
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "underline":
			var v RichTextUnderline
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "strikethrough":
			var v RichTextStrikethrough
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "spoiler":
			var v RichTextSpoiler
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "date_time":
			var v RichTextDateTime
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "text_mention":
			var v RichTextTextMention
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "subscript":
			var v RichTextSubscript
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "superscript":
			var v RichTextSuperscript
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "marked":
			var v RichTextMarked
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "code":
			var v RichTextCode
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "custom_emoji":
			var v RichTextCustomEmoji
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "mathematical_expression":
			var v RichTextMathematicalExpression
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "url":
			var v RichTextUrl
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "email_address":
			var v RichTextEmailAddress
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "phone_number":
			var v RichTextPhoneNumber
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "bank_card_number":
			var v RichTextBankCardNumber
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "mention":
			var v RichTextMention
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "hashtag":
			var v RichTextHashtag
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "cashtag":
			var v RichTextCashtag
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "bot_command":
			var v RichTextBotCommand
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "button":
			var v RichTextButton
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "anchor":
			var v RichTextAnchor
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "anchor_link":
			var v RichTextAnchorLink
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "reference":
			var v RichTextReference
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		case "reference_link":
			var v RichTextReferenceLink
			if err := json.Unmarshal(data, &v); err != nil {
				return nil, err
			}
			return &v, nil
		default:
			return nil, fmt.Errorf("telegram: unknown RichText type %q", probe.Type)
		}
	default:
		return nil, fmt.Errorf("telegram: invalid RichText JSON")
	}
}

func (v *RichTextBold) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextBold{
		Text: Text,
	}
	return nil
}

func (v *RichTextItalic) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextItalic{
		Text: Text,
	}
	return nil
}

func (v *RichTextUnderline) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextUnderline{
		Text: Text,
	}
	return nil
}

func (v *RichTextStrikethrough) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextStrikethrough{
		Text: Text,
	}
	return nil
}

func (v *RichTextSpoiler) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextSpoiler{
		Text: Text,
	}
	return nil
}

func (v *RichTextDateTime) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text           json.RawMessage `json:"text"`
		UnixTime       int64           `json:"unix_time"`
		DateTimeFormat string          `json:"date_time_format"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextDateTime{
		Text:           Text,
		UnixTime:       raw.UnixTime,
		DateTimeFormat: raw.DateTimeFormat,
	}
	return nil
}

func (v *RichTextTextMention) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
		User User            `json:"user"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextTextMention{
		Text: Text,
		User: raw.User,
	}
	return nil
}

func (v *RichTextSubscript) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextSubscript{
		Text: Text,
	}
	return nil
}

func (v *RichTextSuperscript) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextSuperscript{
		Text: Text,
	}
	return nil
}

func (v *RichTextMarked) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextMarked{
		Text: Text,
	}
	return nil
}

func (v *RichTextCode) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextCode{
		Text: Text,
	}
	return nil
}

func (v *RichTextUrl) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
		URL  string          `json:"url"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextUrl{
		Text: Text,
		URL:  raw.URL,
	}
	return nil
}

func (v *RichTextEmailAddress) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text         json.RawMessage `json:"text"`
		EmailAddress string          `json:"email_address"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextEmailAddress{
		Text:         Text,
		EmailAddress: raw.EmailAddress,
	}
	return nil
}

func (v *RichTextPhoneNumber) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text        json.RawMessage `json:"text"`
		PhoneNumber string          `json:"phone_number"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextPhoneNumber{
		Text:        Text,
		PhoneNumber: raw.PhoneNumber,
	}
	return nil
}

func (v *RichTextBankCardNumber) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text           json.RawMessage `json:"text"`
		BankCardNumber string          `json:"bank_card_number"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextBankCardNumber{
		Text:           Text,
		BankCardNumber: raw.BankCardNumber,
	}
	return nil
}

func (v *RichTextMention) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text     json.RawMessage `json:"text"`
		Username string          `json:"username"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextMention{
		Text:     Text,
		Username: raw.Username,
	}
	return nil
}

func (v *RichTextHashtag) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text    json.RawMessage `json:"text"`
		Hashtag string          `json:"hashtag"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextHashtag{
		Text:    Text,
		Hashtag: raw.Hashtag,
	}
	return nil
}

func (v *RichTextCashtag) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text    json.RawMessage `json:"text"`
		Cashtag string          `json:"cashtag"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextCashtag{
		Text:    Text,
		Cashtag: raw.Cashtag,
	}
	return nil
}

func (v *RichTextBotCommand) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text       json.RawMessage `json:"text"`
		BotCommand string          `json:"bot_command"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextBotCommand{
		Text:       Text,
		BotCommand: raw.BotCommand,
	}
	return nil
}

func (v *RichTextAnchorLink) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text       json.RawMessage `json:"text"`
		AnchorName string          `json:"anchor_name"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextAnchorLink{
		Text:       Text,
		AnchorName: raw.AnchorName,
	}
	return nil
}

func (v *RichTextReference) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
		Name string          `json:"name"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextReference{
		Text: Text,
		Name: raw.Name,
	}
	return nil
}

func (v *RichTextReferenceLink) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text          json.RawMessage `json:"text"`
		ReferenceName string          `json:"reference_name"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichTextReferenceLink{
		Text:          Text,
		ReferenceName: raw.ReferenceName,
	}
	return nil
}

// RichBlockCaption represents a rich API object.
type RichBlockCaption struct {
	Text   RichText `json:"text"`
	Credit RichText `json:"credit,omitempty"`
}

// RichBlockTableCell represents a rich API object.
type RichBlockTableCell struct {
	Text     RichText `json:"text,omitempty"`
	IsHeader *bool    `json:"is_header,omitempty"`
	Colspan  *int     `json:"colspan,omitempty"`
	Rowspan  *int     `json:"rowspan,omitempty"`
	Align    string   `json:"align"`
	VAlign   string   `json:"valign"`
}

// RichBlockListItem represents a rich API object.
type RichBlockListItem struct {
	Label       string      `json:"label"`
	Blocks      []RichBlock `json:"blocks"`
	HasCheckbox *bool       `json:"has_checkbox,omitempty"`
	IsChecked   *bool       `json:"is_checked,omitempty"`
	Value       *int        `json:"value,omitempty"`
	Type        *string     `json:"type,omitempty"`
}

// RichBlock describes one of the typed rich message blocks.
type RichBlock interface{ richBlock() }

// RichBlockParagraph represents a rich API object.
type RichBlockParagraph struct {
	Text RichText `json:"text"`
}

func (*RichBlockParagraph) richBlock() {}

func (v RichBlockParagraph) MarshalJSON() ([]byte, error) {
	type alias RichBlockParagraph
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "paragraph"})
}

// RichBlockSectionHeading represents a rich API object.
type RichBlockSectionHeading struct {
	Text RichText `json:"text"`
	Size int      `json:"size"`
}

func (*RichBlockSectionHeading) richBlock() {}

func (v RichBlockSectionHeading) MarshalJSON() ([]byte, error) {
	type alias RichBlockSectionHeading
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "heading"})
}

// RichBlockPreformatted represents a rich API object.
type RichBlockPreformatted struct {
	Text     RichText `json:"text"`
	Language *string  `json:"language,omitempty"`
}

func (*RichBlockPreformatted) richBlock() {}

func (v RichBlockPreformatted) MarshalJSON() ([]byte, error) {
	type alias RichBlockPreformatted
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "pre"})
}

// RichBlockFooter represents a rich API object.
type RichBlockFooter struct {
	Text RichText `json:"text"`
}

func (*RichBlockFooter) richBlock() {}

func (v RichBlockFooter) MarshalJSON() ([]byte, error) {
	type alias RichBlockFooter
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "footer"})
}

// RichBlockDivider represents a rich API object.
type RichBlockDivider struct {
}

func (*RichBlockDivider) richBlock() {}

func (v RichBlockDivider) MarshalJSON() ([]byte, error) {
	type alias RichBlockDivider
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "divider"})
}

// RichBlockMathematicalExpression represents a rich API object.
type RichBlockMathematicalExpression struct {
	Expression string `json:"expression"`
}

func (*RichBlockMathematicalExpression) richBlock() {}

func (v RichBlockMathematicalExpression) MarshalJSON() ([]byte, error) {
	type alias RichBlockMathematicalExpression
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "mathematical_expression"})
}

// RichBlockAnchor represents a rich API object.
type RichBlockAnchor struct {
	Name string `json:"name"`
}

func (*RichBlockAnchor) richBlock() {}

func (v RichBlockAnchor) MarshalJSON() ([]byte, error) {
	type alias RichBlockAnchor
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "anchor"})
}

// RichBlockList represents a rich API object.
type RichBlockList struct {
	Items []RichBlockListItem `json:"items"`
}

func (*RichBlockList) richBlock() {}

func (v RichBlockList) MarshalJSON() ([]byte, error) {
	type alias RichBlockList
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "list"})
}

// RichBlockBlockQuotation represents a rich API object.
type RichBlockBlockQuotation struct {
	Blocks []RichBlock `json:"blocks"`
	Credit RichText    `json:"credit,omitempty"`
}

func (*RichBlockBlockQuotation) richBlock() {}

func (v RichBlockBlockQuotation) MarshalJSON() ([]byte, error) {
	type alias RichBlockBlockQuotation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "blockquote"})
}

// RichBlockExpandableBlockQuotation represents a rich API object.
type RichBlockExpandableBlockQuotation struct {
	Text   RichText `json:"text"`
	Credit RichText `json:"credit,omitempty"`
}

func (*RichBlockExpandableBlockQuotation) richBlock() {}

func (v RichBlockExpandableBlockQuotation) MarshalJSON() ([]byte, error) {
	type alias RichBlockExpandableBlockQuotation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "expandable_blockquote"})
}

// RichBlockPullQuotation represents a rich API object.
type RichBlockPullQuotation struct {
	Text   RichText `json:"text"`
	Credit RichText `json:"credit,omitempty"`
}

func (*RichBlockPullQuotation) richBlock() {}

func (v RichBlockPullQuotation) MarshalJSON() ([]byte, error) {
	type alias RichBlockPullQuotation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "pullquote"})
}

// RichBlockCollage represents a rich API object.
type RichBlockCollage struct {
	Blocks  []RichBlock       `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockCollage) richBlock() {}

func (v RichBlockCollage) MarshalJSON() ([]byte, error) {
	type alias RichBlockCollage
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "collage"})
}

// RichBlockSlideshow represents a rich API object.
type RichBlockSlideshow struct {
	Blocks  []RichBlock       `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockSlideshow) richBlock() {}

func (v RichBlockSlideshow) MarshalJSON() ([]byte, error) {
	type alias RichBlockSlideshow
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "slideshow"})
}

// RichBlockTable represents a rich API object.
type RichBlockTable struct {
	Cells      [][]RichBlockTableCell `json:"cells"`
	IsBordered *bool                  `json:"is_bordered,omitempty"`
	IsStriped  *bool                  `json:"is_striped,omitempty"`
	IsCompact  *bool                  `json:"is_compact,omitempty"`
	Caption    RichText               `json:"caption,omitempty"`
}

func (*RichBlockTable) richBlock() {}

func (v RichBlockTable) MarshalJSON() ([]byte, error) {
	type alias RichBlockTable
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "table"})
}

// RichBlockDetails represents a rich API object.
type RichBlockDetails struct {
	Summary RichText    `json:"summary"`
	Blocks  []RichBlock `json:"blocks"`
	IsOpen  *bool       `json:"is_open,omitempty"`
}

func (*RichBlockDetails) richBlock() {}

func (v RichBlockDetails) MarshalJSON() ([]byte, error) {
	type alias RichBlockDetails
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "details"})
}

// RichBlockMap represents a rich API object.
type RichBlockMap struct {
	Location Location          `json:"location"`
	Zoom     int               `json:"zoom"`
	Width    int               `json:"width"`
	Height   int               `json:"height"`
	Caption  *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockMap) richBlock() {}

func (v RichBlockMap) MarshalJSON() ([]byte, error) {
	type alias RichBlockMap
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "map"})
}

// RichBlockButtons represents a rich API object.
type RichBlockButtons struct {
	Buttons []RichMessageButton `json:"buttons"`
	Align   *string             `json:"align,omitempty"`
}

func (*RichBlockButtons) richBlock() {}

func (v RichBlockButtons) MarshalJSON() ([]byte, error) {
	type alias RichBlockButtons
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "buttons"})
}

// RichBlockAnimation represents a rich API object.
type RichBlockAnimation struct {
	Animation  Animation         `json:"animation"`
	HasSpoiler *bool             `json:"has_spoiler,omitempty"`
	Caption    *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockAnimation) richBlock() {}

func (v RichBlockAnimation) MarshalJSON() ([]byte, error) {
	type alias RichBlockAnimation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "animation"})
}

// RichBlockAudio represents a rich API object.
type RichBlockAudio struct {
	Audio   Audio             `json:"audio"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockAudio) richBlock() {}

func (v RichBlockAudio) MarshalJSON() ([]byte, error) {
	type alias RichBlockAudio
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "audio"})
}

// RichBlockDocument represents a rich API object.
type RichBlockDocument struct {
	Document Document          `json:"document"`
	Caption  *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockDocument) richBlock() {}

func (v RichBlockDocument) MarshalJSON() ([]byte, error) {
	type alias RichBlockDocument
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "document"})
}

// RichBlockPhoto represents a rich API object.
type RichBlockPhoto struct {
	Photo      []PhotoSize       `json:"photo"`
	HasSpoiler *bool             `json:"has_spoiler,omitempty"`
	Caption    *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockPhoto) richBlock() {}

func (v RichBlockPhoto) MarshalJSON() ([]byte, error) {
	type alias RichBlockPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// RichBlockVideo represents a rich API object.
type RichBlockVideo struct {
	Video      Video             `json:"video"`
	HasSpoiler *bool             `json:"has_spoiler,omitempty"`
	Caption    *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockVideo) richBlock() {}

func (v RichBlockVideo) MarshalJSON() ([]byte, error) {
	type alias RichBlockVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// RichBlockVoiceNote represents a rich API object.
type RichBlockVoiceNote struct {
	VoiceNote Voice             `json:"voice_note"`
	Caption   *RichBlockCaption `json:"caption,omitempty"`
}

func (*RichBlockVoiceNote) richBlock() {}

func (v RichBlockVoiceNote) MarshalJSON() ([]byte, error) {
	type alias RichBlockVoiceNote
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "voice_note"})
}

// RichBlockThinking represents a rich API object.
type RichBlockThinking struct {
	Text RichText `json:"text"`
}

func (*RichBlockThinking) richBlock() {}

func (v RichBlockThinking) MarshalJSON() ([]byte, error) {
	type alias RichBlockThinking
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "thinking"})
}

// InputRichBlock describes a typed block in an outgoing rich message.
type InputRichBlock interface{ inputRichBlock() }

// InputRichBlockListItem represents a rich API object.
type InputRichBlockListItem struct {
	Blocks      []InputRichBlock `json:"blocks"`
	HasCheckbox *bool            `json:"has_checkbox,omitempty"`
	IsChecked   *bool            `json:"is_checked,omitempty"`
	Value       *int             `json:"value,omitempty"`
	Type        *string          `json:"type,omitempty"`
}

// InputRichBlockParagraph represents a rich API object.
type InputRichBlockParagraph struct {
	Text RichText `json:"text"`
}

func (*InputRichBlockParagraph) inputRichBlock() {}

func (v InputRichBlockParagraph) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockParagraph
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "paragraph"})
}

// InputRichBlockSectionHeading represents a rich API object.
type InputRichBlockSectionHeading struct {
	Text RichText `json:"text"`
	Size int      `json:"size"`
}

func (*InputRichBlockSectionHeading) inputRichBlock() {}

func (v InputRichBlockSectionHeading) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockSectionHeading
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "heading"})
}

// InputRichBlockPreformatted represents a rich API object.
type InputRichBlockPreformatted struct {
	Text     RichText `json:"text"`
	Language *string  `json:"language,omitempty"`
}

func (*InputRichBlockPreformatted) inputRichBlock() {}

func (v InputRichBlockPreformatted) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockPreformatted
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "pre"})
}

// InputRichBlockFooter represents a rich API object.
type InputRichBlockFooter struct {
	Text RichText `json:"text"`
}

func (*InputRichBlockFooter) inputRichBlock() {}

func (v InputRichBlockFooter) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockFooter
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "footer"})
}

// InputRichBlockDivider represents a rich API object.
type InputRichBlockDivider struct {
}

func (*InputRichBlockDivider) inputRichBlock() {}

func (v InputRichBlockDivider) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockDivider
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "divider"})
}

// InputRichBlockMathematicalExpression represents a rich API object.
type InputRichBlockMathematicalExpression struct {
	Expression string `json:"expression"`
}

func (*InputRichBlockMathematicalExpression) inputRichBlock() {}

func (v InputRichBlockMathematicalExpression) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockMathematicalExpression
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "mathematical_expression"})
}

// InputRichBlockAnchor represents a rich API object.
type InputRichBlockAnchor struct {
	Name string `json:"name"`
}

func (*InputRichBlockAnchor) inputRichBlock() {}

func (v InputRichBlockAnchor) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockAnchor
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "anchor"})
}

// InputRichBlockList represents a rich API object.
type InputRichBlockList struct {
	Items []InputRichBlockListItem `json:"items"`
}

func (*InputRichBlockList) inputRichBlock() {}

func (v InputRichBlockList) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockList
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "list"})
}

// InputRichBlockBlockQuotation represents a rich API object.
type InputRichBlockBlockQuotation struct {
	Blocks []InputRichBlock `json:"blocks"`
	Credit RichText         `json:"credit,omitempty"`
}

func (*InputRichBlockBlockQuotation) inputRichBlock() {}

func (v InputRichBlockBlockQuotation) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockBlockQuotation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "blockquote"})
}

// InputRichBlockExpandableBlockQuotation represents a rich API object.
type InputRichBlockExpandableBlockQuotation struct {
	Text   RichText `json:"text"`
	Credit RichText `json:"credit,omitempty"`
}

func (*InputRichBlockExpandableBlockQuotation) inputRichBlock() {}

func (v InputRichBlockExpandableBlockQuotation) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockExpandableBlockQuotation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "expandable_blockquote"})
}

// InputRichBlockPullQuotation represents a rich API object.
type InputRichBlockPullQuotation struct {
	Text   RichText `json:"text"`
	Credit RichText `json:"credit,omitempty"`
}

func (*InputRichBlockPullQuotation) inputRichBlock() {}

func (v InputRichBlockPullQuotation) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockPullQuotation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "pullquote"})
}

// InputRichBlockCollage represents a rich API object.
type InputRichBlockCollage struct {
	Blocks  []InputRichBlock  `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*InputRichBlockCollage) inputRichBlock() {}

func (v InputRichBlockCollage) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockCollage
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "collage"})
}

// InputRichBlockSlideshow represents a rich API object.
type InputRichBlockSlideshow struct {
	Blocks  []InputRichBlock  `json:"blocks"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*InputRichBlockSlideshow) inputRichBlock() {}

func (v InputRichBlockSlideshow) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockSlideshow
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "slideshow"})
}

// InputRichBlockTable represents a rich API object.
type InputRichBlockTable struct {
	Cells      [][]RichBlockTableCell `json:"cells"`
	IsBordered *bool                  `json:"is_bordered,omitempty"`
	IsStriped  *bool                  `json:"is_striped,omitempty"`
	IsCompact  *bool                  `json:"is_compact,omitempty"`
	Caption    RichText               `json:"caption,omitempty"`
}

func (*InputRichBlockTable) inputRichBlock() {}

func (v InputRichBlockTable) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockTable
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "table"})
}

// InputRichBlockDetails represents a rich API object.
type InputRichBlockDetails struct {
	Summary RichText         `json:"summary"`
	Blocks  []InputRichBlock `json:"blocks"`
	IsOpen  *bool            `json:"is_open,omitempty"`
}

func (*InputRichBlockDetails) inputRichBlock() {}

func (v InputRichBlockDetails) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockDetails
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "details"})
}

// InputRichBlockMap represents a rich API object.
type InputRichBlockMap struct {
	Location Location          `json:"location"`
	Zoom     *int              `json:"zoom,omitempty"`
	Width    *int              `json:"width,omitempty"`
	Height   *int              `json:"height,omitempty"`
	Caption  *RichBlockCaption `json:"caption,omitempty"`
}

func (*InputRichBlockMap) inputRichBlock() {}

func (v InputRichBlockMap) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockMap
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "map"})
}

// InputRichBlockButtons represents a rich API object.
type InputRichBlockButtons struct {
	Buttons []RichMessageButton `json:"buttons"`
	Align   *string             `json:"align,omitempty"`
}

func (*InputRichBlockButtons) inputRichBlock() {}

func (v InputRichBlockButtons) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockButtons
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "buttons"})
}

// InputRichBlockAnimation represents a rich API object.
type InputRichBlockAnimation struct {
	Animation InputMediaAnimation `json:"animation"`
	Caption   *RichBlockCaption   `json:"caption,omitempty"`
}

func (*InputRichBlockAnimation) inputRichBlock() {}

func (v InputRichBlockAnimation) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockAnimation
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "animation"})
}

// InputRichBlockAudio represents a rich API object.
type InputRichBlockAudio struct {
	Audio   InputMediaAudio   `json:"audio"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*InputRichBlockAudio) inputRichBlock() {}

func (v InputRichBlockAudio) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockAudio
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "audio"})
}

// InputRichBlockDocument represents a rich API object.
type InputRichBlockDocument struct {
	Document InputMediaDocument `json:"document"`
	Caption  *RichBlockCaption  `json:"caption,omitempty"`
}

func (*InputRichBlockDocument) inputRichBlock() {}

func (v InputRichBlockDocument) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockDocument
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "document"})
}

// InputRichBlockPhoto represents a rich API object.
type InputRichBlockPhoto struct {
	Photo   InputMediaPhoto   `json:"photo"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*InputRichBlockPhoto) inputRichBlock() {}

func (v InputRichBlockPhoto) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockPhoto
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "photo"})
}

// InputRichBlockVideo represents a rich API object.
type InputRichBlockVideo struct {
	Video   InputMediaVideo   `json:"video"`
	Caption *RichBlockCaption `json:"caption,omitempty"`
}

func (*InputRichBlockVideo) inputRichBlock() {}

func (v InputRichBlockVideo) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockVideo
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "video"})
}

// InputRichBlockVoiceNote represents a rich API object.
type InputRichBlockVoiceNote struct {
	VoiceNote InputMediaVoiceNote `json:"voice_note"`
	Caption   *RichBlockCaption   `json:"caption,omitempty"`
}

func (*InputRichBlockVoiceNote) inputRichBlock() {}

func (v InputRichBlockVoiceNote) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockVoiceNote
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "voice_note"})
}

// InputRichBlockThinking represents a rich API object.
type InputRichBlockThinking struct {
	Text RichText `json:"text"`
}

func (*InputRichBlockThinking) inputRichBlock() {}

func (v InputRichBlockThinking) MarshalJSON() ([]byte, error) {
	type alias InputRichBlockThinking
	return json.Marshal(struct {
		alias
		Type string `json:"type"`
	}{alias(v), "thinking"})
}

func unmarshalRichBlocks(raw []json.RawMessage) ([]RichBlock, error) {
	blocks := make([]RichBlock, len(raw))
	for i := range raw {
		v, err := unmarshalRichBlock(raw[i])
		if err != nil {
			return nil, err
		}
		blocks[i] = v
	}
	return blocks, nil
}

func unmarshalRichBlock(data []byte) (RichBlock, error) {
	var probe struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return nil, err
	}
	switch probe.Type {
	case "paragraph":
		var v RichBlockParagraph
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "heading":
		var v RichBlockSectionHeading
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "pre":
		var v RichBlockPreformatted
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "footer":
		var v RichBlockFooter
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "divider":
		var v RichBlockDivider
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "mathematical_expression":
		var v RichBlockMathematicalExpression
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "anchor":
		var v RichBlockAnchor
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "list":
		var v RichBlockList
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "blockquote":
		var v RichBlockBlockQuotation
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "expandable_blockquote":
		var v RichBlockExpandableBlockQuotation
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "pullquote":
		var v RichBlockPullQuotation
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "collage":
		var v RichBlockCollage
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "slideshow":
		var v RichBlockSlideshow
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "table":
		var v RichBlockTable
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "details":
		var v RichBlockDetails
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "map":
		var v RichBlockMap
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "buttons":
		var v RichBlockButtons
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "animation":
		var v RichBlockAnimation
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "audio":
		var v RichBlockAudio
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "document":
		var v RichBlockDocument
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "photo":
		var v RichBlockPhoto
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "video":
		var v RichBlockVideo
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "voice_note":
		var v RichBlockVoiceNote
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	case "thinking":
		var v RichBlockThinking
		if err := json.Unmarshal(data, &v); err != nil {
			return nil, err
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("telegram: unknown RichBlock type %q", probe.Type)
	}
}

func (v *RichBlockParagraph) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichBlockParagraph{
		Text: Text,
	}
	return nil
}

func (v *RichBlockSectionHeading) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
		Size int             `json:"size"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichBlockSectionHeading{
		Text: Text,
		Size: raw.Size,
	}
	return nil
}

func (v *RichBlockPreformatted) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text     json.RawMessage `json:"text"`
		Language *string         `json:"language,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichBlockPreformatted{
		Text:     Text,
		Language: raw.Language,
	}
	return nil
}

func (v *RichBlockFooter) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichBlockFooter{
		Text: Text,
	}
	return nil
}

func (v *RichBlockBlockQuotation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Blocks []json.RawMessage `json:"blocks"`
		Credit json.RawMessage   `json:"credit,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Blocks, err := unmarshalRichBlocks(raw.Blocks)
	if err != nil {
		return err
	}
	var Credit RichText
	if len(raw.Credit) > 0 && string(raw.Credit) != "null" {
		Credit, err = unmarshalRichText(raw.Credit)
		if err != nil {
			return err
		}
	}
	*v = RichBlockBlockQuotation{
		Blocks: Blocks,
		Credit: Credit,
	}
	return nil
}

func (v *RichBlockExpandableBlockQuotation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text   json.RawMessage `json:"text"`
		Credit json.RawMessage `json:"credit,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	var Credit RichText
	if len(raw.Credit) > 0 && string(raw.Credit) != "null" {
		Credit, err = unmarshalRichText(raw.Credit)
		if err != nil {
			return err
		}
	}
	*v = RichBlockExpandableBlockQuotation{
		Text:   Text,
		Credit: Credit,
	}
	return nil
}

func (v *RichBlockPullQuotation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text   json.RawMessage `json:"text"`
		Credit json.RawMessage `json:"credit,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	var Credit RichText
	if len(raw.Credit) > 0 && string(raw.Credit) != "null" {
		Credit, err = unmarshalRichText(raw.Credit)
		if err != nil {
			return err
		}
	}
	*v = RichBlockPullQuotation{
		Text:   Text,
		Credit: Credit,
	}
	return nil
}

func (v *RichBlockCollage) UnmarshalJSON(data []byte) error {
	var raw struct {
		Blocks  []json.RawMessage `json:"blocks"`
		Caption json.RawMessage   `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Blocks, err := unmarshalRichBlocks(raw.Blocks)
	if err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockCollage{
		Blocks:  Blocks,
		Caption: Caption,
	}
	return nil
}

func (v *RichBlockSlideshow) UnmarshalJSON(data []byte) error {
	var raw struct {
		Blocks  []json.RawMessage `json:"blocks"`
		Caption json.RawMessage   `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Blocks, err := unmarshalRichBlocks(raw.Blocks)
	if err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockSlideshow{
		Blocks:  Blocks,
		Caption: Caption,
	}
	return nil
}

func (v *RichBlockTable) UnmarshalJSON(data []byte) error {
	var raw struct {
		Cells      [][]RichBlockTableCell `json:"cells"`
		IsBordered *bool                  `json:"is_bordered,omitempty"`
		IsStriped  *bool                  `json:"is_striped,omitempty"`
		IsCompact  *bool                  `json:"is_compact,omitempty"`
		Caption    json.RawMessage        `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption RichText
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		var err error
		Caption, err = unmarshalRichText(raw.Caption)
		if err != nil {
			return err
		}
	}
	*v = RichBlockTable{
		Cells:      raw.Cells,
		IsBordered: raw.IsBordered,
		IsStriped:  raw.IsStriped,
		IsCompact:  raw.IsCompact,
		Caption:    Caption,
	}
	return nil
}

func (v *RichBlockDetails) UnmarshalJSON(data []byte) error {
	var raw struct {
		Summary json.RawMessage   `json:"summary"`
		Blocks  []json.RawMessage `json:"blocks"`
		IsOpen  *bool             `json:"is_open,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Summary, err := unmarshalRichText(raw.Summary)
	if err != nil {
		return err
	}
	Blocks, err := unmarshalRichBlocks(raw.Blocks)
	if err != nil {
		return err
	}
	*v = RichBlockDetails{
		Summary: Summary,
		Blocks:  Blocks,
		IsOpen:  raw.IsOpen,
	}
	return nil
}

func (v *RichBlockMap) UnmarshalJSON(data []byte) error {
	var raw struct {
		Location Location        `json:"location"`
		Zoom     int             `json:"zoom"`
		Width    int             `json:"width"`
		Height   int             `json:"height"`
		Caption  json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockMap{
		Location: raw.Location,
		Zoom:     raw.Zoom,
		Width:    raw.Width,
		Height:   raw.Height,
		Caption:  Caption,
	}
	return nil
}

func (v *RichBlockAnimation) UnmarshalJSON(data []byte) error {
	var raw struct {
		Animation  Animation       `json:"animation"`
		HasSpoiler *bool           `json:"has_spoiler,omitempty"`
		Caption    json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockAnimation{
		Animation:  raw.Animation,
		HasSpoiler: raw.HasSpoiler,
		Caption:    Caption,
	}
	return nil
}

func (v *RichBlockAudio) UnmarshalJSON(data []byte) error {
	var raw struct {
		Audio   Audio           `json:"audio"`
		Caption json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockAudio{
		Audio:   raw.Audio,
		Caption: Caption,
	}
	return nil
}

func (v *RichBlockDocument) UnmarshalJSON(data []byte) error {
	var raw struct {
		Document Document        `json:"document"`
		Caption  json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockDocument{
		Document: raw.Document,
		Caption:  Caption,
	}
	return nil
}

func (v *RichBlockPhoto) UnmarshalJSON(data []byte) error {
	var raw struct {
		Photo      []PhotoSize     `json:"photo"`
		HasSpoiler *bool           `json:"has_spoiler,omitempty"`
		Caption    json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockPhoto{
		Photo:      raw.Photo,
		HasSpoiler: raw.HasSpoiler,
		Caption:    Caption,
	}
	return nil
}

func (v *RichBlockVideo) UnmarshalJSON(data []byte) error {
	var raw struct {
		Video      Video           `json:"video"`
		HasSpoiler *bool           `json:"has_spoiler,omitempty"`
		Caption    json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockVideo{
		Video:      raw.Video,
		HasSpoiler: raw.HasSpoiler,
		Caption:    Caption,
	}
	return nil
}

func (v *RichBlockVoiceNote) UnmarshalJSON(data []byte) error {
	var raw struct {
		VoiceNote Voice           `json:"voice_note"`
		Caption   json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	var Caption *RichBlockCaption
	if len(raw.Caption) > 0 && string(raw.Caption) != "null" {
		Caption = new(RichBlockCaption)
		if err := json.Unmarshal(raw.Caption, Caption); err != nil {
			return err
		}
	}
	*v = RichBlockVoiceNote{
		VoiceNote: raw.VoiceNote,
		Caption:   Caption,
	}
	return nil
}

func (v *RichBlockThinking) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text json.RawMessage `json:"text"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	Text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	*v = RichBlockThinking{
		Text: Text,
	}
	return nil
}

func (c *RichBlockCaption) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text   json.RawMessage `json:"text"`
		Credit json.RawMessage `json:"credit,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	c.Text = text
	c.Credit = nil
	if len(raw.Credit) > 0 && string(raw.Credit) != "null" {
		c.Credit, err = unmarshalRichText(raw.Credit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *RichBlockTableCell) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text     json.RawMessage `json:"text,omitempty"`
		IsHeader *bool           `json:"is_header,omitempty"`
		Colspan  *int            `json:"colspan,omitempty"`
		Rowspan  *int            `json:"rowspan,omitempty"`
		Align    string          `json:"align"`
		VAlign   string          `json:"valign"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.IsHeader, c.Colspan, c.Rowspan, c.Align, c.VAlign = raw.IsHeader, raw.Colspan, raw.Rowspan, raw.Align, raw.VAlign
	c.Text = nil
	if len(raw.Text) > 0 && string(raw.Text) != "null" {
		var err error
		c.Text, err = unmarshalRichText(raw.Text)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *RichBlockListItem) UnmarshalJSON(data []byte) error {
	var raw struct {
		Label       string            `json:"label"`
		Blocks      []json.RawMessage `json:"blocks"`
		HasCheckbox *bool             `json:"has_checkbox,omitempty"`
		IsChecked   *bool             `json:"is_checked,omitempty"`
		Value       *int              `json:"value,omitempty"`
		Type        *string           `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	blocks, err := unmarshalRichBlocks(raw.Blocks)
	if err != nil {
		return err
	}
	i.Label, i.Blocks, i.HasCheckbox, i.IsChecked, i.Value, i.Type = raw.Label, blocks, raw.HasCheckbox, raw.IsChecked, raw.Value, raw.Type
	return nil
}

func (b *RichMessageButton) UnmarshalJSON(data []byte) error {
	var raw struct {
		Text                         json.RawMessage              `json:"text"`
		Style                        *string                      `json:"style,omitempty"`
		URL                          *string                      `json:"url,omitempty"`
		CallbackData                 *string                      `json:"callback_data,omitempty"`
		WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
		LoginURL                     *LoginUrl                    `json:"login_url,omitempty"`
		SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
		SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
		SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
		CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
		Disabled                     *DisabledButton              `json:"disabled,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	text, err := unmarshalRichText(raw.Text)
	if err != nil {
		return err
	}
	b.Text, b.Style, b.URL, b.CallbackData, b.WebApp, b.LoginURL = text, raw.Style, raw.URL, raw.CallbackData, raw.WebApp, raw.LoginURL
	b.SwitchInlineQuery, b.SwitchInlineQueryCurrentChat, b.SwitchInlineQueryChosenChat = raw.SwitchInlineQuery, raw.SwitchInlineQueryCurrentChat, raw.SwitchInlineQueryChosenChat
	b.CopyText, b.Disabled = raw.CopyText, raw.Disabled
	return nil
}

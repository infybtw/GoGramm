# GoGram Agent Guide

## Verify

- Requires Go 1.27. Run the full suite with `go test ./...`; run a focused package or test with `go test ./api -run '^TestSendPhotoUpload$'` or `go test . -run '^TestStart'`.
- Format changed Go files with `gofmt -w <files>` before testing. There is no configured lint, CI, generator, or external test service.
- Tests use `httptest` and do not need a bot token or network access.

## Structure

- The root `gogram` package (`bot.go`) is the high-level long-polling dispatcher. `api/` is the typed Telegram Bot API client; all HTTP requests must go through `Client.Invoke`.
- Add Bot API methods as a `*MethodParams` struct with exact JSON tags and a thin `Client` method that delegates to `Invoke`, matching nearby methods in `api/methods_*.go`.
- `InputFile` uploads mutate the parameter struct to `attach://N` and consume their readers. Do not reuse params containing `FileUpload` after a call.

## Coupled Changes

- Adding an `api.Update` field also requires its `UpdateType` constant and `classify` branch in `bot.go`, plus the `updateTypes` test list in `bot_test.go`; that test deliberately enforces this synchronization.
- Tagged-union interfaces in `api/types_*.go` require discriminator-aware `UnmarshalJSON` implementations. Add decoding coverage in `api/types_decode_test.go` when introducing a new variant.

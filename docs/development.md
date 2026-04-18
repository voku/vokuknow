# Development

## Local workflow

```bash
go test ./...
go run ./cmd/vokuknow init
go run ./cmd/vokuknow build
go run ./cmd/vokuknow lint
```

## Determinism

- The generator sorts schema collections.
- File writes are `write-if-changed` for stable diffs.
- JSON output is pretty-printed with stable key ordering.

## Testing

Current tests cover:

- schema loading and invalid YAML errors
- deterministic generation
- golden markdown outputs
- CLI init/build/lint path


## CI

GitHub Actions runs these checks on pushes to `main` and on pull requests:

- `go mod tidy` drift check
- `gofmt` formatting check
- `go vet`
- `staticcheck`
- `go test ./...`
- `go test -race ./...`
- CLI smoke checks via `init`, `build`, `lint`, and `doctor` against a temporary repo

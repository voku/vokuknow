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
- generated-artifact drift check using `testdata/drift-fixture/`, `init`, `build`, and `git diff --exit-code`

To reproduce the drift check locally:

```bash
fixture_repo=$(mktemp -d)
cp -R testdata/drift-fixture/. "$fixture_repo"/
git -C "$fixture_repo" init
git -C "$fixture_repo" add .
go run ./cmd/vokuknow --repo "$fixture_repo" init
go run ./cmd/vokuknow --repo "$fixture_repo" build
git -C "$fixture_repo" add -N .
git -C "$fixture_repo" diff --exit-code
```

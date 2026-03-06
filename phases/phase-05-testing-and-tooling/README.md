# Phase 05: Testing and Tooling

## Goals
- Write maintainable unit tests
- Use benchmarks for performance checks
- Apply standard Go tooling

## Theory: Documentation Chapters
- Tutorial chapter "Add a test": https://go.dev/doc/tutorial/add-a-test
- Tutorial chapter "Fuzzing": https://go.dev/doc/tutorial/fuzz
- Package docs `testing`: https://pkg.go.dev/testing
- Tool docs `go test`, `go vet`, and `gofmt`: https://pkg.go.dev/cmd/go, https://pkg.go.dev/cmd/vet, https://pkg.go.dev/cmd/gofmt

## Checklist
- Add table-driven tests for all packages
- Add one benchmark per core package
- Run `go test ./...`, `go vet ./...`, `go fmt ./...`

## Extra Practice Projects
- Refactor a CLI app to dependency-injected design for easier testing
- Build a package with table-driven, fuzz, and benchmark tests
- Snapshot testing helper for stable text output
- Test report summarizer CLI that parses `go test` output

## Done When
- New features ship with tests by default

# Phase 03: Interfaces and Errors

## Goals
- Use interfaces for abstraction
- Handle errors idiomatically
- Organize code into packages

## Theory: Documentation Chapters
- Go Tour chapters "Methods and interfaces": https://go.dev/tour/
- Effective Go chapters "Interfaces and other types", "Errors", "Panic": https://go.dev/doc/effective_go
- Go blog chapter "Errors are values": https://go.dev/blog/errors-are-values
- Package docs `errors` and `fmt`: https://pkg.go.dev/errors and https://pkg.go.dev/fmt

## Checklist
- Create logger interface with 2 implementations
- Add a utility package with tests
- Use `defer` for cleanup paths

## Done When
- Your code handles failure cases explicitly and predictably

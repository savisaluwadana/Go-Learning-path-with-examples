# Phase 06: Standard Library

## Goals
- Build CLIs and HTTP clients using stdlib
- Parse JSON and manage file I/O
- Apply contexts and timeouts

## Theory: Documentation Chapters
- Package docs `net/http` and `encoding/json`: https://pkg.go.dev/net/http and https://pkg.go.dev/encoding/json
- Package docs `os`, `io`, `bufio`: https://pkg.go.dev/os, https://pkg.go.dev/io, https://pkg.go.dev/bufio
- Package docs `time`, `context`, `flag`: https://pkg.go.dev/time, https://pkg.go.dev/context, https://pkg.go.dev/flag
- Structured logging chapter (`log/slog`): https://pkg.go.dev/log/slog

## Checklist
- Build JSON config loader
- Build retryable HTTP client wrapper
- Use `flag` package for CLI input

## Done When
- You can solve most tasks without third-party dependencies

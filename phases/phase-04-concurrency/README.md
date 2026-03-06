# Phase 04: Concurrency

## Goals
- Launch goroutines safely
- Coordinate work with channels and sync primitives
- Stop long-running tasks using context cancellation

## Theory: Documentation Chapters
- Go Tour chapter "Concurrency": https://go.dev/tour/
- Effective Go chapter "Concurrency": https://go.dev/doc/effective_go
- Go blog chapter "Pipelines and cancellation": https://go.dev/blog/pipelines
- Package docs `sync` and `context`: https://pkg.go.dev/sync and https://pkg.go.dev/context

## Checklist
- Implement worker pool
- Build concurrent URL checker
- Detect race conditions with `go test -race ./...`

## Extra Practice Projects
- Parallel directory size calculator
- Concurrent log line counter across many files
- Batch thumbnail metadata scanner with worker pool
- Cancellable web scraper with context timeout

## Done When
- You can explain and prevent common race/deadlock cases

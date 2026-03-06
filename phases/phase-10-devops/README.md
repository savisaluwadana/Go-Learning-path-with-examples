# Phase 10: DevOps and Delivery

## Goals
- Containerize Go applications
- Build/test in CI
- Deploy with environment-based configuration

## Theory: Documentation Chapters
- Package and build docs (`go build`, `go install`, `go test`): https://pkg.go.dev/cmd/go
- Modules reference chapter (versioning and reproducibility): https://go.dev/ref/mod
- Tutorial chapter "Compile and install the application": https://go.dev/doc/tutorial/compile-install
- Build info and runtime versioning (`runtime/debug`): https://pkg.go.dev/runtime/debug

## Checklist
- Write multi-stage Dockerfile
- Add CI workflow for test + build
- Deploy and verify health endpoint

## Done When
- You can ship a repeatable build artifact and deploy confidently

# Phase 07: APIs

## Goals
- Build REST APIs with clean layering
- Add middleware for logging/recovery/auth
- Return consistent JSON responses

## Theory: Documentation Chapters
- Package docs `net/http` server primitives: https://pkg.go.dev/net/http
- Package docs `encoding/json` for request/response bodies: https://pkg.go.dev/encoding/json
- Package docs `context` for request lifecycle: https://pkg.go.dev/context
- Package docs `net/http/httptest` for handler tests: https://pkg.go.dev/net/http/httptest

## Checklist
- Build todo CRUD API
- Add pagination and validation
- Add recover middleware and request logs

## Done When
- API handlers are thin and business logic is testable

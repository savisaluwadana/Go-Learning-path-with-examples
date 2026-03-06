# Phase 09: Architecture and Production

## Goals
- Structure services for long-term maintainability
- Add graceful shutdown and health endpoints
- Improve observability

## Theory: Documentation Chapters
- Package docs `context` for cancellation and propagation: https://pkg.go.dev/context
- Package docs `log/slog` for structured logs: https://pkg.go.dev/log/slog
- Package docs `os/signal` and graceful shutdown patterns: https://pkg.go.dev/os/signal
- Package docs `net/http` (`Server.Shutdown`, timeouts, handlers): https://pkg.go.dev/net/http

## Checklist
- Add `/health` and `/ready` routes
- Add structured logs and request IDs
- Add startup/shutdown signal handling

## Extra Practice Projects
- User profile service with clean architecture layers
- Billing microservice with idempotent operations
- Internal admin API with audit logging and role checks
- Event ingestion service with retries and observability

## Done When
- Service behavior is observable and safe in production

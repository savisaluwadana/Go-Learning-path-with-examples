# Go Zero to Hero Learning Path (With Examples)

This roadmap takes you from absolute beginner to building production-ready Go services.

## How To Use This Path

1. Follow phases in order.
2. Build every exercise yourself before checking references.
3. Commit each phase to Git.
4. Do not skip mini projects.

## Recommended Pace

- Fast track: 8-10 weeks (2-3 hours/day)
- Balanced: 14-16 weeks (1-2 hours/day)
- Weekend only: 5-6 months

## Repository Structure

- `phases/`: phase-by-phase checklists and outcomes
- `phases/*/README.md`: each phase includes a `Theory: Documentation Chapters` section
- `projects/phase-01-unit-converter`: runnable CLI mini project
- `projects/phase-02-contact-manager`: runnable CLI mini project

## Quick Start

```bash
go test ./...
go run ./projects/phase-01-unit-converter -kind temperature -from c -to f -value 37
go run ./projects/phase-02-contact-manager
```

## Phase 0: Setup and Mindset (1-2 days)

### Learn
- Install Go and verify with `go version`
- Workspace basics (`go mod init`, `go run`, `go test`)
- Basic editor setup (VS Code + Go extension)

### Do
- Create project: `go-zero-to-hero`
- Initialize module
- Run `hello.go`

### Example
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
```

---

## Phase 1: Core Language Fundamentals (1-2 weeks)

### Learn
- Variables, constants, types
- Conditionals and loops
- Functions and multiple return values
- Pointers basics

### Practice
- Build a calculator (add/subtract/multiply/divide)
- Write number guessing game
- Reverse a string function

### Mini Project
- CLI unit converter (temperature, distance, weight)

---

## Phase 2: Collections, Structs, Methods (1-2 weeks)

### Learn
- Arrays, slices, maps
- Structs and methods
- Composition over inheritance
- Value vs pointer receivers

### Practice
- Student grade manager using `map[string]int`
- Todo list with slice of structs

### Example
```go
type User struct {
	Name string
	Age  int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}
```

### Mini Project
- Contact manager CLI with CRUD operations in memory

---

## Phase 3: Interfaces, Error Handling, Packages (1-2 weeks)

### Learn
- Interfaces and implicit implementation
- Idiomatic error handling (`if err != nil`)
- Creating and importing local packages
- `defer`, `panic`, `recover` basics

### Practice
- Build logger interface with console/file implementations
- Create package `mathutil` with tests

### Example
```go
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
```

### Mini Project
- File analyzer tool (count lines, words, bytes with proper errors)

---

## Phase 4: Concurrency (2 weeks)

### Learn
- Goroutines
- Channels (buffered and unbuffered)
- `select` statement
- `sync.WaitGroup`, `sync.Mutex`
- Context cancellation

### Practice
- Fan-out worker pool
- Concurrent URL checker

### Example
```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * 2
	}
}
```

### Mini Project
- Concurrent file processor (read many files in parallel)

---

## Phase 5: Testing and Tooling (1 week)

### Learn
- Unit tests with `testing`
- Table-driven tests
- Benchmarks
- `go test ./...`, `go fmt`, `go vet`

### Practice
- Add tests for every package built so far
- Benchmark two string parsing methods

### Mini Project
- Refactor old mini projects to be testable with >80% coverage

---

## Phase 6: Standard Library Power (1-2 weeks)

### Learn
- `net/http`, `encoding/json`
- `os`, `io`, `bufio`
- `time`, `context`, `log/slog`
- `flag` for CLI args

### Practice
- Build JSON config loader
- Build HTTP client wrapper with timeout + retry

### Mini Project
- REST client CLI (fetch posts/users and pretty print output)

---

## Phase 7: Build Real APIs (2-3 weeks)

### Learn
- HTTP server routing/middleware
- Request validation
- Clean handler/service/repository separation
- Environment-based config

### Practice
- CRUD API for todos
- Add middleware: logging, recovery, auth token check

### Mini Project
- Task API with:
  - Create/read/update/delete tasks
  - Pagination
  - Proper status codes and JSON errors

---

## Phase 8: Databases and Persistence (2 weeks)

### Learn
- SQL basics (PostgreSQL)
- `database/sql` and query patterns
- Migrations
- Transactions

### Practice
- Connect API to PostgreSQL
- Add unique constraints and indexes

### Mini Project
- Expense tracker backend with user + category + transaction tables

---

## Phase 9: Architecture and Production Practices (2 weeks)

### Learn
- Project structure for growing services
- Dependency injection basics
- Graceful shutdown
- Observability (logs, metrics, health checks)

### Practice
- Add `/health` and `/ready` endpoints
- Add request IDs and structured logs

### Mini Project
- Production-ready user service with layered architecture

---

## Phase 10: Deployment and DevOps Basics (1 week)

### Learn
- Dockerizing Go apps
- Multi-stage builds
- Environment configs and secrets
- CI pipeline basics

### Practice
- Dockerize Task API
- Add GitHub Actions: test + build

### Mini Project
- Deploy to a cloud VM/container platform with monitoring

---

## Hero Projects (Portfolio)

Build at least 2 end-to-end projects:

1. **Realtime Chat Service**
   - WebSockets
   - Rooms
   - Broadcast + persistence

2. **E-commerce Backend**
   - Auth, products, cart, orders
   - PostgreSQL + transactions
   - Admin endpoints

3. **Job Queue Processor**
   - Worker pool
   - Retry with backoff
   - Dead-letter handling

## Practice Project Ladder (Beginner to Advanced)

Choose at least 2 projects from each level.

### Beginner Projects (Phases 0-2)

- CLI calculator with operation history
- Number guessing game with difficulty levels
- Password strength checker
- File extension counter (`.txt`, `.go`, `.md`)
- CSV student marks analyzer
- Unit converter (temperature/distance/weight)
- Contact manager CLI (in-memory CRUD)
- Flashcard quiz CLI
- Expense tracker CLI (in-memory)
- Simple markdown word counter

### Intermediate Projects (Phases 3-5)

- File analyzer with pluggable output formatters via interfaces
- Retryable HTTP fetcher with custom error types
- Log aggregator with console and file logger backends
- Concurrent image metadata scanner
- Worker pool job runner with timeout support
- URL health checker with concurrency and summary report
- CLI backup tool with gzip option
- Table-driven test suite for a reusable utility package
- Benchmark project for string/token parsing approaches
- Mini package library (`mathutil`, `strutil`, `dateutil`) with tests

### Upper-Intermediate Projects (Phases 6-8)

- REST client CLI with pagination and filters
- Todo REST API with layered architecture
- URL shortener API (in-memory, then PostgreSQL)
- Notes API with tags/search and validation
- Expense tracker backend with PostgreSQL transactions
- Inventory service with low-stock alerts
- Authentication service with JWT and refresh tokens
- API rate limiter middleware
- Blog backend with comments and soft delete
- Event booking service with transactional seat reservation

### Advanced Projects (Phases 9-10)

- Production-ready user service with health/readiness endpoints
- Realtime chat server with WebSockets and room management
- Job queue processor with retries and dead-letter queue
- E-commerce backend with orders and payment state machine
- API gateway with auth, rate limit, and request tracing
- Distributed scheduler for periodic jobs
- Multi-tenant SaaS backend with tenant isolation
- Notification service (email/SMS/push) with worker consumers
- Observability-focused microservice (logs, metrics, traces)
- Deployable service template with Docker + CI + release tagging

## Daily Practice Checklist

- Read 10-20 pages from Go docs/book
- Solve 1 algorithm/problem in Go
- Write 30-60 lines of clean Go
- Add tests for new code
- Review one older file and improve it

## Suggested Learning Resources

- Official Go Tour
- Effective Go
- Go by Example
- Go blog (official)
- Standard library docs (`pkg.go.dev`)

## Completion Criteria (Zero to Hero)

You are "hero level" when you can:

- Design and ship a Go REST API from scratch
- Write reliable concurrent code without race conditions
- Test business logic and handlers confidently
- Use PostgreSQL with transactions and migrations
- Dockerize and deploy service with logs and health checks

## Next Step

Start with **Phase 0** today and create your first commit:

`chore: setup go module and hello world`

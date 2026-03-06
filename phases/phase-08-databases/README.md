# Phase 08: Databases

## Goals
- Use PostgreSQL with Go
- Execute queries safely
- Manage migrations and transactions

## Theory: Documentation Chapters
- Tutorial chapter "Accessing a relational database": https://go.dev/doc/tutorial/database-access
- Database docs "Opening a database handle" and "Querying for data": https://go.dev/doc/database/open-handle and https://go.dev/doc/database/querying
- Database docs "Executing SQL statements that don't return data": https://go.dev/doc/database/change-data
- Database docs "Executing transactions": https://go.dev/doc/database/execute-transactions
- Package docs `database/sql`: https://pkg.go.dev/database/sql

## Checklist
- Model schema for tasks/expenses
- Implement repository methods
- Add migration files and rollback path

## Done When
- Data writes are transactional and constraints are enforced

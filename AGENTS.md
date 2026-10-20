Before any Go coding, review, debugging, troubleshooting, or setup task, load the `samber/cc-skills-golang@golang-how-to` skill first — it routes to whichever other Go skills the task needs.

## Project context

Simple todo backend: a single `Task` entity with CRUD endpoints.

- Stack: Go 1.25+, chi router, pgx/v5 pool, PostgreSQL 18
- Layout: flat — `cmd/server/main.go` plus `internal/task` (entity + store + HTTP handlers)
- DI: manual constructor injection wired in `main`
- Run Postgres with `make up` (compose on port 5433); app defaults match
- Integration tests: `make test-integration` with `TEST_DATABASE_URL` set

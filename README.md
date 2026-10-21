# Team-max

Internal project management platform (Plane-inspired). This repo follows ADR-0001 (modular monolith) and Phase 0 decisions in `docs/15` / `docs/adr/`.

## Quick start

```bash
make up        # start Postgres on localhost:5433
make build     # build all packages
make test      # run unit tests
make run       # run API server
```

## Endpoints

- `GET /healthz`
- `GET /readyz`
- `GET/POST /tasks`, `GET/PUT/DELETE /tasks/{id}` (work items)

## Environment

- `HTTP_ADDR` (default `:8080`)
- `DATABASE_URL` (default `postgres://todo:todo@localhost:5433/todo?sslmode=disable`)

## Project layout

- `cmd/server/` – entrypoint
- `internal/modules/*` – bounded contexts (domain/application/infrastructure/interfaces)
- `internal/shared/kernel/` – cross-cutting primitives
- `internal/bootstrap/` – wiring
- `migrations/` – SQL migrations
- `docs/` – requirements/ADRs

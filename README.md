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
- `GET /swagger/index.html` — Swagger UI (served only when `ENABLE_SWAGGER=true`, the default)

## API docs (OpenAPI)

The spec is generated from handler annotations in `internal/modules/*/interfaces/http` with
[swag](https://github.com/swaggo/swag), following `docs/10` ("generate OpenAPI documentation
from the interface layer"):

```bash
make swagger     # regenerate internal/api/docs/{docs.go,swagger.json,swagger.yaml}
```

Re-run `make swagger` after changing routes, handlers, or DTOs, and commit the regenerated files.
Add new annotations with `go tool swag fmt`.

Set `ENABLE_SWAGGER=false` to hide the UI in production; the spec files stay in the binary.

## Environment

- `HTTP_ADDR` (default `:8080`)
- `DATABASE_URL` (default `postgres://todo:todo@localhost:5433/todo?sslmode=disable`)
- `ENABLE_SWAGGER` (default `true`)

## Project layout

- `cmd/server/` – entrypoint
- `internal/modules/*` – bounded contexts (domain/application/infrastructure/interfaces)
- `internal/shared/kernel/` – cross-cutting primitives
- `internal/bootstrap/` – wiring
- `migrations/` – SQL migrations
- `docs/` – requirements/ADRs

# 10 --- Architecture and Technical Decisions

## 1. Recommended starting architecture

**Proposed:** Modular monolith with clear bounded contexts, a relational
database, background worker, and object storage adapter.

Rationale: - Lower operational complexity than early microservices. -
Transactions are easier for core invariants. - Context boundaries can
later guide extraction if evidence justifies it. - Developers can test
domain behavior without starting infrastructure.

## 2. Logical layers

### Domain

-   Entities and aggregate roots.
-   Value objects.
-   Domain services for rules that do not naturally belong to one
    entity.
-   Domain events.
-   Repository interfaces where persistence abstraction supports the
    domain.

Must not depend on framework, database, network, UI, or provider SDKs.

### Application

-   Use cases/command handlers and query handlers.
-   Authorization orchestration.
-   DTOs and input validation orchestration.
-   Transaction boundaries.
-   Ports for persistence, clock, ID generation, event publishing,
    email, files, and external integrations.
-   Mapping between domain results and application responses.

### Infrastructure

-   ORM mappings and repository implementations.
-   Database migrations and transaction implementation.
-   Message broker, outbox dispatcher, email provider, object storage,
    and external API adapters.
-   Framework-specific authentication adapters.

### Interface adapters

-   HTTP controllers/handlers, serializers, API schema, CLI, worker
    consumers.
-   Thin adapters: parse, authenticate, call use case, map result to
    HTTP.

## 3. Suggested repository layout

Adapt to the language; this is conceptual:

``` text
src/
  modules/
    identity/
      domain/
      application/
      infrastructure/
      interfaces/
    organizations/
    teams/
    projects/
    work_items/
    workflows/
    planning/
    collaboration/
    notifications/
    audit/
    reporting/
    integrations/
  shared/
    kernel/          # only stable primitives, IDs, Result/Error abstractions
    observability/
  bootstrap/
tests/
  unit/
  integration/
  architecture/
  e2e/
docs/
  adr/
```

Avoid a single global `services/`, `repositories/`, and `models/` folder
that mixes every context together.

## 4. Persistence strategy

-   Use a relational database for transactional project/work data.
-   Each context owns its schema/tables logically, even if one physical
    database is used.
-   Use migrations committed with code.
-   Keep ORM entities separate from domain entities when the ORM model
    would leak persistence concerns into domain logic.
-   Use unique constraints and foreign keys as a second line of defense
    for invariants.
-   Use optimistic concurrency for frequently edited records.
-   Use transactions for aggregate changes and associated audit/outbox
    records.
-   Use an outbox for durable event publishing; consumers should be
    idempotent.
-   Avoid distributed transactions. Model eventual consistency
    explicitly.

## 5. API strategy

-   Version the public API when breaking changes are possible.
-   Use stable IDs and human-readable keys.
-   Apply authorization before data is serialized.
-   Use explicit request/response DTOs; never bind untrusted request
    bodies directly to domain entities.
-   Use pagination, bounded filters, and allowlisted sorting.
-   Standardize validation, authorization, conflict, not-found, and
    internal-error responses.
-   Do not expose stack traces, SQL details, or secret-bearing fields.
-   Generate OpenAPI documentation from the interface layer where
    possible.

Suggested response envelope, if chosen consistently:

``` json
{
  "success": true,
  "data": {},
  "error": null,
  "message": "Work item created"
}
```

Do not force an envelope if it makes HTTP semantics or pagination
unnecessarily awkward; consistency is more important than a particular
shape.

## 6. Error taxonomy

-   `ValidationError`: Input shape or value is invalid.
-   `Unauthenticated`: No valid actor/session.
-   `Forbidden`: Actor lacks permission.
-   `NotFound`: Resource absent or intentionally hidden by policy.
-   `Conflict`: Stale version, duplicate key, or incompatible state.
-   `DomainRuleViolation`: Business invariant failed.
-   `RateLimited`: Request budget exceeded.
-   `DependencyUnavailable`: Required external dependency unavailable.
-   `InternalError`: Unexpected failure; details only in protected logs.

Map domain/application errors to transport responses in one place.

## 7. Background processing

Use a worker for email, notifications, webhooks, imports, exports, and
scheduled reminders. Jobs should include stable IDs, correlation IDs,
bounded retry policies, deadlines, and idempotency keys. Dead-letter or
permanently failed jobs need an operational inspection path.

## 8. Architecture decision records (ADRs)

Create one ADR per consequential choice: - ADR-001: Modular monolith
versus services. - ADR-002: Authentication/session strategy. - ADR-003:
Organization/workspace tenancy model. - ADR-004: Workflow configuration
and versioning. - ADR-005: Outbox and background processing. - ADR-006:
File storage and malware scanning. - ADR-007: Search engine versus
database search. - ADR-008: Audit retention and deletion semantics.

Each ADR includes context, options, decision, consequences, status, and
date. Revisit only when assumptions materially change.

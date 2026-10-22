# Implementation Plan

## Phase 0 — Discovery and decisions
- [x] Stakeholder-approved MVP scope (resolved in `docs/15`)
- [x] Confirmed tenancy, visibility, roles, retention, RPO/RTO, deployment assumptions
- [x] Context map and first ADRs (ADR-0001-0008)
- [x] Initial threat model and test strategy
**Exit gate:** No unresolved policy decision that would force a major data-model rewrite. *(Subject to named-owner sign-off; ADR-0001 proposed pending technical-owner sign-off.)*

## Phase 1 — Foundation
- [x] Repository structure and dependency rules (modular monolith layout under `internal/modules/`)
- [x] CI pipeline: formatting, lint, type checking, unit tests, integration tests (Makefile, `.golangci.yml`, `.github/workflows/ci.yml`)
- [x] Authentication, actor context, organization membership, and authorization policy (scaffolds; implementation in progress)
- [x] Database migrations, structured logs, request IDs, health/readiness checks (`/healthz`, `/readyz`, bootstrap)
**Exit gate:**
- [x] Cross-organization authorization tests pass (`internal/authz` tests deny cross-org access)
- [x] Domain layer is framework-independent (arch test enforces no framework imports in domain)
- [x] New developer can run tests and app using documented instructions (README updated)

### Phase 1 Remaining Tasks
- [x] Enforce inward dependency rules (architecture tests for domain)
- [x] Implement actor context middleware and deny-by-default authorization policy
- [x] Add cross-organization authorization tests
- [x] Document local setup in README
- [ ] Implement local auth (ADR-0002): email/password, Argon2id, sessions, httpOnly cookies (post-exit gate or continue)
- [ ] Add org/workspace/membership/session migrations and models (ADR-0003: auto-create workspace per org) (post-exit gate or continue)

## Phase 2 — Projects and work items
- [ ] Project CRUD/archive, project keys, project membership
- [ ] Work-item create/update/assign/search, comments, labels, due dates
- [ ] Parent-child relationships and activity history
- [ ] API documentation and core UI
**Exit gate:**
- [ ] Key invariants have unit tests
- [ ] Database uniqueness/concurrency constraints are tested
- [ ] Unauthorized reads and writes are rejected

## Phase 3 — Workflow and notifications
- [ ] Workflow definitions, states, transitions, validation (immutable versions per ADR-0004)
- [ ] In-app notifications and email delivery
- [ ] Outbox/worker and retry visibility (ADR-0005)
- [ ] Audit logging for configuration changes
**Exit gate:**
- [ ] Invalid transitions cannot be performed through API or UI
- [ ] Worker retries are idempotent
- [ ] Notification recipients are authorized

## Phase 4 — Reporting and pilot hardening
- [ ] Project dashboard, overdue items, status summary
- [ ] Performance baseline, backup/restore test, incident runbooks
- [ ] Accessibility and usability review
- [ ] Pilot feedback and defect triage
**Exit gate:**
- [ ] Metrics reconcile with test fixtures
- [ ] Restore procedure is demonstrated
- [ ] No unresolved critical/high security issue

## Phase 5 — Advanced capabilities
- [ ] Cycles, modules, milestones, dependencies
- [ ] Saved views and custom fields
- [ ] Workflow automation and approvals
- [ ] Cross-project workload reports
- [ ] Webhooks/source-control integration
- [ ] Knowledge pages/wiki
*Each capability requires a BRD addendum, domain invariants, acceptance tests, and an ADR if it changes architecture.*

## Release acceptance checklist
- [ ] Requirements and acceptance criteria reviewed
- [ ] TDD workflow used for core domain behavior
- [ ] Unit, integration, authorization, and regression tests pass
- [ ] Database migration and rollback plan reviewed
- [ ] API/event contracts and docs updated
- [ ] Logging, metrics, and alerts reviewed
- [ ] Backup and recovery implications reviewed
- [ ] Security review completed
- [ ] Feature flags and rollout plan prepared if needed
- [ ] Product owner signs off against agreed scope

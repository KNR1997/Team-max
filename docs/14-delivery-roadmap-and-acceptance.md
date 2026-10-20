# 14 --- Delivery Roadmap and Acceptance Gates

## Phase 0 --- Discovery and decisions

Status: decisions recorded 2026-10-20 (see `docs/15` and `docs/adr/`).
The exit gate is satisfied subject to named-owner sign-off in `docs/01`
§1; ADR-0001 (modular monolith) remains proposed until the technical
owner signs it.

Deliverables: - Stakeholder-approved MVP scope. - Confirmed tenancy,
visibility, roles, retention, RPO/RTO, and deployment assumptions. -
Context map and first ADRs. - Initial threat model and test strategy.

Exit gate: - No unresolved policy decision that would force a major
data-model rewrite.

## Phase 1 --- Foundation

Deliverables: - Repository structure and dependency rules. - CI
pipeline: formatting, lint, type checking, unit tests, integration
tests. - Authentication, actor context, organization membership, and
authorization policy. - Database migrations, structured logs, request
IDs, health/readiness checks.

Exit gate: - Cross-organization authorization tests pass. - Domain layer
is framework-independent. - New developer can run tests and app using
documented instructions.

## Phase 2 --- Projects and work items

Deliverables: - Project CRUD/archive, project keys, project
membership. - Work-item create/update/assign/search, comments, labels,
due dates. - Parent-child relationships and activity history. - API
documentation and core UI.

Exit gate: - Key invariants have unit tests. - Database
uniqueness/concurrency constraints are tested. - Unauthorized reads and
writes are rejected.

## Phase 3 --- Workflow and notifications

Deliverables: - Workflow definitions, states, transitions, validation. -
In-app notifications and email delivery. - Outbox/worker and retry
visibility. - Audit logging for configuration changes.

Exit gate: - Invalid transitions cannot be performed through API or
UI. - Worker retries are idempotent. - Notification recipients are
authorized.

## Phase 4 --- Reporting and pilot hardening

Deliverables: - Project dashboard, overdue items, status summary. -
Performance baseline, backup/restore test, incident runbooks. -
Accessibility and usability review. - Pilot feedback and defect triage.

Exit gate: - Metrics reconcile with test fixtures. - Restore procedure
is demonstrated. - No unresolved critical/high security issue.

## Phase 5 --- Advanced capabilities

Prioritize from user evidence: - Cycles, modules, milestones,
dependencies. - Saved views and custom fields. - Workflow automation and
approvals. - Cross-project workload reports. - Webhooks/source-control
integration. - Knowledge pages/wiki.

Each capability requires a BRD addendum, domain invariants, acceptance
tests, and an ADR if it changes architecture.

## Release acceptance checklist

-   [ ] Requirements and acceptance criteria reviewed.
-   [ ] TDD workflow used for core domain behavior.
-   [ ] Unit, integration, authorization, and regression tests pass.
-   [ ] Database migration and rollback plan reviewed.
-   [ ] API/event contracts and docs updated.
-   [ ] Logging, metrics, and alerts reviewed.
-   [ ] Backup and recovery implications reviewed.
-   [ ] Security review completed.
-   [ ] Feature flags and rollout plan prepared if needed.
-   [ ] Product owner signs off against agreed scope.

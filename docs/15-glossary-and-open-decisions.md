# 15 --- Glossary and Open Decisions

## Glossary

Canonical project vocabulary lives in [`GLOSSARY.md`](../GLOSSARY.md) at
the repository root. Update terms there; do not duplicate them in this
document.

## Decisions resolved in Phase 0 (2026-10-20)

All DEC items below were approved by the Phase 0 reviewers. Architectural
consequences are recorded in [`docs/adr/`](adr/).

| ID | Decision | Resolution |
| --- | --- | --- |
| DEC-001 | Organization versus workspace hierarchy | Organization → Workspace modeled; exactly one workspace per organization, owned by the organization and auto-created at org creation. No workspace-level membership. |
| DEC-002 | Project visibility | `Restricted` by default; `Organization-visible` must be chosen explicitly and grants non-members read-only access. Visibility is changed by a project admin or an organization admin. |
| DEC-003 | Can contributors create projects? | No. Project creation requires an explicit permission (organization admin, or a member granted the capability). |
| DEC-004 | Can assignee be empty? | Yes, but dashboards surface unassigned active work prominently. |
| DEC-005 | Can users delete work items? | Archive by default; no ordinary hard-delete path exists in MVP. |
| DEC-006 | Project key mutability | Immutable after creation for MVP; any future rename needs aliases/redirects. |
| DEC-007 | Workflow versioning | Published versions are immutable; edits create a new version and in-flight items migrate explicitly. |
| DEC-008 | Authentication | Local email + password behind an OIDC-ready port; no identity provider is mandated. |
| DEC-009 | Background jobs | Postgres transactional outbox drained by a polling worker; no separate broker for MVP. |
| DEC-010 | Search | Postgres full-text and trigram indexing; a dedicated engine only on measured need. |
| DEC-011 | File storage | Private S3-compatible object storage with short-lived signed URLs; malware scanning deferred, adapter seam present. |
| DEC-012 | Audit retention | Append-only and immutable to ordinary roles confirmed; the retention period is set with legal/security before production (Phase 4 gate). |
| DEC-013 | RPO/RTO | Working targets RPO ≤ 24 h and RTO ≤ 4 h with daily backups; restore demonstrated in Phase 4. No compliance driver tightens these today. |
| DEC-014 | API response envelope | The `success` / `data` / `error` / `message` envelope is standardized and applied consistently (`docs/10` §5). |
| DEC-015 | Time tracking and estimates | Post-MVP unless pilot evidence demands it. |

ADR references:

-   DEC-001 → [ADR-0003](adr/0003-single-workspace-tenancy.md)
-   DEC-007 → [ADR-0004](adr/0004-immutable-workflow-versions.md)
-   DEC-008 → [ADR-0002](adr/0002-local-auth-with-oidc-seam.md)
-   DEC-009 → [ADR-0005](adr/0005-postgres-outbox-and-worker.md)
-   DEC-010 → [ADR-0007](adr/0007-database-search.md)
-   DEC-011 → [ADR-0006](adr/0006-private-object-storage.md)
-   DEC-012 → [ADR-0008](adr/0008-audit-retention-and-deletion.md)

## Additional Phase 0 resolutions

Items decided in Phase 0 that were not part of the DEC register:

-   **Deployment assumptions:** single-host container deployment (Docker
    Compose) for MVP; separate development and production environments
    with separate credentials and data; staging added when a shared test
    environment earns its keep; Kubernetes out of scope for MVP.
-   **Roles:** organization roles are Owner, Admin, and Member (Guest
    later); project roles are Admin, Contributor, and Viewer. Roles are
    presets over an explicit permission matrix, never role-name string
    checks in code. Team lead is a field on the Team aggregate, not an
    organization role.
-   **Comment policy:** authors edit and delete their own comments;
    project admins may delete for moderation; deletion is soft, keeping a
    tombstone and the minimum audit evidence required by retention
    policy.
-   **User deletion:** deleted users are tombstoned (stable ID and
    historical authorship preserved) and cannot be deleted while they own
    active work. See
    [ADR-0008](adr/0008-audit-retention-and-deletion.md).
-   **Confirmed proposed defaults:** project lifecycle is Active and
    Archived only; work item types are Task, Bug, Feature, and Story;
    membership lifecycle is Invited, Active, Suspended, and Removed;
    project keys are 2--10 uppercase alphanumerics.
-   **Approvers:** one product owner and one technical lead sign off on
    MVP scope and the release acceptance checklist (names to be recorded
    in `docs/01` §1).

## Change management

When a decision changes: 1. Update this table and relevant BRD. 2.
Explain the reason and impact. 3. Create/update an ADR for architectural
consequences. 4. Update domain model and acceptance tests. 5. Plan data
migration/backward compatibility.

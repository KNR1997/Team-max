---
status: accepted
date: 2026-10-20
---

# One workspace per organization, owned by the organization

The tenancy model is Organization → Workspace → (Teams, Projects) → Work items. A Workspace is modeled as a real child of the Organization that owns teams and projects — both carry `workspace_id` alongside `organization_id` — but it is capped at exactly one per organization for the MVP and auto-created at organization creation. There is deliberately **no workspace-level membership layer**: access is granted only through organization membership, keeping the authorization model single-layer.

**Considered options:** full multi-workspace from day one (rejected: a second membership/roles layer for a capability no pilot needs); organization-only with no workspace entity (rejected: retrofitting the container later touches nearly every table and permission check, which is exactly the data-model rewrite the Phase 0 exit gate forbids).

**Consequences:** "Organization-visible" and "workspace-visible" are identical while the cap holds; lifting the cap later is a data migration plus new membership work, not a model rewrite; queries can filter on `workspace_id` from day one.

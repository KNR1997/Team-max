# 04 --- BRD: Project Management Context

## 1. Context purpose

Manage project identity, lifecycle, settings, membership, and
visibility. Project execution details live in Work Management.

## 2. Project fields

Required: - `id` (UUID or equivalent stable identifier) -
`organization_id` - `name` - `project_key` (unique within agreed
scope) - `status` - `created_at`, `created_by` - `updated_at`,
`updated_by`

Optional: - Description, lead, start/end dates, timezone, icon/color,
archived timestamp, default workflow, estimation configuration.

## 3. Project lifecycle

Proposed states: - `Active` - `Archived`

A later release may add `Planned`, `Paused`, and `Completed` if those
states represent meaningful business behavior.

Rules: - Archived projects are read-only by default. - Restoring a
project requires project-management permission. - Project keys cannot
collide within their uniqueness scope. - Changing a project key must not
silently break links or external integrations. Prefer immutable keys for
MVP; if rename is required, provide redirects/aliases. - Project
deletion should be restricted and normally replaced by archive.

## 4. Project membership and visibility

Visibility modes: - **Organization-visible:** Members of the
organization can discover/access the project, subject to guest
restrictions. - **Restricted:** Only explicit project members and
authorized administrators can access it.

Policy decisions: - The visibility model must be enforced in list
queries, direct reads, search, exports, notifications, and background
jobs. - Adding a user to a project requires active organization
membership. - Removing a project member does not remove their historical
comments or audit entries. - Project lead must be an eligible
organization member.

## 5. Project configuration

A project may configure: - Workflow definition. - Labels and label
colors. - Default assignee policy. - Priority vocabulary. - Estimation
method. - Cycles/modules/pages feature flags. - Notification defaults. -
Timezone.

Configuration changes should be validated, versioned where necessary,
and audited. Deleting a configuration item already referenced by work
should be prevented or handled through a migration/remapping flow.

## 6. Project key and work-item key policy

-   Project key format: uppercase ASCII letters/digits, 2--10 characters
    by default; configurable only if needed.
-   A work-item key consists of project key plus a monotonically
    allocated sequence, e.g. `OPS-104`.
-   Sequence allocation must be concurrency-safe.
-   Do not derive identity from a database row count.
-   Never reuse a retired work-item sequence.
-   The UUID remains the internal identity; human-readable key is a
    stable external reference.

## 7. Acceptance criteria

-   Project creation rejects duplicate keys in the same scope.
-   Restricted projects do not leak through search, counts,
    notifications, or API responses.
-   Archived projects reject ordinary mutations.
-   Project configuration changes create an audit event.
-   Key allocation remains unique under concurrent creation.

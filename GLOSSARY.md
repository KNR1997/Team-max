# Team-max

The internal project management platform: an organization runs teams and projects, members track work items through configurable workflows, and managers get a trustworthy view of progress.

## Language

### Tenancy and access

**Organization**:
The top-level boundary for a company's data; every user, team, project, and work item belongs to exactly one organization.
_Avoid_: company, tenant, account

**Workspace**:
The container inside an organization that owns its teams and projects. An organization has exactly one workspace in the MVP; it is not a membership boundary — people gain access through organization membership.
_Avoid_: space, hub, org unit

**Membership**:
An active relationship between a person and an organization that grants scoped participation. There is no separate workspace-level membership.
_Avoid_: account link, access grant

**Organization Owner**:
The single person accountable for an organization, including ownership transfer and deletion initiation.
_Avoid_: root admin

**Organization Admin**:
An organization member who manages settings, members, teams, and projects, but cannot transfer ownership unless explicitly permitted.
_Avoid_: superuser

**Organization Member**:
A person with ordinary participation rights; their actual project access depends on project membership and visibility.
_Avoid_: user, participant

**Project Admin**:
A project role that manages project settings, members, and workflow configuration.
_Avoid_: project owner, manager (ambiguous with Team lead)

**Contributor**:
A project role that creates and updates work items within the project's policies.
_Avoid_: editor, member (ambiguous with Organization Member)

**Viewer**:
A project role with read-only access to a project.
_Avoid_: reader, guest (Guest is a separate, later role)

**Team**:
A named group of members who share ownership of work within a workspace.
_Avoid_: squad, group, department

**Team lead**:
The member recorded as owning a Team, stored as a field on the Team — not an organization role.
_Avoid_: team manager, squad lead

**Project**:
A scoped body of work with its own key, members, workflow, and work items; belongs to exactly one workspace.
_Avoid_: board, space, initiative

**Project key**:
The project's human-readable code — 2–10 uppercase alphanumerics, unique in its scope, immutable after creation.
_Avoid_: slug, project code, abbreviation

**Project membership**:
The explicit list of people allowed to participate in a project.
_Avoid_: project users, collaborators

**Visibility**:
A project's setting controlling who may see it: `Restricted` (explicit members and admins only) or `Organization-visible` (all organization members, read-only).
_Avoid_: privacy, access level, sharing

### Work

**Work item**:
The trackable unit of work inside a project — a `Task`, `Bug`, `Feature`, or `Story`.
_Avoid_: ticket, issue, card; bare "task" (that is one type)

**Work item key**:
A stable human-readable reference such as `OPS-104`: the project key plus a monotonically allocated, never-reused sequence number.
_Avoid_: ticket number, row ID

**Workflow definition**:
The versioned set of states and allowed transitions a project (or work-item type) uses.
_Avoid_: status list, board config

**Workflow state**:
A configured status a work item can occupy, identified by a stable ID rather than its display label.
_Avoid_: status, column, tag

**Transition**:
A permitted move from one workflow state to another.
_Avoid_: status change, state edit

**Semantic category**:
The reporting meaning of a workflow state — `Backlog`, `Unstarted`, `Started`, `Completed`, or `Cancelled`.
_Avoid_: status group, kind

### Collaboration

**Comment**:
Discussion authored by a person and attached to a work item.
_Avoid_: message, note, reply

**Attachment**:
A file attached to a work item, held in private object storage with its metadata owned by the application.
_Avoid_: upload, file record

**Notification**:
A permission-checked message generated for a person, delivered to the in-app inbox or by email.
_Avoid_: alert, email, ping

**Activity history**:
The user-facing record of changes to a project or work item.
_Avoid_: audit log (that is separate), timeline feed

**Audit log**:
Append-only evidence of security- and governance-relevant actions, never editable by ordinary application users.
_Avoid_: activity history, event log

### Domain building blocks

**Aggregate**:
A cluster of domain objects treated as one consistency boundary.

**Aggregate root**:
The only object through which an aggregate is modified externally.

**Bounded context**:
A boundary within which a domain model and its language have a specific meaning.

**Command**:
A request to change state by expressing intent.

**Domain event**:
A fact that occurred inside a context.

**Integration event**:
A versioned event contract intended for another context or external system.

**Invariant**:
A rule that must always hold true.

**Port**:
An interface describing a dependency that application or domain logic requires.

**Adapter**:
An implementation of a port for a particular technology.

**Repository**:
An abstraction for retrieving and persisting an aggregate.

**Value object**:
An object defined by its values rather than identity.

**Entity**:
An object whose identity persists through changes.

**Outbox pattern**:
Persisting a business change and its event record atomically, then publishing asynchronously.

**Optimistic concurrency**:
Detecting conflicting updates with a version token rather than silently overwriting.

**Ubiquitous language**:
The shared domain vocabulary used by stakeholders and code.

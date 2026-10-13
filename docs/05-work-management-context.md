# 05 --- BRD: Work Management Context

## 1. Context purpose

Represent units of work and enforce the rules governing their lifecycle,
assignment, hierarchy, and relationships.

## 2. Work item types

Initial types: `Task`, `Bug`, `Feature`, `Story`. Types should be
configurable only after the basic model and reporting semantics are
stable.

## 3. Work item fields

-   Stable ID and human-readable key.
-   Organization and project IDs.
-   Type, title, description.
-   State ID.
-   Priority.
-   Reporter/creator and optional assignee.
-   Labels.
-   Due date.
-   Optional parent ID.
-   Created/updated timestamps and actor IDs.
-   Archived/deleted marker when policy permits.
-   Version/concurrency token.

Later fields: estimate, start date, cycle, module, milestone, custom
fields, time entries.

## 4. Invariants and rules

1.  Title is required, trimmed, and bounded by a configured maximum.
2.  Work item belongs to exactly one project and organization.
3.  Assignee must be eligible for the project at the time of assignment.
4.  Parent and child must belong to the same organization; cross-project
    hierarchy requires an explicit future policy.
5.  A work item cannot be its own parent, and parent relationships
    cannot form cycles.
6.  A state must belong to the active workflow configuration for that
    project/item type.
7.  Work-item key is unique and never reused.
8.  Archived projects reject ordinary work-item mutations.
9.  Every successful meaningful mutation updates `updated_at` and actor
    metadata.
10. Concurrent updates must not silently overwrite another user's
    changes. Use optimistic concurrency or another explicit conflict
    strategy.

## 5. Commands/use cases

-   `CreateWorkItem`
-   `UpdateWorkItemDetails`
-   `AssignWorkItem`
-   `ChangeWorkItemPriority`
-   `SetWorkItemDueDate`
-   `ChangeWorkItemState`
-   `SetParentWorkItem`
-   `AddWorkItemRelation`
-   `ArchiveWorkItem`
-   `RestoreWorkItem`
-   `SearchWorkItems`
-   `GetWorkItemDetails`

Commands express intent; avoid a generic update endpoint that allows
arbitrary fields to bypass business rules.

## 6. Relationships

Initial relationship types: - `ParentOf` / `ChildOf` - `RelatesTo` -
`Blocks` / `BlockedBy` (P1)

For blocking links, prevent contradictory/self-referential links and
define whether completion is blocked by unresolved dependencies. If
blocking dependencies are introduced, test cycles and concurrent
changes.

## 7. Search and listing

-   Search/list operations must apply authorization before returning
    records or aggregates.
-   Filtering and sorting are server-side.
-   Pagination must be stable and bounded.
-   Avoid allowing arbitrary unindexed sort fields.
-   Define case sensitivity and partial-match behavior.
-   Search results must not expose descriptions or counts from
    inaccessible projects.

## 8. Domain events

Potential events: - `WorkItemCreated` - `WorkItemDetailsChanged` -
`WorkItemAssigned` - `WorkItemStateChanged` - `WorkItemParentChanged` -
`WorkItemArchived`

Events describe facts that already happened and should include stable
IDs, organization scope, actor ID, occurrence time, and schema version.
Do not include secrets or unnecessary personal data.

## 9. Acceptance criteria

-   Invalid hierarchy changes are rejected.
-   Concurrent key generation produces no duplicate keys.
-   A stale update returns a conflict response rather than overwriting
    newer data.
-   Search and detail endpoints enforce identical access rules.
-   Domain tests cover every invariant listed in this document.

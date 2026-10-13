# 08 --- BRD: Planning, Reporting, and Dashboards

## 1. Context purpose

Support planning and management visibility without making analytics the
source of truth for operational data.

## 2. Planning features

### Cycles (P1)

A cycle is a time-boxed period associated with a project or team. It has
name, start/end times, status, and selected work items.

Rules: - End must be after start. - Cycle overlap policy must be
explicit; default may allow overlap unless the team chooses a strict
sprint model. - Moving items between cycles is audited. - Cycle
completion does not automatically complete every item. - Carryover
behavior must be defined and visible.

### Modules (P1)

A module groups work items around a theme, feature, component, or
release. Membership is many-to-many unless a product decision requires
otherwise. Removing a module must not delete work items.

### Milestones (P1)

A milestone is a target date/outcome. It may aggregate work items and
report progress using a documented calculation. A milestone date change
should be audited.

## 3. Reporting definitions

Define metrics centrally and test them: - **Open work:** Items not in a
completed or cancelled semantic state. - **Overdue work:** Open items
whose due date is earlier than the current date in the project's
timezone. - **Blocked work:** Items with at least one unresolved
blocking dependency, if dependency feature is enabled. - **Completion
rate:** Completed items divided by the explicitly defined denominator
and period. - **Cycle completion:** Items completed within the cycle,
with policy for items added after cycle start documented. -
**Workload:** Use estimates or explicit effort data; item count alone
must not be presented as an accurate measure of effort.

Avoid using state display names to calculate metrics.

## 4. Dashboard access

-   Authorization is applied to every query and drill-down.
-   A dashboard cannot reveal aggregate counts for hidden projects if
    those counts themselves are sensitive.
-   Exports follow the same access rules as interactive screens.
-   Report generation jobs must re-check the requester's access when the
    job executes and before delivering the result.

## 5. Performance

-   Use indexes and bounded date ranges for common reports.
-   Use read models/materialized aggregates only when measurement shows
    a need.
-   Document freshness expectations for asynchronous reports.
-   Show the "as of" time for eventually consistent data.
-   Large exports should be asynchronous and cancellable.

## 6. Acceptance criteria

-   Report metrics match a curated test dataset.
-   Timezone edge cases, date boundaries, and daylight-saving
    transitions are tested where relevant.
-   Unauthorized projects do not contribute to visible totals.
-   Dashboard queries have agreed latency budgets and pagination/export
    limits.

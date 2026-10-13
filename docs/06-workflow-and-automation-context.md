# 06 --- BRD: Workflow and Automation Context

## 1. Context purpose

Define how work moves through states and how valid transitions can
trigger controlled follow-up actions.

## 2. Workflow model

A workflow definition contains: - ID and version. - Scope
(organization/project and optionally work-item type). - States with
stable IDs, names, categories, and ordering. - Allowed transitions. -
Optional transition guards and required fields. - Active/inactive
lifecycle metadata.

Suggested semantic state categories: `Backlog`, `Unstarted`, `Started`,
`Completed`, `Cancelled`. Display labels are configurable; reporting
uses semantic categories rather than matching labels like `"Done"`.

## 3. Transition policy

A transition request must validate: 1. Actor is authorized to update the
item. 2. Source state is current. 3. Target state exists in the
applicable workflow version. 4. Transition is allowed from source to
target. 5. Required fields/approvals are satisfied. 6. Any domain guard
passes. 7. The work item has not changed concurrently.

On success, persist the state change, transition history, and durable
event/outbox record in the same transaction when supported.

## 4. Workflow configuration rules

-   A workflow must have at least one non-terminal state and an agreed
    valid initial state.
-   Removing a state referenced by existing work items is forbidden
    until items are migrated.
-   A workflow update must not silently invalidate in-flight work.
-   Workflow versions should be immutable once published; edits create a
    new version or use a controlled migration.
-   A project manager may configure a project workflow only if
    authorized.
-   Do not execute arbitrary user-supplied code as a workflow condition.

## 5. Automation model (P1)

An automation rule has: - Trigger (e.g. state changed, due date
approaching, item created). - Conditions (allowlisted
field/operator/value combinations). - Actions (notify user, assign user,
add label, change state if permitted, send webhook). - Enabled flag,
creator, timestamps, and version. - Execution history and failure
details.

Rules: - Prevent infinite loops through idempotency, recursion limits,
and event/action correlation. - Automation executes asynchronously
unless the action is essential to the originating transaction. -
Re-evaluate permissions and domain guards at execution time. - Failures
are retried with bounded exponential backoff; permanent failures are
surfaced. - Secrets are stored in a secret manager or encrypted
configuration, never plaintext in logs. - Rule edits must be audited.

## 6. Approvals (P1)

Approval policy should define: - Required approver(s) or group. -
Whether self-approval is allowed. - Quorum/all-required behavior. -
Expiry/escalation. - Rejection behavior. - Whether approval applies to a
transition or an entity version.

Approval decisions must reference the exact item/configuration version
they reviewed. A material change after approval may invalidate it
according to policy.

## 7. Acceptance criteria

-   Invalid transitions are rejected consistently by API and UI.
-   Workflow edits preserve existing items or require an explicit
    migration.
-   Automation retries do not duplicate side effects.
-   Rule execution is observable and failures can be inspected by
    authorized administrators.
-   Approval bypass attempts are covered by negative tests.

# 03 --- BRD: Identity, Organization, Membership, and Access

## 1. Context purpose

Establish who a user is, which organization they belong to, and what
actions they may perform within a scope.

## 2. Capabilities

-   Sign in/out and manage sessions.
-   Invite, suspend, reactivate, and remove organization members.
-   Assign organization roles.
-   Manage organization settings and default policies.
-   Manage teams and team membership.
-   Resolve authorization for project and work-item operations.
-   Review security-relevant audit events.

## 3. Proposed roles

-   **Organization Owner:** Full control, including ownership transfer
    and organization deletion initiation.
-   **Organization Admin:** Manages settings, members, teams, and
    projects; cannot transfer ownership unless explicitly permitted.
-   **Member:** Can participate in projects according to project
    membership and visibility.
-   **Guest/External Viewer (later):** Restricted access only to
    explicitly shared resources.

Project-level roles: - **Project Admin/Manager:** Manages project
settings, members, and workflow configuration. - **Contributor:**
Creates and updates work items within allowed policies. - **Viewer:**
Read-only access.

Roles are starting presets confirmed in Phase 0: organization roles are
Owner, Admin, and Member (Guest/External Viewer later); project roles are
Admin, Contributor, and Viewer. Prefer explicit permissions/policies in
the authorization layer rather than scattering role-name comparisons
through the code. Team lead is a field on the Team aggregate, not an
organization role.

There is no workspace-level membership: organization membership is the
only tenancy-level membership, since an organization has exactly one
workspace in the MVP (`docs/adr/0003-single-workspace-tenancy.md`).

## 4. Authorization policy

Every use case must authorize the actor against: 1. Authentication
state. 2. Organization membership and active status. 3. Organization
boundary. 4. Resource visibility. 5. Project membership or a documented
organization-wide visibility rule. 6. Requested action. 7. Additional
domain constraints.

Default is deny. UI hiding is not authorization.

### Policy examples

-   Only an active organization owner/admin can invite members.
-   A suspended member cannot read or mutate organization resources.
-   A project manager can edit project settings only for projects they
    manage.
-   A viewer cannot mutate work items.
-   A user cannot assign a person who is not eligible for the project.
-   A user cannot access a work item by guessing its ID if the parent
    project is inaccessible.
-   Organization ownership cannot be removed without a successful
    ownership transfer.
-   Removing a member must not erase audit history or historical
    authorship.

## 5. Authentication requirements

-   Use a vetted authentication library/provider; do not invent password
    cryptography.
-   Store only password hashes when local passwords are supported.
-   Protect sessions/tokens against theft and replay according to the
    selected authentication model.
-   Rate-limit authentication and invitation endpoints.
-   Support revocation of sessions after account suspension or
    credential reset.
-   Avoid logging credentials, access tokens, invitation secrets, or
    reset tokens.
-   Consider OIDC/SSO as a later or deployment-specific requirement.

## 6. Membership lifecycle

Suggested states: `Invited`, `Active`, `Suspended`, `Removed`.

Rules: - Invitations expire after a configurable period. - Invitation
acceptance is single-use and bound to the intended email/account. -
Suspended users lose access immediately or within a clearly defined
cache invalidation window. - Removal revokes active sessions and access
grants. - Historical records preserve the actor's stable ID and may
display a tombstone label if the account is deleted. - Re-adding a user
must not silently restore revoked privileges unless policy says so.

## 7. Acceptance criteria

-   Cross-organization access tests cover read, write, search, export,
    and background jobs.
-   All protected use cases require an explicit actor.
-   Permission denial returns a consistent response without leaking
    whether a hidden resource exists when that would expose sensitive
    information.
-   Member suspension invalidates access according to the documented
    session strategy.
-   Role changes are audited.

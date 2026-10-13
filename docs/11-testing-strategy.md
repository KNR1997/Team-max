# 11 --- Testing Strategy (TDD)

## 1. Goal

Use tests to clarify expected behavior, protect domain invariants, and
enable safe refactoring. TDD means writing a failing test for a behavior
before implementing that behavior, then implementing the smallest
solution and refactoring while keeping tests green.

## 2. Test pyramid

### Domain unit tests

Fast, deterministic, no database/network/framework. Test: - Value-object
validation. - Entity behavior and state transitions. - Assignment and
hierarchy invariants. - Workflow transition guards. - Permission policy
decisions when expressed as pure rules. - Metric calculations.

### Application/use-case tests

Test orchestration with fake ports or focused test doubles: -
Authorization is invoked. - Correct repository operations occur. -
Transaction boundaries and domain events are handled. - Failure paths do
not report success. - Duplicate/retry behavior is safe.

### Integration tests

Use a real test database and relevant infrastructure where practical: -
Repository mappings and query filters. - Unique constraints and
concurrency behavior. - Transactions and outbox persistence. - API
authentication/authorization. - Queue consumer idempotency. - File
storage access rules.

### End-to-end tests

Cover a small number of critical user journeys: - Create project →
create work item → assign → transition → see activity. - Restricted
project remains invisible to a non-member. - Suspended user cannot
access protected resources. - Manager sees correct overdue dashboard
totals.

### Architecture tests

Automate rules such as: - Domain packages cannot import
HTTP/ORM/provider packages. - One context cannot access another
context's repositories/tables directly. - Controllers do not contain
business rules. - Dependency direction is enforced.

## 3. TDD workflow

For each use case: 1. Write a failing test describing one business
behavior. 2. Run it and confirm it fails for the intended reason. 3.
Implement the minimum behavior. 4. Run the focused test. 5. Refactor
names, boundaries, duplication, and design. 6. Run the full relevant
test suite. 7. Add integration tests for persistence/security behavior.
8. Record the acceptance criterion covered by the tests.

Do not write tests that merely mirror implementation details. Test
observable behavior and invariants.

## 4. Test naming

Prefer behavior-focused names: -
`CreateWorkItem_rejectsAssigneeOutsideProject` -
`ChangeState_rejectsTransitionNotDefinedByWorkflow` -
`ProjectQuery_doesNotReturnRestrictedProjectToNonMember` -
`AssignWorkItem_rejectsSuspendedUser`

Use Arrange--Act--Assert or Given--When--Then consistently.

## 5. Test data and isolation

-   Each test controls its own data and clock.
-   Avoid shared mutable fixtures and order-dependent tests.
-   Use factories/builders for valid domain objects.
-   Use a fake clock for date/time rules.
-   Use deterministic IDs where helpful.
-   Never use production data in ordinary test environments.
-   Clean up test resources and avoid tests depending on external
    network services.

## 6. Security testing matrix

For every resource, test: - Anonymous user. - User from another
organization. - Suspended member. - Organization member without project
access. - Project viewer attempting a write. - Contributor attempting
project configuration changes. - Authorized actor. - Background job
executing after access has been revoked.

Apply this matrix to detail, list, search, export, attachment,
notification, and mutation flows as applicable.

## 7. Coverage policy

Do not use a single line-coverage percentage as proof of quality.
Require: - Tests for all domain invariants. - Tests for authorization
policy branches. - Tests for error paths and concurrency-sensitive
operations. - Coverage trend monitoring and a justified baseline. - No
critical behavior left untested merely to meet a percentage target.

## 8. Definition of done for a feature

-   Acceptance criteria are testable and met.
-   Relevant domain and application tests pass.
-   Integration/security tests are added where needed.
-   Lint, formatting, type checks, and static analysis pass.
-   API/schema and docs are updated.
-   Logs and metrics are useful but do not leak sensitive data.
-   Migration and rollback implications are reviewed.
-   No known critical security issue is introduced.

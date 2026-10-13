# 12 --- Clean Code and DDD Engineering Standards

## 1. General principles

-   Prefer clear names over clever names.
-   Keep functions small and focused on one intention.
-   Make invalid states difficult to represent.
-   Prefer explicit dependencies over hidden global state.
-   Separate policy decisions from mechanics.
-   Avoid premature abstractions; introduce abstractions when they
    clarify a stable concept or isolate a volatile dependency.
-   Remove dead code and obsolete comments.
-   Comments should explain why, constraints, or non-obvious
    decisions---not restate the code.
-   Avoid boolean parameters that radically change a function's
    behavior; use explicit methods or options objects.
-   Avoid broad exception catches that hide failure.
-   Keep error handling at boundaries and translate errors once.

## 2. DDD rules

-   Use ubiquitous language from the context map.
-   Entities have identity; value objects are defined by their values
    and should be immutable where practical.
-   Aggregates enforce invariants through behavior, not public field
    mutation.
-   External code should modify an aggregate through intention-revealing
    methods.
-   Repository interfaces represent aggregate persistence, not arbitrary
    table CRUD.
-   Domain services are for domain behavior that does not naturally
    belong to one entity/value object.
-   Application services coordinate use cases; they should not become a
    second domain model.
-   Domain events are named in past tense and represent facts that
    occurred.
-   Avoid an anemic domain model for rules with meaningful invariants,
    but do not force every data structure into an entity.
-   Do not create a generic base class or repository abstraction without
    a concrete need.

## 3. Dependency and layer rules

-   Domain must not import infrastructure or interface frameworks.
-   Application may depend on domain and declared ports.
-   Infrastructure implements ports.
-   Controllers/handlers validate transport shape and call application
    use cases.
-   Controllers should not query ORM models directly for business
    operations.
-   Repositories should not contain workflow or authorization policy.
-   Authorization is explicit and testable; do not rely on obscurity or
    frontend hiding.
-   Cross-context calls go through a published application interface or
    events.

## 4. Naming conventions

-   Use business terms, not database abbreviations.
-   Commands use verbs: `AssignWorkItem`, `ArchiveProject`.
-   Queries use `Get`, `List`, or `Search`.
-   Policies read as decisions: `CanEditProject`, `CanAssignWorkItem`.
-   Domain events use past tense: `WorkItemAssigned`.
-   Avoid ambiguous names like `Manager`, `Helper`, `Processor`,
    `CommonService`, or `Utils` without a precise responsibility.
-   Boolean names should be readable as questions or states,
    e.g. `isArchived`, `canTransition`.

## 5. Validation and invariants

-   Boundary validation checks syntax, shape, size, and required fields.
-   Domain validation checks business meaning and invariants.
-   Database constraints protect uniqueness and referential integrity.
-   Avoid duplicating complex business rules in UI and API. UI may
    mirror rules for usability, but server/domain is authoritative.
-   Return typed, actionable errors rather than magic strings.
-   Never use silent fallback values for required business decisions.

## 6. Refactoring and duplication

-   Duplicate knowledge is more dangerous than duplicated lines.
-   Do not prematurely abstract two similar flows if their business
    rules differ.
-   When extracting shared logic, keep it in the context that owns the
    rule or in a deliberately small shared kernel.
-   Refactor with tests in place.
-   Prefer composition over deep inheritance.
-   Avoid "god objects" and services that own unrelated
    responsibilities.

## 7. Code review checklist

-   Does the change match the BRD and ubiquitous language?
-   Are invariants enforced in the domain?
-   Is authorization checked on every relevant path?
-   Are transactions and concurrency considered?
-   Are side effects isolated and retry-safe?
-   Are errors stable and safe?
-   Are tests behavior-focused and deterministic?
-   Are context boundaries respected?
-   Are logs free of secrets and sensitive payloads?
-   Are docs, migrations, and operational implications updated?

## 8. Anti-patterns to avoid

-   One service class for the whole application.
-   Generic CRUD that bypasses domain behavior.
-   Domain entities importing ORM/framework packages.
-   Controllers with business decisions.
-   Repositories calling other repositories across contexts.
-   One transaction spanning unrelated contexts without a documented
    consistency model.
-   Direct external API calls inside a database transaction.
-   Fire-and-forget goroutines/jobs without lifecycle, retries, or
    observability.
-   Mocking every dependency so tests only confirm that methods were
    called.
-   Using tests only after implementation instead of using TDD for core
    behavior.

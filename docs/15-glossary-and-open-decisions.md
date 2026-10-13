# 15 --- Glossary and Open Decisions

## Glossary

-   **Aggregate:** Cluster of domain objects treated as one consistency
    boundary.
-   **Aggregate root:** The only object through which an aggregate is
    modified externally.
-   **Bounded context:** A boundary within which a domain model and
    language have a specific meaning.
-   **Command:** Request to change state by expressing intent.
-   **Domain event:** A fact that occurred in the domain.
-   **Entity:** Object whose identity persists through changes.
-   **Integration event:** Versioned event contract for another
    context/system.
-   **Invariant:** Rule that must always hold true.
-   **Port:** An interface describing a dependency required by
    application/domain logic.
-   **Adapter:** Implementation of a port for a particular technology.
-   **Repository:** Abstraction for retrieving and persisting
    aggregates.
-   **TDD:** Test-Driven Development: red, green, refactor.
-   **Ubiquitous language:** Shared domain vocabulary used by
    stakeholders and code.
-   **Value object:** Object defined by its values rather than identity.
-   **Outbox pattern:** Persist business change and event record
    atomically, then publish asynchronously.
-   **Optimistic concurrency:** Detect conflicting updates using a
    version/token rather than silently overwriting changes.

## Open decisions requiring stakeholder approval

  -----------------------------------------------------------------------
  ID                      Decision                Suggested default
  ----------------------- ----------------------- -----------------------
  DEC-001                 Organization versus     Model both, initially
                          workspace hierarchy     permit one workspace
                                                  per organization if
                                                  simpler

  DEC-002                 Project visibility      Restricted by default;
                                                  organization-visible
                                                  must be explicit

  DEC-003                 Can contributors create No, unless granted a
                          projects?               permission

  DEC-004                 Can assignee be empty?  Yes, but dashboard
                                                  should surface
                                                  unassigned active work

  DEC-005                 Can users delete work   Archive by default;
                          items?                  hard delete restricted

  DEC-006                 Project key mutability  Immutable after
                                                  creation for MVP

  DEC-007                 Workflow versioning     Publish new version and
                                                  explicitly migrate
                                                  affected items

  DEC-008                 Authentication          Use vetted local auth
                                                  or OIDC based on
                                                  company needs

  DEC-009                 Background jobs         Start with a reliable
                                                  worker/queue adapter;
                                                  select broker based on
                                                  operational needs

  DEC-010                 Search                  Start with database
                                                  capabilities; add
                                                  dedicated search only
                                                  when needed

  DEC-011                 File storage            Private S3-compatible
                                                  storage or equivalent

  DEC-012                 Audit retention         Define with
                                                  legal/security owner
                                                  before production

  DEC-013                 RPO/RTO                 Define with business
                                                  owner; test recovery

  DEC-014                 API response envelope   Standardize one
                                                  response/error contract

  DEC-015                 Time tracking and       Post-MVP unless
                          estimates               required by pilot teams
  -----------------------------------------------------------------------

## Change management

When a decision changes: 1. Update this table and relevant BRD. 2.
Explain the reason and impact. 3. Create/update an ADR for architectural
consequences. 4. Update domain model and acceptance tests. 5. Plan data
migration/backward compatibility.

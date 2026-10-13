# 02 --- Domain-Driven Design: Context Map

## 1. Purpose

Define bounded contexts, their responsibilities, ownership, and
integration contracts. Context boundaries are logical boundaries first;
the initial deployment may remain a modular monolith.

## 2. Proposed bounded contexts

  ------------------------------------------------------------------------
  Context                 Owns                    Must not own
  ----------------------- ----------------------- ------------------------
  Identity & Access       Authentication          Project workflow or task
                          identity,               lifecycle
                          sessions/tokens, global 
                          user identity,          
                          authorization policy    
                          primitives              

  Organization &          Organizations,          Work-item business rules
  Membership              workspace settings,     
                          memberships,            
                          organization roles      

  Teams                   Teams, team membership, Organization identity
                          team ownership metadata records

  Project Management      Project lifecycle,      Work-item transition
                          project identifiers,    logic beyond
                          project configuration,  project-level policy
                          project membership,     
                          visibility              

  Work Management         Work items, hierarchy,  User
                          labels, priorities, due authentication/session
                          dates, estimates,       management
                          relationships           

  Workflow                Workflow definitions,   Direct persistence of
                          states, transitions,    work-item aggregates
                          transition guards and   
                          configuration           

  Planning                Cycles, modules,        Core work-item ownership
                          milestones, releases    
                          and planning membership 

  Collaboration           Comments, attachments   Authorization source of
                          metadata, reactions,    truth
                          mentions                

  Notification            Notification            Business transaction
                          preferences, inbox,     ownership
                          delivery attempts and   
                          templates               

  Audit & Activity        Append-only audit       Mutating source
                          records and             aggregates
                          human-readable activity 
                          history                 

  Reporting               Read models, metrics,   Authoritative writes to
                          dashboard queries and   operational aggregates
                          exports                 

  Integration             Webhooks, external      Owning core business
                          connections,            entities
                          idempotency and         
                          delivery status         

  Knowledge (later)       Pages, collections,     Project lifecycle
                          links to projects/work  
                          items                   
  ------------------------------------------------------------------------

## 3. Context map

``` mermaid
flowchart TD
    IA[Identity & Access] --> OM[Organization & Membership]
    OM --> TM[Teams]
    OM --> PM[Project Management]
    TM --> PM
    PM --> WM[Work Management]
    WF[Workflow] --> WM
    PM --> PL[Planning]
    WM --> CO[Collaboration]
    WM --> AU[Audit & Activity]
    PM --> AU
    WM --> NO[Notification]
    AU --> RP[Reporting]
    WM --> RP
    PM --> IN[Integration]
    WM --> IN
    PM --> KN[Knowledge - later]
    WM --> KN
```

This diagram shows conceptual dependencies, not direct database joins.
Prefer published application interfaces or events over another context
reaching into a context's tables.

## 4. Shared language

Use these terms consistently in code, API docs, tests, and product UI:

-   `OrganizationId`, `UserId`, `TeamId`, `ProjectId`, `WorkItemId`
-   `ProjectKey` or `ProjectIdentifier`: human-readable project code
    used in work-item references
-   `WorkItemKey`: stable human-readable key such as `APP-123`
-   `WorkflowStateId`: identity of a configured state, not a hard-coded
    string such as `"done"`
-   `Actor`: authenticated principal performing an action
-   `Membership`: relationship that grants scoped participation; not
    equivalent to authentication
-   `DomainEvent`: internal fact emitted after a valid domain change
-   `IntegrationEvent`: versioned message intended for another context
    or external consumer

Avoid overloaded terms such as "user role" without specifying
organization role, project role, or workflow responsibility.

## 5. Aggregate boundaries

Proposed initial aggregates:

-   `Organization` --- organization identity and lifecycle invariants.
-   `Team` --- team details and membership rules, or team membership as
    a separate aggregate if concurrency requires it.
-   `Project` --- project identity/configuration and project membership
    policy.
-   `WorkItem` --- title, description reference, state, priority,
    assignment, due date, parent/child references, and domain methods.
-   `WorkflowDefinition` --- states, transition rules, and configuration
    version.
-   `Cycle` --- date range and planning rules.
-   `Module` --- grouping metadata and work-item links.
-   `NotificationPreference` --- per-user delivery preferences.
-   `WebhookSubscription` --- endpoint configuration and signing secret
    metadata.

Do not create one giant aggregate containing an entire project and all
work items. Aggregates should protect invariants within a small
consistency boundary.

## 6. Cross-context rules

1.  A context owns its tables and migration lifecycle.
2.  Other contexts must not write directly to those tables.
3.  Cross-context references use IDs and explicit interfaces/events.
4.  Do not call another context's repository directly.
5.  Application services may orchestrate multiple contexts through
    exposed application ports; avoid service-to-service chains that hide
    transaction boundaries.
6.  If a business transaction crosses contexts, explicitly document
    consistency requirements and failure handling.
7.  Use an outbox for durable event publication where losing an event
    would violate business expectations.
8.  Reporting may use denormalized read models; it must not become the
    write source of truth.
9.  Avoid generic "common" modules that become a dumping ground. Shared
    kernel types must be small and stable.

## 7. Dependency direction

Preferred direction:

`Interface adapters → Application use cases → Domain model`

Infrastructure implements ports required by application/domain layers.
Domain code must not import web frameworks, ORM models, message brokers,
or provider SDKs.

A context may contain: - `domain/`: entities, value objects, domain
services, domain events, repository interfaces where appropriate. -
`application/`: use cases, commands/queries, DTOs, ports, transaction
orchestration. - `infrastructure/`: ORM mappings, repositories, broker
adapters, email providers, external clients. - `interfaces/`: HTTP
handlers/controllers, serializers, CLI, event consumers.

Adjust folder names to the chosen language, but preserve the dependency
rule.

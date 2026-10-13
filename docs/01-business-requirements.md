# 01 --- Business Requirements Document (BRD)

## 1. Document control

-   **Status:** Draft v0.1
-   **Business owner:** To be assigned
-   **Technical owner:** To be assigned
-   **Approval:** Product owner and engineering lead
-   **Review cadence:** At each major release

## 2. Business problem

Work is often distributed across chat, spreadsheets, documents, and
individual task tools. This causes unclear ownership, inconsistent
processes, weak auditability, duplicated reporting, and limited
visibility into delivery risks.

The platform should centralize work tracking and provide configurable,
permission-aware project execution.

## 3. Business outcomes

-   Shared visibility of projects and responsibilities.
-   Reduced manual status-report preparation.
-   More consistent work handoffs and approvals.
-   Better detection of blocked and overdue work.
-   A maintainable system that can grow with organizational needs.

## 4. Functional requirements

Priority: **P0** required for MVP, **P1** important next phase, **P2**
later enhancement.

  -----------------------------------------------------------------------
  ID                      Requirement             Priority
  ----------------------- ----------------------- -----------------------
  BR-001                  Users can authenticate  P0
                          and manage their        
                          profile.                

  BR-002                  Administrators can      P0
                          manage organization     
                          members and roles.      

  BR-003                  Users can create teams  P0
                          and manage team         
                          membership subject to   
                          permission.             

  BR-004                  Authorized users can    P0
                          create projects with a  
                          unique identifier.      

  BR-005                  Projects have explicit  P0
                          visibility and          
                          membership rules.       

  BR-006                  Users can create,       P0
                          update, assign, search, 
                          and filter work items.  

  BR-007                  Work items support      P0
                          priority, due date,     
                          labels, descriptions,   
                          comments, and           
                          attachments metadata.   

  BR-008                  Project workflows can   P0
                          define states and       
                          allowed transitions.    

  BR-009                  Important changes are   P0
                          recorded in an          
                          activity/audit history. 

  BR-010                  Users receive           P0
                          notifications for       
                          configured important    
                          events.                 

  BR-011                  Managers can view       P0
                          project progress and    
                          overdue work.           

  BR-012                  Work items can be       P0
                          linked as parent/child  
                          and related items.      

  BR-013                  Users can save          P1
                          frequently used filters 
                          as views.               

  BR-014                  Teams can plan work in  P1
                          cycles and modules.     

  BR-015                  Administrators can      P1
                          configure automation    
                          rules.                  

  BR-016                  The platform exposes    P1
                          documented APIs and     
                          webhooks.               

  BR-017                  Users can create        P2
                          project knowledge pages 
                          linked to work.         

  BR-018                  Managers can analyze    P2
                          cross-project workload  
                          and delivery trends.    
  -----------------------------------------------------------------------

## 5. Non-functional requirements

-   **Security:** Enforce authorization on the server for every read and
    write; do not trust client-side visibility.
-   **Isolation:** No user can access data from an organization they do
    not belong to.
-   **Consistency:** Domain invariants are enforced transactionally
    where required.
-   **Availability:** Target 99.5% monthly availability after production
    stabilization, excluding agreed maintenance.
-   **Performance:** Establish baseline under representative load.
    Initial target: p95 under 500 ms for common non-reporting API
    requests, excluding network latency and external dependencies.
-   **Accessibility:** Target WCAG 2.2 AA for key workflows where
    feasible.
-   **Observability:** Structured logs, request/correlation IDs,
    metrics, health/readiness checks, and actionable error reporting.
-   **Maintainability:** Bounded contexts, automated tests, documented
    decisions, and enforced dependency rules.
-   **Recovery:** Daily backups at minimum; define and test RPO/RTO with
    the business before production.
-   **Privacy:** Collect only necessary personal data and define
    retention rules.

## 6. Business rules

1.  Every project belongs to exactly one organization and has a unique
    identifier within the defined scope.
2.  Every work item belongs to exactly one project.
3.  A user may only be assigned to a work item if they are an eligible
    member of that project's organization and satisfy project assignment
    policy.
4.  A work item state must be valid for the project's workflow
    configuration.
5.  Workflow transitions must be validated on the server.
6.  Deletion of referenced business records should generally be
    soft-delete/archive or blocked; destructive deletion requires a
    deliberate policy.
7.  Cross-organization relationships are forbidden unless a future,
    explicit sharing feature defines them.
8.  Audit records must not be editable by ordinary users.
9.  All timestamps are stored in UTC.
10. Reports must define what "completed", "overdue", "blocked", and
    "active" mean.

## 7. User journeys

### Create and track a task

1.  User opens a project they can access.
2.  User creates a work item with title and required fields.
3.  System validates membership, permissions, and project configuration.
4.  System persists the item and emits a `WorkItemCreated` domain event.
5.  Activity history is recorded and relevant subscribers are notified.

### Move a task through workflow

1.  User requests a state transition.
2.  Application loads the work item and relevant workflow configuration.
3.  Policy checks permission and domain rules validate the transition.
4.  State change, history record, and outbox event are committed
    atomically where practical.
5.  Notifications/automation execute asynchronously.

### Review project health

1.  Manager opens a project dashboard.
2.  System applies authorization before querying data.
3.  Dashboard reports counts and trends using documented metric
    definitions.
4.  User can navigate from an aggregate metric to the underlying
    authorized items.

## 8. Acceptance criteria for MVP

-   A new organization can be configured without direct database edits.
-   Unauthorized access tests cover list, detail, search, export, and
    mutation endpoints.
-   Workflow state changes reject invalid transitions with a stable
    error response.
-   Audit/activity history captures actor, action, target, timestamp,
    and relevant change metadata.
-   Notifications are retryable and do not block core work-item
    creation.
-   Reports reconcile with underlying work-item data for agreed test
    fixtures.
-   Backup restoration and a basic incident runbook are demonstrated.

## 9. Open decisions

-   Is one organization allowed to contain multiple workspaces in the
    first release?
-   Can a project be visible to all organization members, or only
    explicit members?
-   Is a work item required to have an assignee?
-   What is the policy for deleting users who own historical work?
-   Which identity provider/SSO is required?
-   What RPO/RTO and retention periods are acceptable?

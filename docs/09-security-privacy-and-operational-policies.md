# 09 --- Security, Privacy, and Operational Policies

## 1. Security baseline

-   Deny by default; enforce authorization server-side.
-   Validate all external input at the boundary.
-   Use parameterized queries/ORM-safe APIs.
-   Apply CSRF protection where cookie-based authentication requires it.
-   Configure secure cookies, TLS, security headers, and appropriate
    CORS.
-   Rate-limit login, invitations, password reset, search, and expensive
    exports.
-   Protect against XSS, SSRF, injection, insecure direct object
    references, and mass assignment.
-   Keep secrets outside source control and ordinary configuration
    files.
-   Pin and review dependencies; automate vulnerability scanning.
-   Separate production, staging, and development credentials and data.
-   Do not expose stack traces or internal infrastructure details to
    clients.

## 2. Data classification

Proposed categories: - **Public:** Approved for public disclosure. -
**Internal:** Ordinary company work data. - **Confidential:** Restricted
project, customer, or business information. - **Sensitive:**
Credentials, tokens, security secrets, or legally restricted data.

Default project data is Internal. Confidential projects require
restricted membership. Secrets must never be stored in work-item
descriptions or ordinary metadata.

## 3. Privacy and retention

-   Collect only data required for operation.
-   Define retention periods for audit records, notifications,
    attachments, and deleted accounts.
-   Provide a documented account deactivation/removal process.
-   Preserve audit integrity without unnecessarily retaining personal
    content.
-   Exports and backups are subject to the same protection and retention
    expectations as source data.
-   Privacy/legal requirements must be reviewed for the jurisdictions in
    which the organization operates.

## 4. Audit policy

Record at minimum: - Authentication and account security events. -
Membership/role changes. - Project visibility and permission changes. -
Workflow configuration changes. - Work-item state, assignment, due-date,
and archive changes. - Export, webhook configuration, and administrative
actions.

Audit records should be append-only to ordinary application roles.
Restrict access to audit search/export. Include a correlation ID but
avoid tokens or secrets.

## 5. Reliability and recovery

-   Define RPO and RTO with the business.
-   Back up database and object storage consistently.
-   Encrypt backups at rest and restrict access.
-   Test restoration regularly; a successful backup job alone is
    insufficient.
-   Document migration/rollback strategy and release procedure.
-   Health endpoints should distinguish liveness from readiness.
-   Use graceful shutdown and bounded retries for background workers.
-   Add idempotency for retried jobs and external webhook deliveries.

## 6. Observability

-   Structured logs with timestamp, severity, service, environment,
    request ID, actor ID where safe, and error category.
-   Metrics for request latency, error rate, database health, queue
    depth, retry counts, and notification delivery.
-   Traces/correlation across API and workers where supported.
-   Alert on actionable symptoms rather than every transient error.
-   Redact personal and confidential payloads from logs.

## 7. Incident response

Maintain runbooks for: - Suspected account compromise. -
Cross-organization data exposure. - Failed deployment/migration. -
Database restore. - Queue backlog or notification failure. - Storage
outage. - Secret rotation.

Every incident should have an owner, timeline, impact assessment,
remediation, and follow-up actions.

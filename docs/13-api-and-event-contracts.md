# 13 --- API and Event Contract Standards

## 1. HTTP API conventions

-   Use resource-oriented paths and standard HTTP methods.
-   Use stable UUIDs for machine identity and human-readable keys for
    user-facing references.
-   Version APIs when backward-incompatible changes are introduced.
-   Use ISO 8601 timestamps in UTC.
-   Validate query parameters and enforce maximum page sizes.
-   Allowlist sort fields and filter operators.
-   Return consistent errors with a stable code, safe message, and
    optional field errors.
-   Include a request/correlation ID in responses and logs.
-   Use ETags/version numbers or equivalent optimistic concurrency for
    mutable resources where useful.

Example paths: - `GET /api/v1/projects` - `POST /api/v1/projects` -
`GET /api/v1/projects/{projectId}/work-items` -
`POST /api/v1/projects/{projectId}/work-items` -
`POST /api/v1/work-items/{workItemId}/transitions` -
`GET /api/v1/work-items/{workItemId}/activity`

## 2. Example error response

``` json
{
  "success": false,
  "error": {
    "code": "WORKFLOW_TRANSITION_NOT_ALLOWED",
    "message": "The requested state transition is not allowed.",
    "fields": []
  },
  "requestId": "correlation-id"
}
```

Do not return stack traces, SQL messages, secret values, or internal
service topology.

## 3. Domain events versus integration events

Domain events are internal to a context/application model. Integration
events are stable, versioned contracts intended for asynchronous
consumers or external systems. Do not expose internal domain class
serialization as a public event schema.

Example integration event:

``` json
{
  "eventId": "uuid",
  "eventType": "work_item.state_changed",
  "schemaVersion": 1,
  "occurredAt": "2026-01-01T12:00:00Z",
  "organizationId": "uuid",
  "actorId": "uuid",
  "aggregate": {
    "type": "work_item",
    "id": "uuid",
    "version": 4
  },
  "data": {
    "projectId": "uuid",
    "workItemKey": "OPS-104",
    "fromStateId": "uuid",
    "toStateId": "uuid"
  }
}
```

## 4. Event delivery rules

-   Use an outbox for events that must not be lost after a database
    commit.
-   Delivery is at-least-once unless the chosen infrastructure proves
    otherwise.
-   Consumers deduplicate by `eventId` or a consumer-specific
    idempotency key.
-   Consumers must tolerate duplicate and out-of-order events where
    possible.
-   Include schema version and define compatibility policy.
-   Use bounded retries and a dead-letter/manual-replay process.
-   Sign outbound webhooks and support secret rotation.
-   Do not include confidential descriptions or comments in generic
    events by default.

## 5. API compatibility

-   Additive optional fields are generally backward compatible.
-   Renaming/removing fields, changing meaning, or tightening validation
    can be breaking.
-   Document deprecation windows.
-   Contract tests should validate representative API schemas and event
    consumers.

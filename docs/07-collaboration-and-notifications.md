# 07 --- BRD: Collaboration, Notifications, and Activity

## 1. Context purpose

Support discussions and timely coordination while preserving access
boundaries and historical context.

## 2. Comments and mentions

-   Authorized users can add, edit, and delete their comments according
    to policy.
-   Comment policy (Phase 0): authors edit and delete their own
    comments; a project admin may delete a comment for moderation;
    deletion is soft — content is scrubbed and a tombstone plus the
    minimum audit evidence required by retention policy is kept.
-   Edited comments show an edited marker and timestamp.
-   Mentioned users must be eligible to view the referenced resource.
-   Mention parsing must not create notifications to unauthorized users.
-   Rich text content must be sanitized before rendering.

## 3. Attachments

-   Store files in private object storage or an equivalent controlled
    storage service.
-   Store metadata and ownership references in the application database.
-   Enforce authorization before upload, download, preview, or deletion.
-   Use size/type limits and malware scanning where available.
-   Do not trust filename extensions or client-provided MIME types.
-   Use short-lived signed URLs if direct object-storage downloads are
    used.

## 4. Notification channels

MVP: - In-app notification inbox. - Email notifications for selected
events.

Later: - Slack/Teams integrations, digest emails, push notifications.

Notification preferences should support event type, channel, and
frequency where appropriate. Security and mandatory system notifications
may not be suppressible.

## 5. Delivery semantics

-   Notification creation should not block a core work-item transaction.
-   Use a durable queue/outbox pattern for important events.
-   Delivery is at-least-once where retries are used; consumers must be
    idempotent.
-   Record attempts, provider response category, retry schedule, and
    final status.
-   Do not log full email bodies if they contain sensitive project data.
-   A notification must never reveal a title, comment, or link to a user
    who no longer has access.

## 6. Activity history versus audit log

Activity history is user-facing and explains project/work-item changes.
Audit log is security/governance evidence. They may share event sources
but should have separate retention, access, and content policies.

Audit records should include actor, action, target type/ID, timestamp,
request/correlation ID, and safe before/after metadata when justified.
Do not store secrets or entire sensitive payloads.

## 7. Acceptance criteria

-   Mentions never notify unauthorized recipients.
-   Notification retries do not create uncontrolled duplicate inbox
    entries.
-   Revoked access is respected when notification content is generated
    and when links are opened.
-   Attachment access is tested for authorized and unauthorized users.
-   Audit log is append-only for ordinary application users.

---
status: accepted
date: 2026-10-20
---

# Private S3-compatible object storage for attachments

Attachment binaries live in private S3-compatible object storage; the application database stores only metadata and ownership references. Downloads and uploads go through authorization-checked endpoints that mint short-lived signed URLs, so object URLs never grant lasting access. Malware scanning is deferred past MVP, but the storage side sits behind an adapter so a scanner can be inserted into the upload path without changing collaboration code.

**Considered options:** filesystem storage on the app host (rejected: no clean path to backups, retention, or horizontal growth, and it couples file lifecycle to the deployable); public-bucket URLs with long-lived tokens (rejected: contradicts "enforce authorization before download" in `docs/07` §3).

**Consequences:** backups must cover object storage consistently with the database (`docs/09` §5); retention and deletion of attachments follow the same policy as other stored data; size/type limits are enforced at the boundary, never trusting client MIME types.

---
status: accepted
date: 2026-10-20
---

# Postgres outbox and polling worker; no broker for MVP

Background work (notifications, email, exports, scheduled reminders) runs from a transactional outbox table in the same Postgres database, drained by a polling worker in the same deployable. Jobs carry stable IDs, correlation IDs, bounded retries, deadlines, and idempotency keys; permanently failed jobs land in a dead-letter state with an operational inspection path. No separate message broker is introduced until measurement shows the polling loop or the database cannot carry the load.

**Considered options:** Redis/RabbitMQ/NATS from day one (rejected: extra operational surface and a second consistency boundary for an internal tool with modest volume); in-process fire-and-forget goroutines/threads (rejected: events lost on crash, violating the durability expectation behind the outbox pattern in `docs/13` §4).

**Consequences:** delivery is at-least-once and consumers must stay idempotent; notification latency is bounded by the poll interval; the outbox dispatcher sits behind a port so a broker can replace it without touching use cases.

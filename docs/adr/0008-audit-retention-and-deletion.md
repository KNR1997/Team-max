---
status: accepted
date: 2026-10-20
---

# Append-only audit log; archive by default; tombstoned users

The audit log is append-only and not editable or deletable by ordinary application roles; its retention period is set with the legal/security owner before production (tracked as a Phase 4 gate, not fixed here). Business records follow an archive-over-delete policy: work items are archived, not hard-deleted, and no ordinary hard-delete path exists for them in MVP. Deleted users are tombstoned — the stable ID and historical authorship are preserved, content displays a tombstone label, and deletion is blocked while the user still owns active work.

**Considered options:** hard delete with cascade (rejected: erases the "every important state change can be explained through history" guarantee in `docs/00` §4, and `docs/03` §4 forbids erasing historical authorship); full anonymization that rewrites actor references (rejected for MVP: heavier than a tombstone and unnecessary while stable IDs already decouple identity from the person).

**Consequences:** retention numbers are a policy input, not an engineering default — Phase 4 cannot ship without them; archive flags must be honored by every list, search, and reporting query; the user-deletion flow is a domain command with invariants, not a DELETE statement.

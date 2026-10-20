---
status: proposed
date: 2026-10-20
---

# Modular monolith over microservices

We are building Team-max as a modular monolith — one deployable whose internal module boundaries follow the bounded contexts in `docs/02`, with ports/adapters at every infrastructure seam. Chosen because transactions over core invariants (work item + workflow + audit + outbox) stay simple, operational cost stays low for an internal tool, and module boundaries can later guide service extraction if measurement justifies it.

**Considered options:** microservices from day one (rejected: no measured scaling or team-split reason, and distributed transactions would weaken invariants); an unstructured single app with the intention of restructuring later (rejected: dependency rot makes context boundaries unrecoverable).

**Status:** proposed — awaiting named technical-owner sign-off before Phase 1 locks the dependency rules. See `docs/10` §1.

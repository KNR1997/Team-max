---
status: accepted
date: 2026-10-20
---

# Database search for MVP

Search over work items and projects uses Postgres full-text and trigram indexing for the MVP. A dedicated search engine is added only when measurement shows a relevance or scale need that the database cannot meet.

**Considered options:** OpenSearch/Meilisearch/Elasticsearch from day one (rejected: a second datastore, an index-sync problem, and an extra failure mode for an internal deployment whose corpus is small; it also complicates the authorization story, since results must be filtered per-project anyway).

**Consequences:** search queries must apply authorization before returning rows *or* aggregates (`docs/05` §7) — the index is not a bypass; case-sensitivity and partial-match behavior are defined explicitly in the API contract; swapping engines later is confined to the search adapter.

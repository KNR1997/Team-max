---
status: accepted
date: 2026-10-20
---

# Immutable workflow versions with explicit migration

Published workflow definitions are immutable. Editing a project's workflow creates a new version; existing work items keep serving on the version they were created under until an explicit migration moves them, decided case by case. Removing a state that existing items reference is blocked until those items are migrated. A workflow update can therefore never silently invalidate in-flight work.

**Considered options:** mutable definitions with automatic remapping of affected items (rejected: "automatic" means the system guesses where open work goes, and a wrong guess corrupts reporting and history with no audit trail of the choice); blocking all workflow edits once items exist (rejected: teams legitimately need to evolve their process).

**Consequences:** workflow configuration is a two-step operation (publish version, migrate items); reporting must always resolve items against their own workflow version; migration is itself an audited action.

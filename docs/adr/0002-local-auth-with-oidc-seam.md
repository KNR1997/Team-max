---
status: accepted
date: 2026-10-20
---

# Local authentication behind an OIDC-ready port

MVP ships email + password authentication using a vetted library (Argon2id hashing, server-side session records delivered as httpOnly cookies), wrapped behind a port so an OIDC provider can be added later without touching domain or application code. The business has no mandated identity provider, so no IdP integration is built now.

**Considered options:** OIDC/SSO only from day one (rejected: no mandated IdP, and it would block pilots on infrastructure the org has not chosen); stateless JWT for the API (rejected for now: server-side sessions make revocation on suspension immediate, which `docs/03` §6 requires; can be revisited if a stateless consumer appears).

**Consequences:** cookie-based auth means CSRF protection is mandatory (`docs/09` §1); sessions must be revocable on suspend/removal; the OIDC seam is a real interface, not a comment — the auth adapter is replaceable.

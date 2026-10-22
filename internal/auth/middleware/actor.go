package middleware

import (
	"net/http"

	"todo/internal/shared/kernel"
)

// Actor injects a test/dev actor into context if present.
// In production this would validate sessions/tokens.
func Actor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// For now, inject actor from headers if present (dev only)
		orgID := r.Header.Get("X-Org-ID")
		userID := r.Header.Get("X-User-ID")
		wsID := r.Header.Get("X-Workspace-ID")
		if orgID != "" && userID != "" {
			actor := &kernel.Actor{
				UserID:         kernel.UserID(userID),
				OrganizationID: kernel.OrganizationID(orgID),
				WorkspaceID:    kernel.WorkspaceID(wsID),
			}
			next.ServeHTTP(w, r.WithContext(kernel.WithActor(r.Context(), actor)))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RequireAuth ensures actor is present.
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := kernel.ActorFromContext(r.Context()); !ok {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

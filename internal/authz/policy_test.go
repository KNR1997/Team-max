package authz

import (
	"testing"

	"todo/internal/shared/kernel"

	"github.com/stretchr/testify/require"
)

func TestCrossOrgAccessDenied(t *testing.T) {
	p := New()
	actor := &kernel.Actor{UserID: "u1", OrganizationID: "org-a", WorkspaceID: "ws-a"}
	err := p.AuthorizeProjectRead(actor, "org-b")
	require.ErrorIs(t, err, kernel.ErrForbidden)
}

func TestSameOrgAllowed(t *testing.T) {
	p := New()
	actor := &kernel.Actor{UserID: "u1", OrganizationID: "org-a", WorkspaceID: "ws-a"}
	err := p.AuthorizeProjectRead(actor, "org-a")
	require.NoError(t, err)
}

func TestMissingActorUnauthorized(t *testing.T) {
	p := New()
	err := p.AuthorizeProjectRead(nil, "org-a")
	require.ErrorIs(t, err, kernel.ErrUnauthorized)
}

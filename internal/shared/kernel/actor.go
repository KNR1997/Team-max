package kernel

// Actor represents the authenticated principal performing an action.
type Actor struct {
	UserID         UserID
	OrganizationID OrganizationID
	WorkspaceID    WorkspaceID
	SessionID      string
}

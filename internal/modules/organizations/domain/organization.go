package domain

import (
	"time"

	"todo/internal/shared/kernel"
)

type Organization struct {
	ID        kernel.OrganizationID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Workspace struct {
	ID             kernel.WorkspaceID
	OrganizationID kernel.OrganizationID
	Name           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Membership struct {
	ID             kernel.OrganizationID
	OrganizationID kernel.OrganizationID
	UserID         kernel.UserID
	Role           string
	Status         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

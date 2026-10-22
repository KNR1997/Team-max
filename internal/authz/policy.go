package authz

import "todo/internal/shared/kernel"

// Policy enforces deny-by-default authorization.
type Policy struct{}

func New() *Policy { return &Policy{} }

// AuthorizeProjectRead denies cross-org access by default.
func (p *Policy) AuthorizeProjectRead(actor *kernel.Actor, projectOrgID kernel.OrganizationID) error {
	if actor == nil {
		return kernel.ErrUnauthorized
	}
	if actor.OrganizationID != projectOrgID {
		return kernel.ErrForbidden
	}
	return nil
}

// AuthorizeProjectWrite denies cross-org access by default.
func (p *Policy) AuthorizeProjectWrite(actor *kernel.Actor, projectOrgID kernel.OrganizationID) error {
	if actor == nil {
		return kernel.ErrUnauthorized
	}
	if actor.OrganizationID != projectOrgID {
		return kernel.ErrForbidden
	}
	return nil
}

// AuthorizeOrgAction denies cross-org access by default.
func (p *Policy) AuthorizeOrgAction(actor *kernel.Actor, orgID kernel.OrganizationID) error {
	if actor == nil {
		return kernel.ErrUnauthorized
	}
	if actor.OrganizationID != orgID {
		return kernel.ErrForbidden
	}
	return nil
}

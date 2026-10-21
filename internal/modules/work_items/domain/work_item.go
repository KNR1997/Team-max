// Package domain contains the WorkItem entity and repository port.
package domain

import (
	"context"
	"errors"
	"strings"
)

// ErrNotFound is returned by the repository when a work item does not exist.
var ErrNotFound = errors.New("work item not found")

// WorkItem is a trackable unit of work.
type WorkItem struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Task is kept for backward compatibility while the codebase migrates.
type Task = WorkItem

// Repository persists work items. Implementations translate storage errors into
// ErrNotFound or wrapped system errors.
type Repository interface {
	// Create inserts w and fills its ID.
	Create(ctx context.Context, w *WorkItem) error
	// Get returns the work item with the given id, or ErrNotFound.
	Get(ctx context.Context, id int64) (*WorkItem, error)
	// List returns all work items ordered by id.
	List(ctx context.Context) ([]WorkItem, error)
	// Update replaces the stored work item with the given id, or returns ErrNotFound.
	Update(ctx context.Context, w *WorkItem) error
	// Delete removes the work item with the given id, or returns ErrNotFound.
	Delete(ctx context.Context, id int64) error
}

// Store is kept for backward compatibility.
type Store = Repository

// Validate normalizes the work item and reports invalid input.
func (w *WorkItem) Validate() error {
	w.Title = strings.TrimSpace(w.Title)
	if w.Title == "" {
		return errors.New("title is required")
	}

	const maxTitleLen = 500
	if len(w.Title) > maxTitleLen {
		return errors.New("title must be at most 500 characters")
	}

	return nil
}

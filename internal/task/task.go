// Package task holds the Task domain: the entity, the storage contract, and the HTTP API.
package task

import (
	"context"
	"errors"
	"strings"
)

// ErrNotFound is returned by Store when a task does not exist.
var ErrNotFound = errors.New("task not found")

// Task is a single todo item.
type Task struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Store persists tasks. Implementations translate storage errors into
// ErrNotFound or wrapped system errors.
type Store interface {
	// Create inserts t and fills its ID.
	Create(ctx context.Context, t *Task) error
	// Get returns the task with the given id, or ErrNotFound.
	Get(ctx context.Context, id int64) (*Task, error)
	// List returns all tasks ordered by id.
	List(ctx context.Context) ([]Task, error)
	// Update replaces the stored task with the given id, or returns ErrNotFound.
	Update(ctx context.Context, t *Task) error
	// Delete removes the task with the given id, or returns ErrNotFound.
	Delete(ctx context.Context, id int64) error
}

// Validate normalizes the task and reports invalid input.
func (t *Task) Validate() error {
	t.Title = strings.TrimSpace(t.Title)
	if t.Title == "" {
		return errors.New("title is required")
	}

	const maxTitleLen = 500
	if len(t.Title) > maxTitleLen {
		return errors.New("title must be at most 500 characters")
	}

	return nil
}

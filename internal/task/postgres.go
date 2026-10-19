package task

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresStore implements Store on PostgreSQL via pgxpool.
type PostgresStore struct {
	pool *pgxpool.Pool
}

// NewPostgresStore returns a Store backed by pool.
func NewPostgresStore(pool *pgxpool.Pool) *PostgresStore {
	return &PostgresStore{pool: pool}
}

// Create inserts the task and fills ID.
func (s *PostgresStore) Create(ctx context.Context, t *Task) error {
	const query = `INSERT INTO tasks (title, completed)
		VALUES ($1, $2)
		RETURNING id`

	err := s.pool.QueryRow(ctx, query, t.Title, t.Completed).Scan(&t.ID)
	if err != nil {
		return fmt.Errorf("inserting task: %w", err)
	}

	return nil
}

// Get returns the task with the given id, or ErrNotFound.
func (s *PostgresStore) Get(ctx context.Context, id int64) (*Task, error) {
	const query = `SELECT id, title, completed FROM tasks WHERE id = $1`

	t := new(Task)

	err := s.pool.QueryRow(ctx, query, id).Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, fmt.Errorf("selecting task %d: %w", id, err)
	}

	return t, nil
}

// List returns all tasks ordered by id.
func (s *PostgresStore) List(ctx context.Context) ([]Task, error) {
	const query = `SELECT id, title, completed FROM tasks ORDER BY id`

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("listing tasks: %w", err)
	}
	defer rows.Close()

	tasks := make([]Task, 0)

	for rows.Next() {
		var t Task

		if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
			return nil, fmt.Errorf("scanning task: %w", err)
		}

		tasks = append(tasks, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating tasks: %w", err)
	}

	return tasks, nil
}

// Update replaces the task with the given id, or returns ErrNotFound.
func (s *PostgresStore) Update(ctx context.Context, t *Task) error {
	const query = `UPDATE tasks
		SET title = $2, completed = $3
		WHERE id = $1
		RETURNING id`

	err := s.pool.QueryRow(ctx, query, t.ID, t.Title, t.Completed).Scan(&t.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrNotFound
		}

		return fmt.Errorf("updating task %d: %w", t.ID, err)
	}

	return nil
}

// Delete removes the task with the given id, or returns ErrNotFound.
func (s *PostgresStore) Delete(ctx context.Context, id int64) error {
	const query = `DELETE FROM tasks WHERE id = $1`

	tag, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("deleting task %d: %w", id, err)
	}

	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

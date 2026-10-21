package infrastructure

import (
	"context"
	"errors"
	"fmt"

	"todo/internal/modules/work_items/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository implements domain.Repository on PostgreSQL via pgxpool.
type PostgresRepository struct {
	pool *pgxpool.Pool
}

// PostgresStore is kept for backward compatibility.
type PostgresStore = PostgresRepository

// NewPostgresRepository returns a Repository backed by pool.
func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

// NewPostgresStore returns a Repository backed by pool.
func NewPostgresStore(pool *pgxpool.Pool) *PostgresRepository {
	return NewPostgresRepository(pool)
}

// Create inserts the work item and fills ID.
func (s *PostgresRepository) Create(ctx context.Context, w *domain.WorkItem) error {
	const query = `INSERT INTO tasks (title, completed)
		VALUES ($1, $2)
		RETURNING id`

	err := s.pool.QueryRow(ctx, query, w.Title, w.Completed).Scan(&w.ID)
	if err != nil {
		return fmt.Errorf("inserting work item: %w", err)
	}

	return nil
}

// Get returns the work item with the given id, or ErrNotFound.
func (s *PostgresRepository) Get(ctx context.Context, id int64) (*domain.WorkItem, error) {
	const query = `SELECT id, title, completed FROM tasks WHERE id = $1`

	w := new(domain.WorkItem)

	err := s.pool.QueryRow(ctx, query, id).Scan(&w.ID, &w.Title, &w.Completed)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("selecting work item %d: %w", id, err)
	}

	return w, nil
}

// List returns all work items ordered by id.
func (s *PostgresRepository) List(ctx context.Context) ([]domain.WorkItem, error) {
	const query = `SELECT id, title, completed FROM tasks ORDER BY id`

	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("listing work items: %w", err)
	}
	defer rows.Close()

	items := make([]domain.WorkItem, 0)

	for rows.Next() {
		var w domain.WorkItem

		if err := rows.Scan(&w.ID, &w.Title, &w.Completed); err != nil {
			return nil, fmt.Errorf("scanning work item: %w", err)
		}

		items = append(items, w)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating work items: %w", err)
	}

	return items, nil
}

// Update replaces the work item with the given id, or returns ErrNotFound.
func (s *PostgresRepository) Update(ctx context.Context, w *domain.WorkItem) error {
	const query = `UPDATE tasks
		SET title = $2, completed = $3
		WHERE id = $1
		RETURNING id`

	err := s.pool.QueryRow(ctx, query, w.ID, w.Title, w.Completed).Scan(&w.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.ErrNotFound
		}

		return fmt.Errorf("updating work item %d: %w", w.ID, err)
	}

	return nil
}

// Delete removes the work item with the given id, or returns ErrNotFound.
func (s *PostgresRepository) Delete(ctx context.Context, id int64) error {
	const query = `DELETE FROM tasks WHERE id = $1`

	tag, err := s.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("deleting work item %d: %w", id, err)
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}

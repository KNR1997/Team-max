//go:build integration

package task

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestPostgresStore(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Fatalf("creating pool: %v", err)
	}
	t.Cleanup(pool.Close)

	_, err = pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS tasks (
		id         BIGSERIAL PRIMARY KEY,
		title      TEXT NOT NULL CHECK (length(trim(title)) > 0),
		completed  BOOLEAN NOT NULL DEFAULT FALSE,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
	)`)
	if err != nil {
		t.Fatalf("creating schema: %v", err)
	}

	_, err = pool.Exec(ctx, "TRUNCATE tasks RESTART IDENTITY")
	if err != nil {
		t.Fatalf("truncating tasks: %v", err)
	}

	store := NewPostgresStore(pool)

	created := &Task{Title: "buy milk", Completed: false}

	if err := store.Create(ctx, created); err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if created.ID != 1 {
		t.Errorf("Create() ID = %d, want 1", created.ID)
	}

	got, err := store.Get(ctx, created.ID)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if got.Title != "buy milk" || got.Completed {
		t.Errorf("Get() = %+v, want title %q completed false", got, "buy milk")
	}

	created.Title = "buy oat milk"
	created.Completed = true

	if err := store.Update(ctx, created); err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	list, err := store.List(ctx)
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if len(list) != 1 || list[0].Title != "buy oat milk" || !list[0].Completed {
		t.Errorf("List() = %+v, want one updated task", list)
	}

	if err := store.Delete(ctx, created.ID); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	if _, err := store.Get(ctx, created.ID); !errors.Is(err, ErrNotFound) {
		t.Errorf("Get() after Delete() error = %v, want ErrNotFound", err)
	}

	if err := store.Delete(ctx, created.ID); !errors.Is(err, ErrNotFound) {
		t.Errorf("Delete() twice error = %v, want ErrNotFound", err)
	}

	if err := store.Update(ctx, &Task{ID: 999, Title: "ghost"}); !errors.Is(err, ErrNotFound) {
		t.Errorf("Update() missing error = %v, want ErrNotFound", err)
	}
}

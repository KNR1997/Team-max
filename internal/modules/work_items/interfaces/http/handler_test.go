package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo/internal/modules/work_items/domain"

	"github.com/go-chi/chi/v5"
)

const (
	caseOK     = "success"
	titleMilk  = "buy milk"
	taskPathID = "/tasks/1"
)

type fakeStore struct {
	create func(ctx context.Context, t *domain.WorkItem) error
	get    func(ctx context.Context, id int64) (*domain.WorkItem, error)
	list   func(ctx context.Context) ([]domain.WorkItem, error)
	update func(ctx context.Context, t *domain.WorkItem) error
	del    func(ctx context.Context, id int64) error
}

var _ domain.Repository = (*fakeStore)(nil)

func (f *fakeStore) Create(ctx context.Context, t *domain.WorkItem) error {
	if f.create == nil {
		return errors.New("create not implemented")
	}

	return f.create(ctx, t)
}

func (f *fakeStore) Get(ctx context.Context, id int64) (*domain.WorkItem, error) {
	if f.get == nil {
		return nil, errors.New("get not implemented")
	}

	return f.get(ctx, id)
}

func (f *fakeStore) List(ctx context.Context) ([]domain.WorkItem, error) {
	if f.list == nil {
		return nil, errors.New("list not implemented")
	}

	return f.list(ctx)
}

func (f *fakeStore) Update(ctx context.Context, t *domain.WorkItem) error {
	if f.update == nil {
		return errors.New("update not implemented")
	}

	return f.update(ctx, t)
}

func (f *fakeStore) Delete(ctx context.Context, id int64) error {
	if f.del == nil {
		return errors.New("delete not implemented")
	}

	return f.del(ctx, id)
}

func newTestHandler(t *testing.T, store domain.Repository) http.Handler {
	t.Helper()

	log := slog.New(slog.DiscardHandler)

	r := chi.NewRouter()
	NewHandler(store, log).Routes(r)

	return r
}

func doJSON(t *testing.T, h http.Handler, method, target, body string) *httptest.ResponseRecorder {
	t.Helper()

	var reader io.Reader
	if body != "" {
		reader = bytes.NewBufferString(body)
	}

	req := httptest.NewRequestWithContext(t.Context(), method, target, reader)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	return rec
}

func decodeTaskResponse(t *testing.T, rec *httptest.ResponseRecorder) domain.WorkItem {
	t.Helper()

	var got domain.WorkItem

	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("decoding response body: %v", err)
	}

	return got
}

func TestHandler_create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		store      domain.Repository
		body       string
		wantStatus int
		wantTitle  string
	}{
		{
			name: caseOK,
			store: &fakeStore{create: func(_ context.Context, task *domain.WorkItem) error {
				task.ID = 1
				return nil
			}},
			body:       `{"title":"buy milk"}`,
			wantStatus: http.StatusCreated,
			wantTitle:  titleMilk,
		},
		{
			name:       "empty title",
			store:      &fakeStore{},
			body:       `{"title":"  "}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "unknown field",
			store:      &fakeStore{},
			body:       `{"title":"x","bogus":1}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "malformed json",
			store:      &fakeStore{},
			body:       `{`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "store error",
			store: &fakeStore{create: func(_ context.Context, _ *domain.WorkItem) error {
				return errors.New("boom")
			}},
			body:       `{"title":"buy milk"}`,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := newTestHandler(t, tt.store)
			rec := doJSON(t, h, http.MethodPost, "/tasks", tt.body)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusCreated {
				if got := decodeTaskResponse(t, rec); got.Title != tt.wantTitle {
					t.Errorf("title = %q, want %q", got.Title, tt.wantTitle)
				}
			}
		})
	}
}

func TestHandler_list(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		store     domain.Repository
		wantIDs   []int64
		wantEmpty bool
	}{
		{
			name: "multiple tasks",
			store: &fakeStore{list: func(_ context.Context) ([]domain.WorkItem, error) {
				return []domain.WorkItem{{ID: 1, Title: "a"}, {ID: 2, Title: "b"}}, nil
			}},
			wantIDs: []int64{1, 2},
		},
		{
			name: "empty list is json array",
			store: &fakeStore{list: func(_ context.Context) ([]domain.WorkItem, error) {
				return []domain.WorkItem{}, nil
			}},
			wantEmpty: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := newTestHandler(t, tt.store)
			rec := doJSON(t, h, http.MethodGet, "/tasks", "")

			if rec.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
			}

			var got []domain.WorkItem

			if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
				t.Fatalf("decoding response body: %v", err)
			}

			if tt.wantEmpty {
				if len(got) != 0 {
					t.Fatalf("len(got) = %d, want 0", len(got))
				}

				return
			}

			if len(got) != len(tt.wantIDs) {
				t.Fatalf("len(got) = %d, want %d", len(got), len(tt.wantIDs))
			}

			for i, id := range tt.wantIDs {
				if got[i].ID != id {
					t.Errorf("got[%d].ID = %d, want %d", i, got[i].ID, id)
				}
			}
		})
	}
}

func TestHandler_get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		store      domain.Repository
		target     string
		wantStatus int
	}{
		{
			name: caseOK,
			store: &fakeStore{get: func(_ context.Context, id int64) (*domain.WorkItem, error) {
				return &domain.WorkItem{ID: id, Title: titleMilk}, nil
			}},
			target:     "/tasks/1",
			wantStatus: http.StatusOK,
		},
		{
			name: "not found",
			store: &fakeStore{get: func(_ context.Context, _ int64) (*domain.WorkItem, error) {
				return nil, domain.ErrNotFound
			}},
			target:     "/tasks/99",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "invalid id",
			store:      &fakeStore{},
			target:     "/tasks/abc",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := newTestHandler(t, tt.store)
			rec := doJSON(t, h, http.MethodGet, tt.target, "")

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestHandler_update(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		store      domain.Repository
		body       string
		wantStatus int
	}{
		{
			name: caseOK,
			store: &fakeStore{update: func(_ context.Context, _ *domain.WorkItem) error {
				return nil
			}},
			body:       `{"title":"buy oat milk","completed":true}`,
			wantStatus: http.StatusOK,
		},
		{
			name: "not found",
			store: &fakeStore{update: func(_ context.Context, _ *domain.WorkItem) error {
				return domain.ErrNotFound
			}},
			body:       `{"title":"x"}`,
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "empty title",
			store:      &fakeStore{},
			body:       `{"title":""}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := newTestHandler(t, tt.store)
			rec := doJSON(t, h, http.MethodPut, "/tasks/1", tt.body)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestHandler_delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		store      domain.Repository
		target     string
		wantStatus int
	}{
		{
			name: caseOK,
			store: &fakeStore{del: func(_ context.Context, _ int64) error {
				return nil
			}},
			target:     "/tasks/1",
			wantStatus: http.StatusNoContent,
		},
		{
			name: "not found",
			store: &fakeStore{del: func(_ context.Context, _ int64) error {
				return domain.ErrNotFound
			}},
			target:     "/tasks/99",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := newTestHandler(t, tt.store)
			rec := doJSON(t, h, http.MethodDelete, tt.target, "")

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestTask_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		title   string
		wantErr bool
		want    string
	}{
		{name: "valid", title: titleMilk, want: titleMilk},
		{name: "trims whitespace", title: "  buy milk  ", want: titleMilk},
		{name: "empty", title: "", wantErr: true},
		{name: "whitespace only", title: "   ", wantErr: true},
		{name: "too long", title: string(make([]byte, 501)), wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			task := &domain.WorkItem{Title: tt.title}

			err := task.Validate()

			if tt.wantErr {
				if err == nil {
					t.Fatal("Validate() error = nil, want error")
				}

				return
			}

			if err != nil {
				t.Fatalf("Validate() error = %v", err)
			}

			if task.Title != tt.want {
				t.Errorf("Title = %q, want %q", task.Title, tt.want)
			}
		})
	}
}

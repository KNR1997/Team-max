package task

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Handler serves the task HTTP API.
type Handler struct {
	store Store
	log   *slog.Logger
}

// NewHandler wires a task HTTP handler on store.
func NewHandler(store Store, log *slog.Logger) *Handler {
	return &Handler{store: store, log: log}
}

// Routes mounts the task endpoints on r.
func (h *Handler) Routes(r chi.Router) {
	r.Post("/tasks", h.create)
	r.Get("/tasks", h.list)
	r.Get("/tasks/{id}", h.get)
	r.Put("/tasks/{id}", h.update)
	r.Delete("/tasks/{id}", h.delete)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	t, ok := decodeTask(w, r)
	if !ok {
		return
	}

	if err := t.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.store.Create(r.Context(), t); err != nil {
		h.internalError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.List(r.Context())
	if err != nil {
		h.internalError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}

	t, err := h.store.Get(r.Context(), id)
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}

	body, ok := decodeTask(w, r)
	if !ok {
		return
	}

	if err := body.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	body.ID = id

	if err := h.store.Update(r.Context(), body); err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, body)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}

	if err := h.store.Delete(r.Context(), id); err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, ErrNotFound) {
		writeError(w, http.StatusNotFound, ErrNotFound.Error())
		return
	}

	h.internalError(w, err)
}

func (h *Handler) internalError(w http.ResponseWriter, err error) {
	h.log.Error("internal error", "error", err)
	writeError(w, http.StatusInternalServerError, "internal server error")
}

func pathID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id < 1 {
		writeError(w, http.StatusBadRequest, "invalid task id")
		return 0, false
	}

	return id, true
}

func decodeTask(w http.ResponseWriter, r *http.Request) (*Task, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	t := new(Task)

	if err := dec.Decode(t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body: "+err.Error())
		return nil, false
	}

	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		writeError(w, http.StatusBadRequest, "body must contain a single JSON object")
		return nil, false
	}

	return t, true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("encoding response", "error", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

package http

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"todo/internal/modules/work_items/domain"

	"github.com/go-chi/chi/v5"
)

// ErrorResponse is the JSON body returned for API failures.
type ErrorResponse struct {
	Error string `json:"error" example:"title is required"`
}

// Handler serves the work item HTTP API.
type Handler struct {
	repo domain.Repository
	log  *slog.Logger
}

// NewHandler wires a work item HTTP handler on the repository.
func NewHandler(repo domain.Repository, log *slog.Logger) *Handler {
	return &Handler{repo: repo, log: log}
}

// Routes mounts the work item endpoints on r.
func (h *Handler) Routes(r chi.Router) {
	r.Post("/tasks", h.create)
	r.Get("/tasks", h.list)
	r.Get("/tasks/{id}", h.get)
	r.Put("/tasks/{id}", h.update)
	r.Delete("/tasks/{id}", h.delete)
}

// create handles POST /tasks.
//
//	@Summary		Create a work item
//	@Description	Creates a work item and returns it with its generated ID.
//	@Tags			work-items
//	@Accept			json
//	@Produce		json
//	@Param			body	body		domain.WorkItem	true	"Work item to create"
//	@Success		201		{object}	domain.WorkItem
//	@Failure		400		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/tasks [post]
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	t, ok := decodeTask(w, r)
	if !ok {
		return
	}

	if err := t.Validate(); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.repo.Create(r.Context(), t); err != nil {
		h.internalError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

// list handles GET /tasks.
//
//	@Summary		List work items
//	@Description	Returns all work items ordered by ID.
//	@Tags			work-items
//	@Produce		json
//	@Success		200	{array}		domain.WorkItem
//	@Failure		500	{object}	ErrorResponse
//	@Router			/tasks [get]
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.List(r.Context())
	if err != nil {
		h.internalError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, tasks)
}

// get handles GET /tasks/{id}.
//
//	@Summary		Get a work item
//	@Description	Returns a single work item by ID.
//	@Tags			work-items
//	@Produce		json
//	@Param			id	path		int	true	"Work item ID"
//	@Success		200	{object}	domain.WorkItem
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/tasks/{id} [get]
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}

	t, err := h.repo.Get(r.Context(), id)
	if err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// update handles PUT /tasks/{id}.
//
//	@Summary		Update a work item
//	@Description	Replaces the stored work item with the given ID.
//	@Tags			work-items
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int				true	"Work item ID"
//	@Param			body	body		domain.WorkItem	true	"Replacement work item"
//	@Success		200		{object}	domain.WorkItem
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/tasks/{id} [put]
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

	if err := h.repo.Update(r.Context(), body); err != nil {
		h.handleError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, body)
}

// delete handles DELETE /tasks/{id}.
//
//	@Summary		Delete a work item
//	@Description	Removes the work item with the given ID.
//	@Tags			work-items
//	@Param			id	path	int	true	"Work item ID"
//	@Success		204	"No content"
//	@Failure		400	{object}	ErrorResponse
//	@Failure		404	{object}	ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/tasks/{id} [delete]
func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		h.handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		writeError(w, http.StatusNotFound, domain.ErrNotFound.Error())
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

func decodeTask(w http.ResponseWriter, r *http.Request) (*domain.WorkItem, bool) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	t := new(domain.WorkItem)

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
	writeJSON(w, status, ErrorResponse{Error: msg})
}

package tasks

import (
	"net/http"
	"log/slog"

	"github.com/gabrafo/todo_golang/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.ListTasks(r.Context())
	if err != nil {
		slog.Error("Failed to list tasks", "error", err)
		http.Error(w, "Failed to list tasks", http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, tasks)
}
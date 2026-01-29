package tasks

import (
	"context"
	"github.com/gabrafo/todo_golang/internal/adapters/sqlc"
)

type Service interface {
	ListTasks(ctx context.Context) ([]sqlc.Task, error)
}

type svc struct {

}

func NewService() Service {
	return &svc {
	}
}

func (s *svc) ListTasks(ctx context.Context) ([]sqlc.Task, error) {
	return nil, nil
}
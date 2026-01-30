package tasks

import (
	"context"
	repo "github.com/gabrafo/todo_golang/internal/adapters/sqlc"
)

type Service interface {
	ListTasks(ctx context.Context) ([]repo.Task, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc {
		repo: repo,
	}
}

func (s *svc) ListTasks(ctx context.Context) ([]repo.Task, error) {
	return s.repo.ListTasks(ctx)
}
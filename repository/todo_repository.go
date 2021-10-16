package repository

import (
	"context"
	"todo-backend/model/domain"
)

type TodoRepository interface {
	Save(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	Update(ctx context.Context, todo domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, todo domain.Todo) error
	FindByID(ctx context.Context, todoID uint) (domain.Todo, error)
	FindAll(ctx context.Context) ([]domain.Todo, error)
}

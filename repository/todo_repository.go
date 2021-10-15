package repository

import (
	"context"
	"todo-backend/model/domain"
	"todo-backend/model/web"
)

type TodoRepository interface {
	Save(ctx context.Context, todo web.TodoCreateRequest) domain.Todo
	Update(ctx context.Context, todo web.TodoCreateRequest) domain.Todo
	Delete(ctx context.Context, todoID uint)
	FindByID(ctx context.Context, todoID uint) domain.Todo
	FindAll(ctx context.Context) []domain.Todo
}

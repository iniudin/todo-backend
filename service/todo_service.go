package service

import (
	"context"
	"todo-backend/model/web"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse
	Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse
	Delete(ctx context.Context, todoID uint)
	FindByID(ctx context.Context, todoID uint) web.TodoResponse
	FindAll(ctx context.Context) []web.TodoResponse
}

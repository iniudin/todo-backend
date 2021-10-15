package repository

import (
	"context"
	"todo-backend/app"
	"todo-backend/model/domain"
	"todo-backend/model/web"
)

type TodoRepositoryImpl struct {
	Server *app.Server
}

func NewTodoRepository(server *app.Server) TodoRepository {
	return &TodoRepositoryImpl{
		Server: server,
	}
}

func (r *TodoRepositoryImpl) Save(ctx context.Context, todo web.TodoCreateRequest) domain.Todo {
	panic("not implemented") // TODO: Implement
}

func (r *TodoRepositoryImpl) Update(ctx context.Context, todo web.TodoCreateRequest) domain.Todo {
	panic("not implemented") // TODO: Implement
}

func (r *TodoRepositoryImpl) Delete(ctx context.Context, todoID uint) {
	panic("not implemented") // TODO: Implement
}

func (r *TodoRepositoryImpl) FindByID(ctx context.Context, todoID uint) domain.Todo {
	panic("not implemented") // TODO: Implement
}

func (r *TodoRepositoryImpl) FindAll(ctx context.Context) []domain.Todo {
	panic("not implemented") // TODO: Implement
}

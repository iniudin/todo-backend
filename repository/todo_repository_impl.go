package repository

import (
	"context"
	"todo-backend/app"
	"todo-backend/model/domain"
)

type TodoRepositoryImpl struct {
	Server *app.Server
}

func NewTodoRepository(server *app.Server) TodoRepository {
	return &TodoRepositoryImpl{
		Server: server,
	}
}

func (repository *TodoRepositoryImpl) Save(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	return todo, repository.
		Server.
		DB.WithContext(ctx).
		Create(&todo).Error
}

func (repository *TodoRepositoryImpl) Update(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	return todo, repository.
		Server.DB.
		WithContext(ctx).
		Model(&todo).
		Updates(&todo).Error
}

func (repository *TodoRepositoryImpl) Delete(ctx context.Context, todo domain.Todo) error {
	if err := repository.
		Server.DB.
		WithContext(ctx).
		First(&todo, todo.ID).Error; err != nil {
		return err
	}

	return repository.
		Server.DB.
		WithContext(ctx).
		Delete(&todo).Error
}

func (repository *TodoRepositoryImpl) FindByID(ctx context.Context, todoID uint) (todo domain.Todo, err error) {
	return todo, repository.
		Server.DB.
		WithContext(ctx).
		First(&todo, todoID).Error
}

func (repository *TodoRepositoryImpl) FindAll(ctx context.Context) (todos []domain.Todo, err error) {
	return todos, repository.
		Server.DB.
		WithContext(ctx).
		Find(&todos).Error
}

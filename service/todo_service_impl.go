package service

import (
	"context"
	"todo-backend/helper"
	"todo-backend/model/domain"
	"todo-backend/model/web"
	"todo-backend/repository"

	"gorm.io/gorm"
)

type TodoServiceImpl struct {
	TodoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
	}
}
func (service *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) web.TodoResponse {
	todo, err := service.TodoRepository.Save(ctx, domain.Todo{
		Title:       request.Title,
		Description: request.Description,
		IsDone:      request.IsDone,
	})
	helper.PanicIfError(err)

	return web.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}

}

func (service *TodoServiceImpl) Update(ctx context.Context, request web.TodoUpdateRequest) web.TodoResponse {
	var getTodo domain.Todo
	getTodo.ID = request.ID

	_, err := service.TodoRepository.FindByID(ctx, getTodo.ID)
	helper.PanicIfError(err)

	todo, err := service.TodoRepository.Update(ctx, domain.Todo{
		Model: gorm.Model{
			ID: request.ID,
		},
		Title:       request.Title,
		Description: request.Description,
		IsDone:      request.IsDone,
	})
	helper.PanicIfError(err)

	return web.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}
}

func (service *TodoServiceImpl) Delete(ctx context.Context, todoID uint) {
	var todo domain.Todo
	todo.ID = todoID

	err := service.TodoRepository.Delete(ctx, todo)
	helper.PanicIfError(err)

}

func (service *TodoServiceImpl) FindByID(ctx context.Context, todoID uint) web.TodoResponse {
	todo, err := service.TodoRepository.FindByID(ctx, todoID)
	helper.PanicIfError(err)

	return web.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		IsDone:      todo.IsDone,
	}

}

func (service *TodoServiceImpl) FindAll(ctx context.Context) []web.TodoResponse {
	var newTodoResponses []web.TodoResponse

	todoResponses, err := service.TodoRepository.FindAll(ctx)
	helper.PanicIfError(err)

	for _, todo := range todoResponses {
		newTodoResponses = append(newTodoResponses, web.TodoResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			IsDone:      todo.IsDone,
		})
	}
	return newTodoResponses
}

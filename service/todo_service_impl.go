package service

import "context"

type TodoServiceImpl struct {
}

func NewTodoService() TodoService {
	return &TodoServiceImpl{}
}

func (service *TodoServiceImpl) Create(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (service *TodoServiceImpl) Update(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (service *TodoServiceImpl) Delete(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (service *TodoServiceImpl) FindByID(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (service *TodoServiceImpl) FindAll(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

package service

import (
	"context"
)

type TodoService interface {
	Create(ctx context.Context) error
	Update(ctx context.Context) error
	Delete(ctx context.Context) error
	FindByID(ctx context.Context) error
	FindAll(ctx context.Context) error
}

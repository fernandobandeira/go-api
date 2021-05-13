package interfaces

import (
	"api/entities"
	"context"
)

type TodoRepository interface {
	Find(ctx context.Context) ([]entities.Todo, error)
	FindByID(ctx context.Context, id uint) (entities.Todo, error)
	Create(ctx context.Context, todo *entities.Todo) error
	Update(ctx context.Context, id uint, updateData entities.Todo) error
	Delete(ctx context.Context, id uint) error
}

type TodoService interface {
	Find(ctx context.Context) ([]entities.Todo, error)
	FindByID(ctx context.Context, id uint) (entities.Todo, error)
	Create(ctx context.Context, name string) (entities.Todo, error)
	Update(ctx context.Context, id uint, name string) error
	Delete(ctx context.Context, id uint) error
}

package services

import (
	"api/entities"
	"api/interfaces"
	"context"
)

func (service *todoService) Find(ctx context.Context) ([]entities.Todo, error) {
	return service.todoRepositorty.Find(ctx)
}

func (service *todoService) FindByID(ctx context.Context, id uint) (entities.Todo, error) {
	return service.todoRepositorty.FindByID(ctx, id)
}

func (service *todoService) Create(ctx context.Context, name string) (entities.Todo, error) {
	todo := entities.NewTodo(name)
	err := service.todoRepositorty.Create(ctx, &todo)
	return todo, err
}

func (service *todoService) Update(ctx context.Context, id uint, name string) error {
	updateData := entities.Todo{Name: name}
	return service.todoRepositorty.Update(ctx, id, updateData)
}

func (service *todoService) Delete(ctx context.Context, id uint) error {
	return service.todoRepositorty.Delete(ctx, id)
}

type todoService struct {
	todoRepositorty interfaces.TodoRepository
}

func NewTodoService(
	todoRepositorty interfaces.TodoRepository,
) interfaces.TodoService {
	return &todoService{
		todoRepositorty: todoRepositorty,
	}
}

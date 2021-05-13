package repositories

import (
	"api/entities"
	"api/interfaces"
	"context"

	"gorm.io/gorm"
)

func (repository *todoRepository) Find(ctx context.Context) ([]entities.Todo, error) {
	var todos = []entities.Todo{}
	result := repository.db.Find(&todos)
	return todos, result.Error
}

func (repository *todoRepository) FindByID(ctx context.Context, id uint) (entities.Todo, error) {
	var todo = entities.Todo{}
	result := repository.db.First(&todo, "id = ?", id)
	return todo, result.Error
}

func (repository *todoRepository) Create(ctx context.Context, todo *entities.Todo) error {
	result := repository.db.Create(todo)
	return result.Error
}

func (repository *todoRepository) Update(ctx context.Context, id uint, updateData entities.Todo) error {
	result := repository.db.Model(entities.Todo{
		Model: gorm.Model{
			ID: id,
		},
	}).Updates(updateData)
	return result.Error
}

func (repository *todoRepository) Delete(ctx context.Context, id uint) error {
	result := repository.db.Delete(&entities.Todo{
		Model: gorm.Model{
			ID: id,
		},
	})
	return result.Error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(
	db *gorm.DB,
) interfaces.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

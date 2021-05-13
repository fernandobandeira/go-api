package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name string `json:"name"`
}

func NewTodo(name string) Todo {
	return Todo{Name: name}
}

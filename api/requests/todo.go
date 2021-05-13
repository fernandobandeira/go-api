package requests

type PostTodo struct {
	Name string `json:"name" validate:"required" example:"Todo this"`
}

type PatchTodo struct {
	Name string `json:"name" validate:"required" example:"Todo this"`
}

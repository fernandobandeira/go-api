package handlers

import (
	"api/api/requests"
	"api/interfaces"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

func (handlers *todoHandlers) getTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := handlers.todoService.Find(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, todos)
}

func (handlers *todoHandlers) getTodo(w http.ResponseWriter, r *http.Request) {
	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	todo, err := handlers.todoService.FindByID(r.Context(), uint(ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, todo)
}

func (handlers *todoHandlers) postTodo(w http.ResponseWriter, r *http.Request) {
	var input requests.PostTodo
	err := utils.ReadJson(r, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	todo, err := handlers.todoService.Create(r.Context(), input.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, todo)
}

func (handlers *todoHandlers) patchTodo(w http.ResponseWriter, r *http.Request) {
	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var input requests.PatchTodo
	err = utils.ReadJson(r, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handlers.todoService.Update(r.Context(), uint(ID), input.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handlers *todoHandlers) deleteTodo(w http.ResponseWriter, r *http.Request) {
	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = handlers.todoService.Delete(r.Context(), uint(ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type todoHandlers struct {
	todoService interfaces.TodoService
}

func newTodoHandlers(todoService interfaces.TodoService) *todoHandlers {
	return &todoHandlers{todoService: todoService}
}

func TodoRouter(todoService interfaces.TodoService) http.Handler {
	handlers := newTodoHandlers(todoService)

	r := chi.NewRouter()
	r.Get("/", handlers.getTodos)
	r.Get("/{id:[0-9]+}", handlers.getTodo)
	r.Post("/", handlers.postTodo)
	r.Patch("/{id:[0-9]+}", handlers.patchTodo)
	r.Delete("/{id:[0-9]+}", handlers.deleteTodo)

	return r
}

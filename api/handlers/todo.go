package handlers

import (
	"api/api/requests"
	"api/entities"
	"api/interfaces"
	"api/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog"
)

func (handlers *todoHandlers) getTodos(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	todos, err := handlers.todoService.Find(ctx)
	if err != nil {
		handlers.presenters.Error(w, r, err)
		return
	}

	handlers.presenters.JSON(w, r, todos)
}

func (handlers *todoHandlers) getTodo(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	todo, err := handlers.todoService.FindByID(ctx, uint(ID))
	if err != nil {
		handlers.presenters.Error(w, r, err)
		return
	}

	handlers.presenters.JSON(w, r, todo)
}

func (handlers *todoHandlers) postTodo(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	var input requests.PostTodo
	err := utils.ReadJson(r, &input)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	v := validator.New()
	err = v.Struct(input)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	todo, err := handlers.todoService.Create(ctx, input.Name)
	if err != nil {
		handlers.presenters.Error(w, r, err)
		return
	}

	handlers.presenters.JSON(w, r, todo)
}

func (handlers *todoHandlers) patchTodo(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	var input requests.PatchTodo
	err = utils.ReadJson(r, &input)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	v := validator.New()
	err = v.Struct(input)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	err = handlers.todoService.Update(ctx, uint(ID), input.Name)
	if err != nil {
		handlers.presenters.Error(w, r, err)
		return
	}
}

func (handlers *todoHandlers) deleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx, span := utils.StartSpan(r.Context())
	defer span.End()

	IDArg := chi.URLParam(r, "id")
	ID, err := strconv.ParseUint(IDArg, 10, 64)
	if err != nil {
		handlers.presenters.Error(w, r, entities.ErrorBadRequest(err))
		return
	}

	err = handlers.todoService.Delete(ctx, uint(ID))
	if err != nil {
		handlers.presenters.Error(w, r, err)
		return
	}
}

type todoHandlers struct {
	logger      zerolog.Logger
	presenters  interfaces.Presenters
	todoService interfaces.TodoService
}

func newTodoHandlers(logger zerolog.Logger, presenters interfaces.Presenters, todoService interfaces.TodoService) *todoHandlers {
	return &todoHandlers{
		logger:      logger,
		presenters:  presenters,
		todoService: todoService,
	}
}

func TodoRouter(logger zerolog.Logger, presenters interfaces.Presenters, todoService interfaces.TodoService) http.Handler {
	handlers := newTodoHandlers(logger, presenters, todoService)

	r := chi.NewRouter()
	r.Get("/", handlers.getTodos)
	r.Get("/{id:[0-9]+}", handlers.getTodo)
	r.Post("/", handlers.postTodo)
	r.Patch("/{id:[0-9]+}", handlers.patchTodo)
	r.Delete("/{id:[0-9]+}", handlers.deleteTodo)

	return r
}

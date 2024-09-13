package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oreshkindev/snt-central-backend/common"
	"github.com/oreshkindev/snt-central-backend/internal/user/entity"
)

type UserController struct {
	usecase entity.UserUsecase
}

func NewUserController(usecase entity.UserUsecase) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	entity := &entity.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Create(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result.NewResponse())
}

func (controller *UserController) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	for i := range result {
		result[i] = *result[i].NewResponse()
	}

	render.JSON(w, r, result)
}

func (controller *UserController) First(w http.ResponseWriter, r *http.Request) {
	// get id from request
	email := chi.URLParam(r, "email")

	result, err := controller.usecase.First(email)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result.NewResponse())
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	id64, _ := strconv.ParseUint(id, 10, 64)

	entity := &entity.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Update(entity, id64)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result.NewResponse())
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	// get id from request
	id := chi.URLParam(r, "id")

	id64, _ := strconv.ParseUint(id, 10, 64)

	err := controller.usecase.Delete(id64)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}

func (controller *UserController) Authenticate(w http.ResponseWriter, r *http.Request) {
	entity := &entity.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Authenticate(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result.NewResponse())
}

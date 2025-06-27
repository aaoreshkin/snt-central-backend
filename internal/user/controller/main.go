package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/oreshkindev/snt-central-backend/common"
	"github.com/oreshkindev/snt-central-backend/internal/user/model"
)

type (
	Controller struct {
		usecase model.Usecase
	}
)

func New(usecase model.Usecase) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (controller *Controller) Create(w http.ResponseWriter, r *http.Request) {
	entity := &model.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Create(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) Find(w http.ResponseWriter, r *http.Request) {
	result, err := controller.usecase.Find()
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) First(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.First(id)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) Update(w http.ResponseWriter, r *http.Request) {

	entity := &model.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Update(entity)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	if err := controller.usecase.Delete(id); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, nil)
}

func (controller *Controller) Authenticate(w http.ResponseWriter, r *http.Request) {
	entity := &model.User{}

	if err := render.DecodeJSON(r.Body, entity); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	result, err := controller.usecase.Authenticate(entity)
	if err != nil {
		if err.Error() == "no rows in result set" {
			render.Render(w, r, common.ErrInvalidRequest(errors.New("Пользователь не найден или неверный пароль")))
		} else {
			render.Render(w, r, common.ErrInvalidRequest(err))
		}
		return
	}

	render.JSON(w, r, result)
}

func (controller *Controller) Revoke(w http.ResponseWriter, r *http.Request) {
	// Извлекаем заголовок Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		render.Render(w, r, common.ErrInvalidRequest(errors.New("отсутствует заголовок авторизации")))
		return
	}

	// Проверяем, что заголовок начинается с "Bearer "
	const bearerPrefix = "Bearer "
	if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
		render.Render(w, r, common.ErrInvalidRequest(errors.New("неверный формат заголовка авторизации")))
		return
	}

	// Извлекаем токен
	token := authHeader[len(bearerPrefix):]

	result, err := controller.usecase.Revoke(token)
	if err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, result)
}

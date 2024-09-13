package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/oreshkindev/rbac-middleware"
)

const (
	superuser rbac.Access = "superuser"
	manager   rbac.Access = "manager"
)

func (mux *Mux) handleUser() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.User.UserController

	router.With(rbac.Guard([]rbac.Access{superuser})).Post("/", controller.Create)
	router.With(rbac.Guard([]rbac.Access{superuser, manager})).Get("/", controller.Find)
	router.With(rbac.Guard([]rbac.Access{superuser, manager})).Get("/{email}", controller.First)
	router.With(rbac.Guard([]rbac.Access{superuser})).Put("/{id}", controller.Update)
	router.With(rbac.Guard([]rbac.Access{superuser})).Delete("/{id}", controller.Delete)
	router.Post("/auth", controller.Authenticate)

	return router
}

func (mux *Mux) handlePost() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Post.PostController

	router.With(rbac.Guard([]rbac.Access{superuser})).Post("/", controller.Create)
	router.Get("/", controller.Find)
	router.Get("/{id}", controller.First)
	router.With(rbac.Guard([]rbac.Access{superuser, manager})).Put("/{id}", controller.Update)
	router.With(rbac.Guard([]rbac.Access{superuser})).Delete("/{id}", controller.Delete)

	return router
}

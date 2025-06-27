package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/oreshkindev/rbac-middleware"
)

const (
	g int64 = iota + 1 // guest
	m                  // manager
	s                  // superuser
)

func (mux *Mux) handleUser() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.User.Controller

	router.With(rbac.Middleware(s, m)).Post("/", controller.Create)
	router.With(rbac.Middleware(s, m)).Get("/", controller.Find)
	router.With(rbac.Middleware(s, m)).Get("/{id}", controller.First)
	router.With(rbac.Middleware(s, m)).Put("/", controller.Update)
	router.With(rbac.Middleware(s, m)).Delete("/{id}", controller.Delete)
	router.Post("/auth", controller.Authenticate)
	router.Get("/revoke", controller.Revoke)

	return router
}

func (mux *Mux) handlePosts() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Posts.Controller

	router.With(rbac.Middleware(s, m)).Post("/", controller.Create)
	router.Get("/", controller.Find)
	router.Get("/{id}", controller.First)
	router.With(rbac.Middleware(s, m)).Put("/", controller.Update)
	router.With(rbac.Middleware(s, m)).Delete("/{id}", controller.Delete)

	return router
}

func (mux *Mux) handlerSchedules() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Schedules.Controller

	router.With(rbac.Middleware(s, m)).Post("/", controller.Create)
	router.Get("/", controller.Find)
	router.Get("/{id}", controller.First)
	router.With(rbac.Middleware(s, m)).Put("/", controller.Update)
	router.With(rbac.Middleware(s, m)).Delete("/{id}", controller.Delete)

	return router
}

func (mux *Mux) handlerInvoices() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Invoices.Controller

	router.With(rbac.Middleware(s, m)).Post("/", controller.Create)
	router.Get("/", controller.Find)
	router.Get("/{id}", controller.First)
	router.With(rbac.Middleware(s, m)).Put("/", controller.Update)
	router.With(rbac.Middleware(s, m)).Delete("/{id}", controller.Delete)

	return router
}

func (mux *Mux) handleEvents() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Event.Controller

	router.With(rbac.Middleware(s, m)).Post("/", controller.Create)
	router.Get("/", controller.Find)
	router.Get("/{id}", controller.First)
	router.With(rbac.Middleware(s, m)).Put("/", controller.Update)
	router.With(rbac.Middleware(s, m)).Delete("/{id}", controller.Delete)

	return router
}

func (mux *Mux) handlerAttachments() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Attachments.Controller

	router.With(rbac.Middleware(s, m)).Post("/", controller.Create)
	router.With(rbac.Middleware(s, m)).Get("/", controller.Find)

	return router
}

package router

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oreshkindev/rbac-middleware"
	"github.com/oreshkindev/snt-central-backend/internal"
)

type (
	Mux struct {
		*chi.Mux
		manager *internal.Manager
	}
)

const (
	superuser rbac.Access = "superuser"
	manager   rbac.Access = "manager"
)

func New(ctx context.Context, manager *internal.Manager) (*Mux, error) {

	// New router instance
	router := &Mux{
		chi.NewRouter(),
		manager,
	}

	// Add middlewares
	router.Use(
		// Enable CORS for all routes
		cors.AllowAll().Handler,
		// Set the content type to JSON for all responses
		render.SetContentType(render.ContentTypeJSON),
	)

	// Define routes for the router
	router.Route("/v1", func(r chi.Router) {

		// Mount the user handler on the "/v1/users" route
		r.Mount("/users", router.handlerUser())

		// Mount the user handler on the "/v1/posts" route
		r.Mount("/posts", router.handlerPost())
	})

	return router, nil
}

func (mux *Mux) handlerUser() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.User.UserController

	router.With(rbac.Guard([]rbac.Access{superuser})).Post("/", controller.Create)
	router.With(rbac.Guard([]rbac.Access{superuser})).Get("/", controller.Find)
	router.With(rbac.Guard([]rbac.Access{superuser})).Get("/{email}", controller.First)
	router.With(rbac.Guard([]rbac.Access{superuser})).Put("/{id}", controller.Update)
	router.With(rbac.Guard([]rbac.Access{superuser})).Delete("/{id}", controller.Delete)
	router.Post("/auth", controller.Authenticate)

	return router
}

func (mux *Mux) handlerPost() chi.Router {
	router := chi.NewRouter()

	controller := mux.manager.Post.PostController

	router.With(rbac.Guard([]rbac.Access{superuser})).Post("/", controller.Create)
	router.Get("/", controller.Find)
	router.Get("/{id}", controller.First)
	router.With(rbac.Guard([]rbac.Access{superuser})).Put("/{id}", controller.Update)
	router.With(rbac.Guard([]rbac.Access{superuser})).Delete("/{id}", controller.Delete)

	return router
}

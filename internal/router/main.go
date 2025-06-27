package router

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oreshkindev/snt-central-backend/internal"
)

type (
	Mux struct {
		*chi.Mux
		manager *internal.Manager
	}
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
		r.Mount("/users", router.handleUser())
		// Mount the user handler on the "/v1/posts" route
		r.Mount("/posts", router.handlePosts())
		// Mount the user handler on the "/v1/events" route
		r.Mount("/events", router.handleEvents())
		// Mount the user handler on the "/v1/attachment" route
		r.Mount("/attachments", router.handlerAttachments())
		// Mount the documents handler on the "/v1/schedules" route
		r.Mount("/schedules", router.handlerSchedules())
		// Mount the invoices handler on the "/v1/invoices" route
		r.Mount("/invoices", router.handlerInvoices())
	})

	return router, nil
}

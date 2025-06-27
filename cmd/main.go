package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/oreshkindev/snt-central-backend/internal"
	"github.com/oreshkindev/snt-central-backend/internal/database"
	"github.com/oreshkindev/snt-central-backend/internal/router"
)

var (
	connection *database.Connection
	mux        *router.Mux

	err error
)

func main() {
	// Create a context that is cancellable and cancel it on exit.
	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Run the application
	if err = run(context); err != nil {
		log.Println(err)
	}
}

func run(context context.Context) error {
	// Create a new database connection.
	if connection, err = database.New(context); err != nil {
		// There is no need to run the application without connecting to the database.
		panic(err)
	}
	// Close the connection when the program exits.
	defer connection.Close()

	// Create a new manager instance.
	manager, err := internal.New(context, connection)
	if err != nil {
		return err
	}

	// Create a new router with the necessary middlewares and routes.
	if mux, err = router.New(context, manager); err != nil {
		log.Println(err)
	}

	return http.ListenAndServe(os.Getenv("SERVICE_PORT"), mux)
}

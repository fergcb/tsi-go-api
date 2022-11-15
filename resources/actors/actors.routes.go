package actors

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", ListActors)
	router.Post("/", CreateActor)

	return router
}

package actors

import (
	"net/http"

	"github.com/go-chi/render"
	db "tsi.co/go-api/database"
	e "tsi.co/go-api/error"
)

func ListActors(w http.ResponseWriter, r *http.Request) {
	var actors []*Actor
	db.DB.Find(&actors)
	render.RenderList(w, r, NewActorListResponse(actors))
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	var data ActorRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, e.ErrInvalidRequest(err))
	}

	actor := data.Actor

	db.DB.Create(actor)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewActorResponse(actor))
}

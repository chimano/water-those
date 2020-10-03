package resources

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// PlantResource implements profile management handler.
type PlantResource struct {
}

// NewPlantResource creates and returns a plant resource.
func NewPlantResource() *PlantResource {
	return &PlantResource{}
}

func (rs *PlantResource) router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", rs.get)
	return r
}

type plantResponse struct {
	Msg string `json:"msg"`
}

func newPlantResponse() *plantResponse {
	return &plantResponse{
		Msg: "Hello World!",
	}
}

func (rs *PlantResource) get(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, newPlantResponse())
}

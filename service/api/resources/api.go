package resources

import (
	"github.com/go-chi/chi"
)

// API provides application resources and handlers.
type API struct {
	Plant *PlantResource
}

// NewAPI configures and returns application API.
func NewAPI() *API {
	plant := NewPlantResource()
	api := &API{
		Plant: plant,
	}
	return api
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/plant", a.Plant.router())

	return r
}

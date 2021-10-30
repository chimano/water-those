package plants

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chimano/water-those-service/api/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlantResource struct {
	plantCollection *mongo.Collection
}

type plant struct {
	Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
}

type plantResponse struct {
	Msg string `json:"msg"`
}

type plantInsertResponse struct {
	ID string `json:"id" bson:"_id"`
}

// NewPlantResource creates and returns a plant resource.
func NewPlantResource(plantCollection *mongo.Collection) *PlantResource {
	return &PlantResource{
		plantCollection: plantCollection,
	}
}

func (rs *PlantResource) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{plantId}", rs.get)
	r.Post("/", rs.post)
	return r
}

func newPlantInsertResponse(id interface{}) *plantInsertResponse {
	idString := id.(primitive.ObjectID).String()
	return &plantInsertResponse{
		ID: idString,
	}
}

func (rs *PlantResource) get(w http.ResponseWriter, r *http.Request) {
	plantId := chi.URLParam(r, "plantId")
	hexId, err := primitive.ObjectIDFromHex(plantId)
	if err != nil {
		log.Println("Invalid id")
	}
	filter := bson.D{{Key: "_id", Value: hexId}}

	var p plant

	ctx, cancel := db.GetContextWithDefaultTimeout()
	defer cancel()

	err = rs.plantCollection.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		log.Println(err)
		return
	}
	render.Respond(w, r, p)
}

func (rs *PlantResource) post(w http.ResponseWriter, r *http.Request) {
	var p plant

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}
	fmt.Print(p)

	ctx, cancel := db.GetContextWithDefaultTimeout()
	defer cancel()

	res, err := rs.plantCollection.InsertOne(ctx, p)
	if err != nil {
		render.Respond(w, r, err.Error())
		return
	}
	fmt.Print(res)

	render.Respond(w, r, newPlantInsertResponse(res.InsertedID))
}

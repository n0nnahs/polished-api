package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"github.com/n0nnahs/polished-api/configs"
	"github.com/n0nnahs/polished-api/models"
	"github.com/n0nnahs/polished-api/responses"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var polishCollection *mongo.Collection = configs.GetCollection(configs.DB, "polishes")
var validate = validator.New()

func CreatePolish() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var polish models.Polish
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&polish); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PolishResponse{
				Status: http.StatusBadRequest, 
				Message: "error", 
				Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&polish); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PolishResponse{
				Status: http.StatusBadRequest, 
				Message: "error", 
				Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newPolish := models.Polish{
			ID:               primitive.NewObjectID(),
			Name:             polish.Name,
			Brand:            polish.Brand,
			Collection:       polish.Collection,
			Type:             polish.Type,
			Color:            polish.Color,
			ColorDescription: polish.ColorDescription,
			PurchaseDate:     polish.PurchaseDate,
			Description:      polish.Description,
			Swatched:         polish.Swatched,
		}

		result, err := polishCollection.InsertOne(ctx, newPolish)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PolishResponse{
				Status: http.StatusInternalServerError, 
				Message: "error", 
				Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.PolishResponse{
			Status: http.StatusCreated, 
			Message: "success", 
			Data: map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAPolishId() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        params := mux.Vars(r)
        polishId := params["polishId"]
        var polish models.Polish
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(polishId)
	   fmt.Print(objId)

        err := polishCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&polish)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.PolishResponse{
			  Status: http.StatusInternalServerError, 
			  Message: "error", 
			  Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        rw.WriteHeader(http.StatusOK)
        response := responses.PolishResponse{
		   Status: http.StatusOK, 
		   Message: "success", 
		   Data: map[string]interface{}{"data": polish}}
        json.NewEncoder(rw).Encode(response)
    }
}

func GetAPolishName() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        params := mux.Vars(r)
	   polishName := params["name"]
        var polish models.Polish
        defer cancel()


        err := polishCollection.FindOne(ctx, bson.M{"name": polishName}).Decode(&polish)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.PolishResponse{
			  Status: http.StatusInternalServerError, 
			  Message: "error", 
			  Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        rw.WriteHeader(http.StatusOK)
        response := responses.PolishResponse{
		   Status: http.StatusOK, 
		   Message: "success", 
		   Data: map[string]interface{}{"data": polish}}
        json.NewEncoder(rw).Encode(response)
    }
}
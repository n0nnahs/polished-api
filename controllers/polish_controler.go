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
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&polish); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PolishResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}}
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
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.PolishResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetPolishId() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		polishId := params["polishId"]
		var polish models.Polish
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(polishId)
		fmt.Print(objId)

		err := polishCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&polish)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PolishResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PolishResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": polish}}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetPolishName() http.HandlerFunc {
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
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}

			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PolishResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": polish}}
		json.NewEncoder(rw).Encode(response)
	}
}

func EditPolish() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		polishId := params["polishId"]
		var polish models.Polish
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(polishId)

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&polish); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PolishResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&polish); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.PolishResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		update := bson.M{"name": polish.Name, "brand": polish.Brand, "collection": polish.Collection, "type": polish.Type, "color": polish.Color, "colorDescription": polish.ColorDescription, "purchaseDate": polish.PurchaseDate, "swatched": polish.Swatched}

		result, err := polishCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PolishResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//get updated polish details
		var updatedPolish models.Polish
		if result.MatchedCount == 1 {
			err := polishCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedPolish)

			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.PolishResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
				return
			}
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PolishResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPolish}}
		json.NewEncoder(rw).Encode(response)
	}
}

func DeletePolish() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		polishId := params["polishId"]
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(polishId)

		result, err := polishCollection.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PolishResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			rw.WriteHeader(http.StatusNotFound)
			response := responses.PolishResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "Polish with specified ID not found!"}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PolishResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "Polish successfully deleted!"}}
		json.NewEncoder(rw).Encode(response)
	}

}

func GetAllPolish() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Get All Polishes")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.Polish
		defer cancel()

		results, err := polishCollection.Find(ctx, bson.M{})

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.PolishResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.Polish
			if err = results.Decode(&singleUser); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.PolishResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(rw).Encode(response)
			}

			users = append(users, singleUser)
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.PolishResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}}
		json.NewEncoder(rw).Encode(response)
	}
}

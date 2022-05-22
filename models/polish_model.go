package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Polish struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Name             string        	 `bson:"name" json:"name" validate:"required"`
	Brand            string        	 `bson:"brand" json:"brand" validate:"required"`
	Collection       string       	 `bson:"collection" json:"collection,omitempty"`
	Type             []string    		 `bson:"type" json:"type,omitempty"`
	Color            []string     	 `bson:"color" json:"color,omitempty"`
	ColorDescription []string     	 `bson:"color_description" json:"colorDescription,omitempty"`
	PurchaseDate     string       	 `bson:"purchase_date" json:"purchaseDate,omitempty"`
	Description      string       	 `bson:"description" json:"description,omitempty"`
	Swatched         string       	 `bson:"swatched" json:"swatched" validate:"required"`
}
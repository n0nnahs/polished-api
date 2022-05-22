package polish

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Polish struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Name             string        	 `bson:"name" json:"name"`
	Brand            string        	 `bson:"brand" json:"brand"`
	Collection       string       	 `bson:"collection" json:"collection"`
	Type             []string    		 `bson:"type" json:"type"`
	Color            []string     	 `bson:"color" json:"color"`
	ColorDescription []string     	 `bson:"color_description" json:"colorDescription"`
	PurchaseDate     string       	 `bson:"purchase_date" json:"purchaseDate"`
	Description      string       	 `bson:"description" json:"description"`
	Swatched         string       	 `bson:"swatched" json:"swatched"`
}
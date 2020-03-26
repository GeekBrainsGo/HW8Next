package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mongo stands for mongo BSON ObjectID type.
type Mongo struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
}

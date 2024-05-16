package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User        string             `json:"user" bson:"user"`
	Description string             `json:"description" bson:"description"`
	Date        time.Time          `json:"datetime" bson:"datetime"`
	Trigger     bool               `json:"trigger" bson:"trigger"`
}

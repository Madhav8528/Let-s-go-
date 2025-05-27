package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cinema struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie     string             `json:"movie,omitempty"`
	IsWatched bool               `json:"watched,omitempty"`
}

package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type Image struct {
	ObjectId bson.ObjectId `bson:"_id" json:"_id"`
	Name     string        `bson:"name" json:"name"`
	Path     string        `bson:"path" json:"path"`
}
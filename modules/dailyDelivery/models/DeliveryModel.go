package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Delivery struct {
		Id        bson.ObjectId `bson:"_id,omitempty"`
		Address   string        `bson:"address"`
		ZipCode   string        `bson:"zipCode"`
		sussDelvd int           `bson:"sussDelvd"`
	}
)

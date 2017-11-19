package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Delivery struct {
		Id       bson.ObjectId `bson:"_id,omitempty"`
		IpPort   string        `bson:"ipPort"`
		Protocol string        `bson:"protocol"`
		Counter  int           `bson:"counter"`
	}
)

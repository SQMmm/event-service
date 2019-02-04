package domain

import "gopkg.in/mgo.v2/bson"

type Event struct {
	ID bson.ObjectId `bson:"_id"`
	Type string `bson:"type"`
	State int32 `bson:"state"`
}



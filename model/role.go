package model

import "gopkg.in/mgo.v2/bson"

type Role struct {
	Id      bson.ObjectId `bson:"_id,json:"id"`
	Name    string        `json:"name"`
	Created int64         `json:"created"`
}

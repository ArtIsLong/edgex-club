package model

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name"          json:"name"`
	Password string        `bson:"password"      json:"password"`
	GitHubId string        `bson:"gitHubId"      json:"gitHubId"`
	AvatarUrl string       `bson:"avatarUrl"     json:"avatarUrl"`
	Created  int64         `bson:"created"       json:"created"`
	Modified int64         `bson:"modified"      json:"modified"`
}

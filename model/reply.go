package model

import "gopkg.in/mgo.v2/bson"

type Reply struct {
	Id        bson.ObjectId  `bson:"_id,omitempty" json:"id"`
	Content   string         `bson:"content"       json:"content"`
	CommentId string         `bson:"commentId"     json:"commentId"`
	FromUserName  string     `bson:"fromUserName"  json:"fromUserName"`
  ToUserName  string       `bson:"toUserName"    json:"toUserName"`
	Created   int64          `bson:"created"       json:"created"`
}

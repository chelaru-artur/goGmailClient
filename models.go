package main

import (
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
)

type UserToken struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	TokenObj *oauth2.Token `json:"tokenObj" bson:"tokenObj"`
}

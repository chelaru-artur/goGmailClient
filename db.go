package main

import (
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

const dbURI string = "mongodb://user:user@ds031842.mongolab.com:31842/gmailclient"

func GetMongoSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(dbURI)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

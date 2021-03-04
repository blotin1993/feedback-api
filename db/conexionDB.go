package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN is the object of connecting to the database*/
var MongoCN = ConnectionDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://blotin:fp6QFdFZg5uffDaO@cluster0.6umd3.mongodb.net/test?authSource=admin&replicaSet=atlas-3vavae-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

/*ConectarBD is the feature that allows me to connect to the database
  Returns a connection to the BD of type Mongo Client*/
func ConnectionDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Successful connection to BD")
	return client
}

/*CheckConnection it's ping to BD*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

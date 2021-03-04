package main

import (
	"log"

	"github.com/blotin1993/feedback-api/db"
	//"github.com/vincent/feedback-api/handler"
	"github.com/joho/godotenv"
)

//MyEnv is the container for all the .env variables.
var MyEnv map[string]string

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("No connection to the BD")
		return
	}

	//handler.handlers()

	myEnv, err := godotenv.Read()
	s3Bucket := MyEnv["S3_BUCKET"]
}

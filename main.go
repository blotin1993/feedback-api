package main

import (
<<<<<<< HEAD
	"log"

	"github.com/blotin1993/feedback-api/db"
	//"github.com/vincent/feedback-api/handler"
=======
	"fmt"

>>>>>>> bf396c6bb10b12fa2166fd70e5637890ebbfa72e
	"github.com/joho/godotenv"
)

//MyEnv is the container for all the .env variables.
var MyEnv map[string]string

func main() {
<<<<<<< HEAD

	if db.CheckConnection() == 0 {
		log.Fatal("No connection to the BD")
		return
	}

	//handler.handlers()

	myEnv, err := godotenv.Read()
	s3Bucket := MyEnv["S3_BUCKET"]
=======
	MyEnv, err := godotenv.Read()
	if err == nil {
		dbURI := MyEnv["DB_URI"]
		fmt.Println(dbURI)
	}
>>>>>>> bf396c6bb10b12fa2166fd70e5637890ebbfa72e
}

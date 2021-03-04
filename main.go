package main

import (
	"log"

	"github.com/blotin1993/feedback-api/db"
	//"github.com/vincent/feedback-api/handler"
	"fmt"

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

	MyEnv, err := godotenv.Read()
	if err == nil {
		dbURI := MyEnv["DB_URI"]
		fmt.Println(dbURI)
	}
}

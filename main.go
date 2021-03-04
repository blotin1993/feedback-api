package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

//MyEnv is the container for all the .env variables.
var MyEnv map[string]string

func main() {
	MyEnv, err := godotenv.Read()
	if err == nil {
		dbURI := MyEnv["DB_URI"]
		fmt.Println(dbURI)
	}
}

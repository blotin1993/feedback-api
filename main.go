package main

import "github.com/joho/godotenv"

//MyEnv is the container for all the .env variables.
var MyEnv map[string]string

func main() {
	myEnv, err := godotenv.Read()
	s3Bucket := MyEnv["S3_BUCKET"]
}

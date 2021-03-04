package main

import "github.com/joho/godotenv"

//Environment map.
var myEnv map[string]string

func main() {
	myEnv, err := godotenv.Read()
	s3Bucket := myEnv["S3_BUCKET"]
}

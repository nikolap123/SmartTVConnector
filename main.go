package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)



func main() {

	godotenv.Load(".env")

	http.HandleFunc("/get-data" , HandleGetData)

	http.HandleFunc("/run-command" , HandleRunCommand)
	//
	//http.HandleFunc("/upload-dist" , HandleUploadDist)
	//
	//http.HandleFunc("/upload-build" , HandleUploadBuild)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

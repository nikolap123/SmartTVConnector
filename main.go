package main

import (
	"net/http"
	"log"
	
	"github.com/joho/godotenv"

)



func main() {

	godotenv.Load(".env")

	http.HandleFunc("/get-devices" , HandleGetDevices)

	http.HandleFunc("/get-applications" , HandleGetApplications)

	http.HandleFunc("/run-command" , HandleRunCommand)

	http.HandleFunc("/upload-dist" , HandleUploadDist)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

package main

import (
	"net/http"
	"log"
)



func main() {

	http.HandleFunc("/get-devices" , HandleGetDevices)

	http.HandleFunc("/run-command" , HandleRunCommand)

	http.HandleFunc("/upload-dist" , HandleUploadDist)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

package main

import (
	"log"
	"net/http"
)



func main() {

	http.HandleFunc("/get-devices" , HandleGetDevices)

	http.HandleFunc("/run-command" , HandleRunCommand)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

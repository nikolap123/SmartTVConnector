package main

import (
	"encoding/json"
	"log"
	"net/http"
)



func main() {

	http.HandleFunc("/get-devices" , HandleGetDevices)

	http.HandleFunc("/run-command" , HandleRunCommand)

	http.HandleFunc("/db-test",func (w http.ResponseWriter, r *http.Request) {

		var D Device 

		D = getDevice(1)

		res, err := json.Marshal(D)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

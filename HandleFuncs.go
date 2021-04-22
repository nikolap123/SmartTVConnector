package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "os"
	"net/http"
)



func HandleRunCommand(w http.ResponseWriter, r *http.Request) {

	var RCR RunCommandRequest
	var M Connector


	err := json.NewDecoder(r.Body).Decode(&RCR)

	M.init(RCR)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	res, err := json.Marshal(M)

	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}

func HandleGetDevices(w http.ResponseWriter, r *http.Request) {


	var devices Devices
	

	
	devices = getDevicesWithApplications()

	fmt.Printf("%+v\n",devices)
	res, err := json.Marshal(devices)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}
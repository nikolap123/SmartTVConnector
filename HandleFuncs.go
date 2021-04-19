package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

    devicesJson, err := os.Open("devices.json")
	applicationsJson, err := os.Open("applications.json")
    
    if err != nil {
        fmt.Println(err)
    }

    defer devicesJson.Close()
	defer applicationsJson.Close()

	devicesByteValue, _ := ioutil.ReadAll(devicesJson)
	applicationsByteValue, _ := ioutil.ReadAll(applicationsJson)

	var devices Devices
	var applications Applications

	json.Unmarshal(devicesByteValue, &devices)
	json.Unmarshal(applicationsByteValue, &applications)


	for _i, device := range devices.Devices {
		for _, application := range applications.Applications {
			
			if device.Id == application.Device_id {
				devices.Devices[_i].Applications = append(devices.Devices[_i].Applications,application)
			}
		}
		
	}
	res, err := json.Marshal(devices)

	// devices = getDevicesWithApplications()

	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}
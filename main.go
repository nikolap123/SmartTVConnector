package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "os"
	"io/ioutil"
)



func main() {

	http.HandleFunc("/get-devices",HandleGetDevices)

	http.HandleFunc("/get-applications",HandleGetApplications)

	http.HandleFunc("/run-command", HandleRunCommand)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

func HandleRunCommand(w http.ResponseWriter, r *http.Request) {

	/*
	* Initialize Micun with Command, Application and Device
	*/
	var M Micun
	err := json.NewDecoder(r.Body).Decode(&M.Command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	fmt.Fprintf(w, "Command: %+v", M.connect())
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

	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}

func HandleGetApplications(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test"))
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"os"
)


func HandleRunCommand(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r);

	if r.Method == "OPTIONS" {
		return
	}

	var RCR RunCommandRequest
	var M Connector

	err := json.NewDecoder(r.Body).Decode(&RCR)

	if err != nil {
		http.Error(w, "Something went wrong in GO", http.StatusBadRequest)
		return
	}

	M.init(RCR)

	command_response,command_error := RunCommand(M)

	if command_error != nil {
		w.Header().Set("Content-Type", "application/json")
		message := CommandResponse{Message:command_error.Error()}

		json,_ := json.Marshal(message)


		w.Write(json)
		return
	}

	message := CommandResponse{Message:"",Data:command_response.ResponseArray}

	json,_ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}

func HandleGetDevices(w http.ResponseWriter, r *http.Request) {

	var devices Devices
	var applications Applications
	
	devices = getDevices()
	applications = getApplications()

	var ADResponse = DevicesAndApplicationsResponse{Applications:applications.Applications,Devices:devices.Devices}
	res, err := json.Marshal(ADResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}

func HandleGetApplications(w http.ResponseWriter, r *http.Request) {

	var applications Applications
	
	applications = getApplications()

	res, err := json.Marshal(applications)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}

func HandleUploadDist(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r);

	if r.Method == "OPTIONS" {
		return
	}

	var ProjectsPath = os.Getenv("PROJECTS_PATH")
	
	var UPR UploadDistRequest

	r.ParseMultipartForm(10 << 20)

	ApplicationId,_ := strconv.Atoi(r.PostFormValue("ApplicationId"))
	DeviceType,_ := strconv.Atoi(r.PostFormValue("DeviceType"))

	UPR.ApplicationId = ApplicationId
	UPR.DeviceType = DeviceType

	var Application Application

	Application = getApplication(UPR.ApplicationId)

	fmt.Printf("%v",Application)

    file, _, err := r.FormFile("dist")

    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    defer file.Close()

    tempFile, err := ioutil.TempFile("tmp-zips", "upload-*.zip")

    if err != nil {
        fmt.Println(err)
    }
	defer os.Remove(tempFile.Name())
    defer tempFile.Close()


	fileBytes, err := ioutil.ReadAll(file)

	tempFile.Write(fileBytes)

	if err != nil {
        fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
    }

	fmt.Println(ResolveDeviceTypeUploadDist(DeviceType))

	for _,deviceType := range ResolveDeviceTypeUploadDist(DeviceType) {

		fmt.Println(ProjectsPath + deviceType + Application.Name)
		_,errUnzip := Unzip(tempFile.Name(),ProjectsPath + "/" + deviceType + "/" + Application.Name )

		if errUnzip != nil {
			fmt.Println(errUnzip)
			http.Error(w, errUnzip.Error(), http.StatusBadRequest)
		}

	}

	message := CommandResponse{Message:"Succesfully uploaded dist"}

	json,_ := json.Marshal(message)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return

	

}
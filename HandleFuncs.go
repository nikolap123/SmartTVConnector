package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const ProjectsPath = "C:/Users/Popa/SmartTV"

func HandleRunCommand(w http.ResponseWriter, r *http.Request) {

	var RCR RunCommandRequest
	var M Connector

	err := json.NewDecoder(r.Body).Decode(&RCR)

	M.init(RCR)

	res := RunCommand(M)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	w.Header().Set("Content-Type", "text/html")
  	w.Write([]byte(res))
}

func HandleGetDevices(w http.ResponseWriter, r *http.Request) {

	var devices Devices
	
	devices = getDevicesWithApplications()

	res, err := json.Marshal(devices)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	w.Header().Set("Content-Type", "application/json")
  	w.Write(res)
}

func HandleUploadDist(w http.ResponseWriter, r *http.Request) {

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
        return
    }

    defer file.Close()

    tempFile, err := ioutil.TempFile("temp-zips", "upload-*.zip")

    if err != nil {
        fmt.Println(err)
    }
    defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	tempFile.Write(fileBytes)

	if err != nil {
        fmt.Println(err)
    }
	fmt.Println(ResolveDeviceTypeUploadDist(DeviceType))
	for _,deviceType := range ResolveDeviceTypeUploadDist(DeviceType) {

		fmt.Println(ProjectsPath + deviceType + Application.Name)
		_,errUnzip := Unzip(tempFile.Name(),ProjectsPath + "/" + deviceType + "/" + Application.Name )

		if errUnzip != nil {
			fmt.Println(errUnzip)
		}

	}

	

}
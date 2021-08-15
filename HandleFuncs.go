package main

import (
	db "SmartTVConnector/database"
	"SmartTVConnector/smarttv"
	"encoding/json"
	"fmt"
	"net/http"
)


func HandleRunCommand(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r);

	if r.Method == "OPTIONS" {
		return
	}

	var RCR RunCommandRequest
	var C smarttv.ConnectorDTO

	err := json.NewDecoder(r.Body).Decode(&RCR)

	if err != nil {
		http.Error(w, "Request cannot be decoded!", http.StatusBadRequest)
		return
	}

	C.Init(RCR.DeviceId,RCR.ApplicationId,RCR.CommandName)

	commandResponse,err := RunCommand(C)

	if err != nil {
		fmt.Println(err.Error())
		message := Response{Message:err.Error()}
		json,_ := json.Marshal(message)
		w.WriteHeader(http.StatusBadRequest)

		w.Write(json)

		return
	}

	message := CommandResponse{Data:commandResponse}

	json,_ := json.Marshal(message)

	w.Write(json)
	return
}

func HandleGetData(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w,r);

	var devices db.Devices
	var applications db.Applications

	devices = db.GetDevices()
	applications = db.GetApplications()

	var ADResponse = DevicesAndApplicationsResponse{Applications:applications.Applications,Devices:devices.Devices}
	res, err := json.Marshal(ADResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

 	w.Write(res)
}
//
//func HandleUploadDist(w http.ResponseWriter, r *http.Request) {
//
//	setupResponse(&w,r);
//
//	if r.Method == "OPTIONS" {
//		return
//	}
//
//	var ProjectsPath = os.Getenv("PROJECTS_PATH")
//
//	var UPR UploadDistRequest
//
//	err := r.ParseMultipartForm(10 << 20)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	UPR.ApplicationId,_ = strconv.Atoi(r.PostFormValue("ApplicationId"))
//	UPR.DeviceType,_ = strconv.Atoi(r.PostFormValue("DeviceType"))
//
//	var Application Application
//
//	Application = getApplication(UPR.ApplicationId)
//
//	fmt.Printf("%v",Application)
//
//   file, _, err := r.FormFile("dist")
//
//   if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//       return
//   }
//
//   defer file.Close()
//
//   tempFile, err := ioutil.TempFile("tmp-zips", "upload-*.zip")
//
//   if err != nil {
//       fmt.Println(err)
//   }
//	defer os.Remove(tempFile.Name())
//   defer tempFile.Close()
//
//
//	fileBytes, err := ioutil.ReadAll(file)
//
//	tempFile.Write(fileBytes)
//
//	if err != nil {
//       fmt.Println(err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//   }
//
//	for _,deviceType := range ResolveDeviceTypeUpload(UPR.DeviceType) {
//
//		_,errUnzip := Unzip(tempFile.Name(),ProjectsPath + "/" + deviceType + "/" + Application.Name )
//
//		if errUnzip != nil {
//			http.Error(w, errUnzip.Error(), http.StatusBadRequest)
//		}
//
//	}
//
//	json,_ := json.Marshal(CommandResponse{Message:"Succesfully uploaded dist"})
//
//	w.Write(json)
//	return
//
//}
//
//func HandleUploadBuild(w http.ResponseWriter, r *http.Request) {
//
//	setupResponse(&w,r);
//
//	if r.Method == "OPTIONS" {
//		return
//	}
//
//	var ProjectsPath = os.Getenv("PROJECTS_PATH")
//
//	var UPR UploadDistRequest
//
//	r.ParseMultipartForm(10 << 20)
//
//	ApplicationId,_ := strconv.Atoi(r.PostFormValue("ApplicationId"))
//	DeviceType,_ := strconv.Atoi(r.PostFormValue("DeviceType"))
//
//	UPR.ApplicationId = ApplicationId
//	UPR.DeviceType = DeviceType
//
//	var Application Application
//
//	Application = getApplication(UPR.ApplicationId)
//
//	fmt.Printf("%v",Application)
//
//	file, h, err := r.FormFile("build")
//
//	if err != nil {
//		fmt.Println("Error Retrieving the File")
//		fmt.Println(err)
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	deviceType := ResolveDeviceTypeUpload(DeviceType);
//
//	f, err := os.OpenFile(ProjectsPath + "/" + deviceType[0] + "/" + h.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	io.Copy(f, file)
//
//	defer f.Close()
//	defer file.Close()
//
//
//	message := CommandResponse{Message:"Succesfully uploaded build"}
//
//	json,_ := json.Marshal(message)
//
//	w.Write(json)
//	return
//
//
//
//}
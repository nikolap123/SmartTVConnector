package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)



func main() {

	http.HandleFunc("/get-devices",HandleGetDevices)

	http.HandleFunc("/get-applications",HandleGetApplications)

	http.HandleFunc("/run-command", HandleRunCommand)

	http.HandleFunc("/get-devices-db", HanndleGetDevicesDB)

	http.HandleFunc("/insert-devices-db", HanndleInsertDevicesDB)

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

func HanndleGetDevicesDB(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	// query
	rows, err := db.Query("SELECT * FROM devices")
	checkErr(err)

	var Id int
	var Name string
	var IpAddress string
	var Year time.Time
	var Type string
	var Applications [] Application

	for rows.Next() {
		err = rows.Scan(&Id, &Name, &IpAddress, &Year, &Type, $Applications)
		checkErr(err)
		fmt.Println(Id)
		fmt.Println(Name)
		fmt.Println(IpAddress)
		fmt.Println(Year)
		fmt.Println(Type)
	}

	rows.Close()
}

func HanndleInsertDevicesDB(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO devices(name, ipAddress, year, type) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("STB BOX", "192.168.0.58", "2019", "Android")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
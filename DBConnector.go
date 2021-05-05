package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
	// "fmt"
)

func getDevice(DeviceId int ) Device {

	db, err := sqlx.Open("sqlite3", "./database_lite_smarttv.db")
	checkErr(err)

	rows,err := db.Queryx("select * from devices where id=?",DeviceId)

	
	var D Device

	for rows.Next() {
        err := rows.StructScan(&D)
        if err != nil {
            log.Fatalln(err)
        } 
    }

	return D
}

func getApplication(ApplicationId int ) Application {


	db, err := sqlx.Open("sqlite3", "./database_lite_smarttv.db")
	checkErr(err)

	rows,err := db.Queryx("select * from applications where id=?",ApplicationId)

	
	var A Application

	for rows.Next() {
        err := rows.StructScan(&A)
        if err != nil {
            log.Fatalln(err)
        } 
    }

	return A

}

func getDevicesWithApplications() Devices {

	db, err := sqlx.Open("sqlite3", "./database_lite_smarttv.db")
	checkErr(err)

	rows, err := db.Queryx("select d.id,d.name,d.ip_address,d.type as tv_type,d.year,a.name as app_name,a.id as app_id from devices as d join application_devices on d.id = application_devices.device_id join applications as a on a.id = application_devices.application_id")
	
	var appeard = make(map[int]int)
	var Ds Devices

	var it int = 0

	for rows.Next() {
		var (
		  id  int
		  app_id int
		  name   string
		  ip_address string
		  tv_type string
		  year int
		  app_name string
		)
		rows.Scan(&id,&name,&ip_address, &tv_type, &year, &app_name,&app_id)

		if _, ok := appeard[id]; ok {

			Ds.Devices[appeard[id]].Applications = append(Ds.Devices[appeard[id]].Applications,Application{Id: app_id, Name: app_name})

		} else {
			Ds.Devices = append(Ds.Devices,Device{Id:id,Name:name,IpAddress:ip_address,Type:tv_type,Year:year,Applications:[]Application{Application{Id: app_id, Name: app_name}}})

			appeard[id] = it
			it++

		}

	
	}


	return Ds
}


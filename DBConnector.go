package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
	// "fmt"
)


func getDevices() Devices {

	db, err := sqlx.Open("sqlite3", "./database_lite_smarttv.db")
	checkErr(err)

	rows,err := db.Queryx("select * from devices")

	
	var D Device
	var devices Devices

	for rows.Next() {
        err := rows.StructScan(&D)
		devices.Devices = append(devices.Devices,D)

        if err != nil {
            log.Fatalln(err)
        } 
    }

	return devices
}

func getApplications() Applications {

	db, err := sqlx.Open("sqlite3", "./database_lite_smarttv.db")
	checkErr(err)

	rows,err := db.Queryx("select * from applications")

	
	var A Application
	var applications Applications

	for rows.Next() {
        err := rows.StructScan(&A)
		applications.Applications = append(applications.Applications,A)

        if err != nil {
            log.Fatalln(err)
        } 
    }

	return applications
}

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


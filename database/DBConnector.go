package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


func GetDevices() Devices {

	db, err := sqlx.Open("sqlite3", "./database/database_lite_smarttv.db")
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

func GetApplications() Applications {

	db, err := sqlx.Open("sqlite3", "./database/database_lite_smarttv.db")
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

func GetDevice(DeviceId int ) Device {

	db, err := sqlx.Open("sqlite3", "./database/database_lite_smarttv.db")
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

func GetApplication(ApplicationId int ) Application {


	db, err := sqlx.Open("sqlite3", "./database/database_lite_smarttv.db")
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

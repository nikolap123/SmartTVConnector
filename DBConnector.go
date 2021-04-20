package main

import (
	"database/sql"
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

func getCommand(CommandId int ) Command {


	db, err := sqlx.Open("sqlite3", "./database_lite_smarttv.db")
	checkErr(err)

	rows,err := db.Queryx("select * from commands where id=?",CommandId)

	
	var C Command

	for rows.Next() {
        err := rows.StructScan(&C)
        if err != nil {
            log.Fatalln(err)
        } 
    }

	return C

}

func getDevicesWithApplications() Devices {

	db, err := sqlx.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	rows, err := db.Queryx("select * from devices inner join application_devices on devices.id = application_devices.device_id inner join on applications.id = application_devices.application_id;")

	log.Fatal(err)
	
	var Ds Devices

	for rows.Next() {
		err = rows.Scan()

	}

	return Ds
}


func createTable(db *sql.DB) {
	createDevicesTableSQL := `CREATE TABLE devices (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"ip_address" TEXT,
		"year" integer,
		"type" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create devices table...")
	statement, err := db.Prepare(createDevicesTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	statement.Exec() // Execute SQL Statements
	log.Println("devices table created")
}
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func getDevice(DeviceId int ) Device {

	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	rows, table_check := db.Query("select * from devices where id='';")

	log.Fatal(table_check)
	
	var D Device

	err = rows.Scan(&D)

	return D
}

func getApplication(ApplicationId int ) Application {


	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	rows, table_check := db.Query("select * from applications where id='';")

	log.Fatal(table_check)
	
	var A Application

	err = rows.Scan(&A)

	return A

}

func getCommand(CommandId int ) Command {


	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	rows, table_check := db.Query("select * from commands where id='';")
	
	log.Fatal(table_check)

	var C Command
	
	err = rows.Scan(&C)


	return C

}

func getDevicesWithApplications() Devices {

	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	rows, table_check := db.Query("select * from devices inner join application_devices on devices.id = application_devices.device_id inner join on applications.id = application_devices.application_id;")

	log.Fatal(table_check)
	
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
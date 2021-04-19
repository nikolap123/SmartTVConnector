package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	//"time"
)



func main() {

	http.HandleFunc("/get-devices" , HandleGetDevices)

	http.HandleFunc("/run-command" , HandleRunCommand)

	http.HandleFunc("/get-devices-db", HanndleGetDevicesDB)

	http.HandleFunc("/insert-devices-db", HanndleInsertDevicesDB)

	log.Fatal(http.ListenAndServe(":8081", nil))

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
	//var Year time.Time
	var Year int
	var Type string
	//var Applications [] Application

	for rows.Next() {
		err = rows.Scan(&Id, &Name, &IpAddress, &Year, &Type, /*&Applications*/)
		checkErr(err)
		fmt.Println(Id)
		fmt.Println(Name)
		fmt.Println(IpAddress)
		fmt.Println(Year)
		fmt.Println(Type)
	}
	w.Write([]byte("Test12"))
	rows.Close()
}

func HanndleInsertDevicesDB(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./database_lite.db")
	checkErr(err)

	_, table_check := db.Query("select * from devices;")

	if table_check == nil {
		fmt.Println("table is there")
	} else {
		fmt.Println("table not there")
		createTable(db) // Create Database Tables
	}

	// insert
	stmt, err := db.Prepare("INSERT INTO devices(name, ip_address, year, type) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("STB BOX", "192.168.0.58", "2019", "Android")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
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

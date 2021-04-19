package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	_ "github.com/mattn/go-sqlite3"
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
	var Year time.Time
	var Type string
	var Applications [] Application

	for rows.Next() {
		err = rows.Scan(&Id, &Name, &IpAddress, &Year, &Type, &Applications)
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

	// insert
	stmt, err := db.Prepare("INSERT INTO devices(name, ipAddress, year, type) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("STB BOX", "192.168.0.58", "2019", "Android")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

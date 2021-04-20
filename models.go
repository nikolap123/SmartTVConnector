package main

type Command struct {
	Id int `db:"id"`
	CommandName string `db:"command_name"`
	Provider string `db:"provider"`
}

type Device struct {
	Id int `db:"id"`
	Name string `db:"name"`
	IpAddress string `db:"ip_address"`
	Year int `db:"year"`
	Type string `db:"type"`
	Applications [] Application `json:"applications"`
}

type Application struct {
	Id int `db:"id"`
	Name string `db:"name"`
	Device_id int `db:"device_id"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}
type Applications struct {
	Applications []Application `json:"applications"`
}
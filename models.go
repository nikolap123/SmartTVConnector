package main

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
	WidgetName string `db:"widget_name"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}
type Applications struct {
	Applications []Application `json:"applications"`
}
package main


type Command struct {
	CommandName string
	Provider string
}

type Devices struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Id int `json:"id"`
	Name string `json:"name"`
	IpAddress string `json:"ip_address"`
	Year int `json:"year"`
	Type string `json:"type"`
	Applications [] Application `json:"applications"`
}

type Applications struct {
	Applications []Application `json:"applications"`
}

type Application struct {
	Name string `json:"name"`
	Device_id int `json:"device_id"`
}
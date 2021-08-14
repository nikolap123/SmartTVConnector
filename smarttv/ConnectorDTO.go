package smarttv

import (
	db "SmartTVConnector/database"
)

type ConnectorDTO struct {
	Device db.Device
	Application db.Application
	CommandName string
}

func (C *ConnectorDTO) Init(DeviceId int,ApplicationId int,CommandName string) {

	//TODO: Refactor

	if DeviceId == -1 {
		C.Device = db.Device{Name:"temp",IpAddress:"temp",Year:0,Type:"LG"}
	} else if DeviceId == -2 {
		C.Device = db.Device{Name:"temp",IpAddress:"temp",Year:0,Type:"Samsung"}
	} else {
		C.Device = db.GetDevice(DeviceId);
	}

	C.Application = db.GetApplication(ApplicationId);
	C.CommandName = CommandName
}






package main


import (
	"encoding/json"
	"fmt"
    "os"
	"io/ioutil"
)

type SmartTV interface {
	connect() bool
	build() Message
	run() Message

}

type Connector struct {
	Command Command
	Device Device
	Application Application
}

func (M Connector) connect() bool {
	fmt.Printf("TestConnection")
	return true
}

func (M Connector) build() bool {
	fmt.Printf("TestBuild")
	return true
}

func (M Connector) run() bool {
	fmt.Printf("TestRun")
	return true
}

func (M *Connector) init(RCR RunCommandRequest) {

	devicesJson, err := os.Open("devices.json")
    
    if err != nil {
        fmt.Println(err)
    }

    defer devicesJson.Close()

	var devices Devices

	devicesByteValue, _ := ioutil.ReadAll(devicesJson)
	json.Unmarshal(devicesByteValue, &devices)

	M.Device = devices.Devices[1]



	// M.Device = getDevice(RCR.DeviceId);
	// M.Application = getApplication(RCR.ApplicationId);
	// M.Command = getCommand(RCR.CommandId);
}



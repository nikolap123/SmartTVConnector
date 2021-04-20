package main


import (
	"fmt"
)

type SmartTV interface {
	connect() bool
	build() 
	run() 

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

func (M Connector) build()  {
	fmt.Printf("TestBuild")
}

func (M Connector) run()  {
	fmt.Printf("TestRun")
}

func (M *Connector) init(RCR RunCommandRequest) {

	M.Device = getDevice(RCR.DeviceId);
	M.Application = getApplication(RCR.ApplicationId);
	M.Command = getCommand(RCR.CommandId);
}



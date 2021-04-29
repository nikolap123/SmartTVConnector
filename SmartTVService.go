package main



type SmartTV interface {
	connect() bool
	build() 
	run() 

}

type Connector struct {
	// Command Command
	Device Device
	Application Application
}

func (M *Connector) init(RCR RunCommandRequest) {

	M.Device = getDevice(RCR.DeviceId);
	M.Application = getApplication(RCR.ApplicationId);
	// M.Command = getCommand(RCR.CommandId);
}






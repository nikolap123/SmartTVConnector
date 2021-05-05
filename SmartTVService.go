package main

type Connector struct {
	Device Device
	Application Application
}

func (M *Connector) init(RCR RunCommandRequest) {

	M.Device = getDevice(RCR.DeviceId);
	M.Application = getApplication(RCR.ApplicationId);
}






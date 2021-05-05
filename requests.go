package main

type RunCommandRequest struct {
	DeviceId int
	ApplicationId int
}

type UploadDistRequest struct {
	ApplicationId int
	DeviceType int // 0 - LG 1 - Samsung 2 - Both
}
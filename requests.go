package main

type RunCommandRequest struct {
	DeviceId int `json:"DeviceId"`
	ApplicationId int `json:"ApplicationId"`
	CommandName string `json:"CommandName"`
}

type UploadDistRequest struct {
	ApplicationId int
	DeviceType int // 0 - LG 1 - Samsung 2 - Both
}

type CommandResponse struct {
	Message string
}
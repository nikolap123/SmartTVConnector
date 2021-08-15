package main

import (
	db "SmartTVConnector/database"
	smartv "SmartTVConnector/smarttv"
)

//Requests

type RunCommandRequest struct {
	DeviceId int `json:"DeviceId"`
	ApplicationId int `json:"ApplicationId"`
	CommandName string `json:"CommandName"`
}

type UploadDistRequest struct {
	ApplicationId int
	DeviceType int // 0 - LG 1 - Samsung 2 - Both
}

// Responses

type Response struct {
	Message string
}

type CommandResponse struct {
	Data []smartv.TVCommandResult
}

type DevicesAndApplicationsResponse struct {
	Devices []db.Device
	Applications []db.Application
}
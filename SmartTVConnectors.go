package main


import (
	"fmt"
)

type SmartTV interface {
	connect() bool
	build() Message
	run() Message

}

type Micun struct {
	Command Command
	Device Device
	Application Application
}

func (M Micun) connect() bool {
	fmt.Printf("TestConnection")
	return true
}

func (M Micun) build() bool {
	fmt.Printf("TestBuild")
	return true
}

func (M Micun) run() bool {
	fmt.Printf("TestRun")
	return true
}



package main

import (
	"errors"
	"io/ioutil"
	"os"
	"fmt"

)

func RunCommand(M Connector) (SequenceResponse,error) {

	err := checkExecution(M)
	
	if err != nil {
		return SequenceResponse{},err
	}

	tvCommands,_ := RunBuilder(M)
	fmt.Println(tvCommands)

	var SequenceResponse SequenceResponse
	 
	tvCommands[0].exec()
	tvCommands[0].getResult(&SequenceResponse)

	return SequenceResponse,nil
}


func checkExecution(C Connector) error {

	if C.CommandName == "create-project" {

		files, _ := ioutil.ReadDir(os.Getenv("PROJECTS_PATH") + "/" + C.Device.Type)
	 
		for _, f := range files {
			if f.IsDir() && f.Name() == C.Application.Name {
				return errors.New("Project for this Application and Device already exists")
			}
		}

	} 

	return nil
}
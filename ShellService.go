package main

import (
	smartv "SmartTVConnector/smarttv"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func RunCommand(C smartv.ConnectorDTO) ([]smartv.TVCommandResult,error) {

	err := checkExecution(C)

	if err != nil {
		return nil,err
	}

	tvCommands,_ := smartv.RunBuilder(C)
	fmt.Println(tvCommands)

	var tvCommandResults []smartv.TVCommandResult

	tvCommands[0].Exec()
	tvCommands[0].GetResult(&tvCommandResults)

	return tvCommandResults,nil
}


func checkExecution(C smartv.ConnectorDTO) error {

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
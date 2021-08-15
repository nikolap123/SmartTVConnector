package smarttv

import (
	"errors"
	"github.com/Jeffail/gabs"
	"strconv"
)

type TVCommandFactory struct {}

type TVCommandFactoryInterface interface {
	Create()
}

func (B * TVCommandFactory) Create(C ConnectorDTO, commandName string) (TVCommand,error) {

	command,args,err := findCommand(C,commandName)

	if err != nil {
		return TVCommand{},err
	}

	var tvCommandArgs[] string

	var ArgumentsFactory ArgumentFactory

	for j := 0; j < len(args); j++ {

		arg,_ := ArgumentsFactory.Create(args[strconv.Itoa(j)].Data().(string),j,C,commandName)

		tvCommandArgs = append(tvCommandArgs,arg)

	}

	var TVCommand TVCommand

	TVCommand.setArgs(tvCommandArgs)
	TVCommand.setName(command)

	return TVCommand, nil
}

func findCommand(M ConnectorDTO,commandName string) (string,map[string]*gabs.Container,error) {

	command := parseJson("./smarttv/json_conf/commands.json").S(M.Device.Type,commandName)

	if command == nil {
		return "",nil,errors.New("Command " + commandName + " not found")
	}

	args,_ := command.S("args").ChildrenMap()

	return command.S("command").Data().(string),args,nil
}


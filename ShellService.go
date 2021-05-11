package main

import (
	"strconv"
	"strings"
	"errors"
	"io/ioutil"
	"os"
	"fmt"

)

func RunCommand(M Connector) (string,error) {


	if !checkExecution(M) {
		return "",errors.New("This command cannot be executed")
	}

	jsonParsed := parseJson("json_conf/commands.json")
	commandSequencesJson := parseJson("json_conf/commands_sequence.json")

	commandsSequence,_ := commandSequencesJson.S(M.Device.Type,M.CommandName).ChildrenMap();

	var tvCommands[] TVCommand

	for i := len(commandsSequence)-1; i >= 0; i-- {

		command := jsonParsed.S(M.Device.Type,commandsSequence[strconv.Itoa(i)].Data().(string))

		if command == nil {
			return "",errors.New("Command not found")
		}

		tvCommand := TVCommand{
			Command : command.S("command").Data().(string),
		}

		args,_ := command.S("args").ChildrenMap()

		var tvCommandArgs[] string

		for j := 0; j < len(args); j++ {

			if args[strconv.Itoa(j)].Data().(string) == "#" {

				arg,err := getDynamicArg(commandsSequence[strconv.Itoa(i)].Data().(string) + strconv.Itoa(j),M)

				if err != nil {
					return "",err
				}

				tvCommandArgs = append(tvCommandArgs,arg)

			} else {

				tvCommandArgs = append(tvCommandArgs, args[strconv.Itoa(j)].Data().(string))
			} 


		}

		tvCommand.Args = tvCommandArgs

		if i < len(commandsSequence)-1 {
			tvCommand.Next = &tvCommands[0]
		} 

		tvCommands = append([]TVCommand{tvCommand},tvCommands...)

	}

	if len(tvCommands) > 0 {
		tvCommands[0].exec()
	} else {
		return "",errors.New("Something went wrong")
	}
	// fmt.Println(tvCommands)
	return tvCommands[0].getResult(),nil
}

func getDynamicArg (key string,M Connector) (string,error) {

	jsonParsedCommandsMap := parseJson("json_conf/command_map.json")
	jsonParsedPropertyMap := parseJson("json_conf/property_to_command_map.json")

	if jsonParsedCommandsMap.S(M.Device.Type,key).Data() == nil {
		fmt.Println(key)
		return "",errors.New("Command key not found")
	}

	var ret_key = jsonParsedCommandsMap.S(M.Device.Type,key).Data().(string)

	exploded := strings.Split(ret_key,".")

	var ret_value string

	for _,a_key := range exploded {

		
		if jsonParsedPropertyMap.S(a_key).Data() == nil {
			
			return "", errors.New("cannot parse key " + a_key)
		}

		var a_key_value = ""

		if strings.HasPrefix(a_key,"DB_") {

			t_val :=  jsonParsedPropertyMap.S(a_key).Data().(string)

			a_key_value	= getField(&M,strings.Split(t_val,".")).Interface().(string)

		} else if strings.HasPrefix(a_key,"E_") {

			t_val :=  jsonParsedPropertyMap.S(a_key).Data().(string)

			a_key_value = os.Getenv(t_val)

		} else {
			a_key_value = jsonParsedPropertyMap.S(a_key).Data().(string)
		}

		ret_value = ret_value + a_key_value


	}

	return ret_value,nil
}

func checkExecution(C Connector) bool {

	if C.CommandName == "create-project" {

		files, _ := ioutil.ReadDir(os.Getenv("PROJECTS_PATH") + "/" +C.Device.Type)
	 
		for _, f := range files {
			if f.IsDir() && f.Name() == C.Application.Name {
				return false
			}
		}

	} 

	return true
}
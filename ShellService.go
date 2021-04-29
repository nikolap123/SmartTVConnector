package main

import (
	"strconv"
	"strings"
	"fmt"

)

func RunCommand(M Connector) {

	jsonParsed := parseJson("commands.json")

	var tvCommands[] TVCommand

	commands,_ := jsonParsed.S(M.Device.Type).ChildrenMap()

	for i := len(commands)-1; i >= 0; i-- {

		command := commands[strconv.Itoa(i)]

		tvCommand := TVCommand{
			Command : command.S("command").Data().(string),
		}

		args,_ := command.S("args").ChildrenMap()

		var tvCommandArgs[] string

		for j := 0; j < len(args); j++ {

			var arg = args[strconv.Itoa(j)].Data().(string)

			if args[strconv.Itoa(j)].Data().(string) == "#" {

				arg = getDynamicArg(string(M.Device.Type[0]) + strconv.Itoa(i) + strconv.Itoa(j),M)
			} 

			tvCommandArgs = append(tvCommandArgs,arg)

		}

		tvCommand.Args = tvCommandArgs

		if i < len(commands)-1 {
			tvCommand.Next = &tvCommands[0]
		} 

		tvCommands = append([]TVCommand{tvCommand},tvCommands...)

	}

	fmt.Printf("%v",tvCommands)
}

func getDynamicArg (key string,M Connector) string {

	jsonParsedCommandsMap := parseJson("samsung_command_map.json")
	jsonParsedPropertyMap := parseJson("samsung_property_to_command_map.json")

	var ret_key = jsonParsedCommandsMap.S(key).Data().(string)
	
	exploded := strings.Split(ret_key,".")

	var ret_value string

	for _,a_key := range exploded {

		var a_key_value = jsonParsedPropertyMap.S(a_key).Data().(string)

		if !strings.HasPrefix(a_key,"H_") {
			a_key_value	= getField(&M,strings.Split(a_key_value,".")).Interface().(string)
		} 

		ret_value = ret_value + a_key_value

	}

	return ret_value
}
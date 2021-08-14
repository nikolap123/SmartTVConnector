package smarttv

import (
	"errors"
	"github.com/Jeffail/gabs"
	"os"
	"strconv"
	"strings"
)

func RunBuilder(M ConnectorDTO) ([]TVCommand,error) {

	commandSequencesJson := parseJson("./smarttv/json_conf/commands_sequence.json")
	commandSequence,err := commandSequencesJson.S(M.Device.Type,M.CommandName).ChildrenMap()

	if err != nil {
		return nil,errors.New("Faild to find command sequence")
	}

	tvCommands := buildCommands(commandSequence,M)

	return tvCommands,nil
}

func findCommand(M ConnectorDTO,commandName string) (string,map[string]*gabs.Container,error) {

	commandJson := parseJson("./smarttv/json_conf/commands.json")
	command := commandJson.S(M.Device.Type,commandName)

	if command == nil {
		return "",nil,errors.New("Command not found")
	}

	args,_ := command.S("args").ChildrenMap()

	return command.S("command").Data().(string),args,nil
}

func buildCommands(commandSequence map[string]*gabs.Container, M ConnectorDTO ) ([] TVCommand) {

	var tvCommands[] TVCommand

	for i := len(commandSequence)-1; i >= 0; i-- {

		tvCommand,_ := buildCommand(commandSequence[strconv.Itoa(i)].Data().(string),M);

		if i < len(commandSequence)-1 {
			tvCommand.Next = &tvCommands[0]
		}

		tvCommands = append([]TVCommand{tvCommand},tvCommands...)

	}

	return tvCommands
}

func buildCommand(commandName string,M ConnectorDTO) (TVCommand,error) {

	command,args,err := findCommand(M,commandName)

	if err != nil {
		return TVCommand{},err
	}

	tvCommand := TVCommand{
		Command : command,
	}


	tvCommand.Args,_ = buildArgumets(args,M,commandName)

	return tvCommand,nil
}

func buildArgumets(args map[string]*gabs.Container, M ConnectorDTO,commandName string) ([]string,error) {

	var tvCommandArgs[] string

	for j := 0; j < len(args); j++ {

		arg,_ := buildArgumet(args[strconv.Itoa(j)].Data().(string),j,M,commandName)

		tvCommandArgs = append(tvCommandArgs,arg)

	}

	return tvCommandArgs,nil
}

func buildArgumet(arg string,j int,M ConnectorDTO,commandName string) (string,error) {

	if arg != "#"  {
		return arg,nil
	}

	arg,err := getDynamicArg(commandName + strconv.Itoa(j),M)

	if err != nil {
		return "",err
	}

	return arg,nil

}



func getDynamicArg (key string,M ConnectorDTO) (string,error) {

	jsonParsedCommandsMap := parseJson("./smarttv/json_conf/command_map.json")
	jsonParsedPropertyMap := parseJson("./smarttv/json_conf/command_args_map.json")

	if jsonParsedCommandsMap.S(M.Device.Type,key).Data() == nil {
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

			if(a_key == "DB_app_id") {
				a_key_value =  strings.ToLower(a_key_value)
			}

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

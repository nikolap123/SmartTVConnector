package smarttv

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type ArgumentFactory struct {}

type ArgumentFactoryInterface interface {
	Create()
}


func (R * ArgumentFactory) Create(arg string,j int,C ConnectorDTO,commandName string) (string,error) {

	if arg != "#"  {
		return arg,nil
	}

	arg,err := getDynamicArg(commandName + strconv.Itoa(j),C)

	if err != nil {
		return "",err
	}

	return arg,nil

}

func getDynamicArg (key string,C ConnectorDTO) (string,error) {

	jsonParsedCommandsMap := parseJson("./smarttv/json_conf/command_map.json")
	jsonParsedPropertyMap := parseJson("./smarttv/json_conf/command_args_map.json")

	if jsonParsedCommandsMap.S(C.Device.Type,key).Data() == nil {
		return "",errors.New("Command key not found")
	}

	var ret_key = jsonParsedCommandsMap.S(C.Device.Type,key).Data().(string)

	exploded := strings.Split(ret_key,".")

	var ret_value string

	for _,a_key := range exploded {


		if jsonParsedPropertyMap.S(a_key).Data() == nil {

			return "", errors.New("cannot parse key " + a_key)
		}

		var a_key_value = ""

		if strings.HasPrefix(a_key,"DB_") {

			t_val :=  jsonParsedPropertyMap.S(a_key).Data().(string)

			a_key_value	= getField(&C,strings.Split(t_val,".")).Interface().(string)

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

package main

import (
	"fmt"
	"io/ioutil"
    "os"
	"github.com/Jeffail/gabs"
	"strings"
	"reflect"
)

type SmartTV interface {
	connect() bool
	build() 
	run() 

}

type Connector struct {
	// Command Command
	Device Device
	Application Application
}

func (M *Connector) init(RCR RunCommandRequest) {

	M.Device = getDevice(RCR.DeviceId);
	M.Application = getApplication(RCR.ApplicationId);
	// M.Command = getCommand(RCR.CommandId);
}

func getField(v *Connector, field []string) reflect.Value {
    r := reflect.ValueOf(v)
	var value = reflect.Indirect(r)

	for _,p_key := range field {
		value = value.FieldByName(p_key)
	}

	return value
}

func getDynamicArg (key string,M Connector) string {

	jsonFileCommandsMap, err := os.Open("samsung_command_map.json")
	jsonFilePropertyMap, err := os.Open("samsung_property_to_command_map.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFileCommandsMap.Close()
	defer jsonFilePropertyMap.Close()

	jsonDataCommandsMap, _ := ioutil.ReadAll(jsonFileCommandsMap)
	jsonDataPropertyMap, _ := ioutil.ReadAll(jsonFilePropertyMap)

	jsonParsedCommandsMap, err := gabs.ParseJSON(jsonDataCommandsMap)
	jsonParsedPropertyMap, err := gabs.ParseJSON(jsonDataPropertyMap)

	if err != nil {
		panic(err)
	}

	
	var ret_key = jsonParsedCommandsMap.S(key).Data().(string)
	
	exploded := strings.Split(ret_key,".")

	for _,a_key := range exploded {

		a_key_value := jsonParsedPropertyMap.S(a_key).Data().(string)

		model_field := getField(&M,strings.Split(a_key_value,"."))

		fmt.Println(model_field)

	}

	return "test"
}



package main

import (
	"fmt"
    "io/ioutil"
    "os"
	"github.com/Jeffail/gabs"
	"strconv"
)

func RunCommand(M Connector) {
	jsonFile, err := os.Open("commands.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()
	
		jsonData, _ := ioutil.ReadAll(jsonFile)
	
		jsonParsed, err := gabs.ParseJSON(jsonData)
		if err != nil {
			panic(err)
		}

		var tvCommands[] TVCommand

		commands,_ := jsonParsed.S("Samsung").ChildrenMap()

		lenOfCommands := len(commands)

		for i := lenOfCommands-1; i >= 0; i-- {

			command := commands[strconv.Itoa(i)]

			tvCommand := TVCommand{
				Command : command.S("command").Data().(string),
			}

			args,_ := command.S("args").ChildrenMap()
			lenOfArgs := len(args)

			var tvCommandArgs[] string

			for j := 0; j < lenOfArgs; j++ {

				var arg = args[strconv.Itoa(j)].Data().(string)

				if args[strconv.Itoa(j)].Data().(string) == "#" {

					var key = "S" + strconv.Itoa(i) + strconv.Itoa(j)

					arg = getDynamicArg(key,M)
				} 

				tvCommandArgs = append(tvCommandArgs,arg)

			}

			tvCommand.Args = tvCommandArgs

			if i < lenOfCommands-1 {
				tvCommand.Next = &tvCommands[0]
			} 


			tvCommands = append([]TVCommand{tvCommand},tvCommands...)

		}
}
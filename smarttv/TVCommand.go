package smarttv

import "C"
import (
	 "bytes"
	 "fmt"
	 "os/exec"
	 "strings"
	 "time"
 )

type TVCommandResult struct {
	CommandExp string
	CommandResult string
}

type TVCommand struct {
	Command string
	Args []string
	Result TVCommandResult
	Next *TVCommand
}

type TVCommandInterface interface {
	Exec()
	GetResult()
}

func (R *TVCommand) setArgs(args []string) 	{ R.Args = args }
func (R *TVCommand) setName(name string)  	{ R.Command = name }

func (R *TVCommand) Exec()  {

	if R.Command == "ares-inspect" {
		// TODO: This should be used with CommandContext to set cmd.Process.Kill timeout
		cmd := exec.Command(R.Command,R.Args...)
		
		var outb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Start()

		time.Sleep(5 * time.Second)

		R.Result.CommandExp = R.Command + " " + strings.Join(R.Args," ")
		R.Result.CommandResult = outb.String()


	} else {
		fmt.Println(R.Command,R.Args)

		out, err := exec.Command(R.Command,R.Args...).Output()

		fmt.Println(string(out))

		R.Result.CommandExp = R.Command + " " + strings.Join(R.Args," ")
		R.Result.CommandResult = string(out)
		
		if err != nil {
			R.Next = nil
    	}
	}

	if R.Next != nil {
		R.Next.Exec()
	} 
}

func (R TVCommand) GetResult(result *[]TVCommandResult ) {

	*result = append((*result),R.Result)

	if R.Next != nil {
		R.Next.GetResult(result)
	}

	return

}


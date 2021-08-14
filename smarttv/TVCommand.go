package smarttv

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

func (C *TVCommand) Exec()  {

	if C.Command == "ares-inspect" {
		// TODO: This should be used with CommandContext to set cmd.Process.Kill timeout
		cmd := exec.Command(C.Command,C.Args...)
		
		var outb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Start()

		time.Sleep(5 * time.Second)

		C.Result.CommandExp = C.Command + " " + strings.Join(C.Args," ")
		C.Result.CommandResult = outb.String()


	} else {
		fmt.Println(C.Command,C.Args)

		out, err := exec.Command(C.Command,C.Args...).Output()

		fmt.Println(string(out))

		C.Result.CommandExp = C.Command + " " + strings.Join(C.Args," ")
		C.Result.CommandResult = string(out)
		
		if err != nil {
			C.Next = nil
    	}
	}

	if C.Next != nil {
		C.Next.Exec()
	} 
}

func (C TVCommand) GetResult(result *[]TVCommandResult ) {

	*result = append((*result),C.Result)

	if C.Next != nil {
		C.Next.GetResult(result)
	}

	return

}


 package main

 import (
	"fmt"
	"os/exec"
	"strings"
	"time"
	"bytes"
)
type SequenceResponse struct {
	ResponseArray []TVCommandResponse
}
type TVCommandResponse struct {
	CommandExp string
	CommandResult string
}

type TVCommand struct {
	Command string
	Args []string
	Result TVCommandResponse
	Next *TVCommand
}

type TVCommandInterface interface {
	exec()
	getResult()
}

func (C *TVCommand) exec()  {

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
		C.Next.exec()
	} 
}

func (C TVCommand) getResult(SR * SequenceResponse) {
	
	if C.Next == nil {
		SR.ResponseArray = append(SR.ResponseArray,C.Result)
		return 
	}

	SR.ResponseArray = append(SR.ResponseArray,C.Result)

	C.Next.getResult(SR)
}


 package main

 import (
	"fmt"
	"os/exec"
	"strings"
)


type TVCommand struct {
	Command string
	Args []string
	Result string
	Next *TVCommand
}

type TVCommandInterface interface {
	exec()
	getResult()
}

func (C *TVCommand) exec()  {

	fmt.Println(C.Command,C.Args)
	out, err := exec.Command(C.Command,C.Args...).Output()

	C.Result = C.Command + " " + strings.Join(C.Args," ") + "\n " + string(out) + "\n"
	
    if err != nil {
		C.Result = C.Result + "ERROR"
		C.Next = nil
    }


	if C.Next != nil {
		C.Next.exec()
	} 
}

func (C TVCommand) getResult() string {
	
	if C.Next == nil {
		return C.Result
	}

	return C.Result + C.Next.getResult()
}


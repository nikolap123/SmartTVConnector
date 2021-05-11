 package main

 import (
	"fmt"
	"os/exec"
	"log"
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

	C.Result = C.Command + " " + strings.Join(C.Args," ") + "\n " + string(out)
	
    if err != nil {
		fmt.Println("Error")
		log.Println(C.Command)
		log.Println(C.Args)
		C.Result = C.Result + "ERROR"
		// C.Result = C.Result + "\n " + C.Command + "\n "
		// C.Result = C.Result + "\n " + strings.Join(C.Args," ") + "\n "
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


 package main

 import (
	"fmt"
	"os/exec"
	"log"
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

	out, err := exec.Command(C.Command,C.Args...).Output()

	
    if err != nil {
		fmt.Println("Error")
		log.Println(C.Command)
		log.Println(C.Args)
        log.Fatal(err)
    }

	C.Result = string(out)

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


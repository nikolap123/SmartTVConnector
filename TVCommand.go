 package main
 import (
	"fmt"
	// "os/exec"
	// "log"
)


type TVCommand struct {
	Command string
	Args []string
	Next *TVCommand

}

type TVCommandInterface interface {
	exec()
	onError()
}

func (C TVCommand) exec()  {

	fmt.Println(C.Args)
// 	out, err := exec.Command(C.Command,C.Args...).Output()

//     if err != nil {
// 		fmt.Println("Error")
//         log.Fatal(err)
//     }

//     fmt.Println(string(out))

	if C.Next != nil {
		C.Next.exec()
	}
}

func (C TVCommand) onError()  {
	fmt.Println("TestRun")
}

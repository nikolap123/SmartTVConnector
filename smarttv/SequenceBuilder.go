package smarttv

import (
	"errors"
	"github.com/Jeffail/gabs"
	"os"
	"strconv"
)

// TODO: Split Sequence Builder and Sequence into to types

type SequenceBuilder struct  {
	connector ConnectorDTO
	sequence map[string]*gabs.Container
	commands []TVCommand
}

type SequenceBuilderInterface interface {
	Init()
	Build()
	GetRoot()
	GetCommands()
}

func (R * SequenceBuilder) GetCommands() ([] TVCommand) { return R.commands }

func (B *SequenceBuilder) Init(C ConnectorDTO) (error) {

	seq,err := parseJson(os.Getenv("COMMAND_SEQUENCE_PATH")).S(C.Device.Type,C.CommandName).ChildrenMap()

	if err != nil {
		return errors.New("Failed to find command " + C.CommandName + " for " + C.Device.Type + " TV")
	}

	B.connector = C
	B.sequence = seq

	return nil
}

func (B *SequenceBuilder) Build() (error) {

	var TVCommandFactory TVCommandFactory

	for i := len(B.sequence)-1; i >= 0; i-- {

		tvCommand,err := TVCommandFactory.Create(B.connector,B.sequence[strconv.Itoa(i)].Data().(string))

		if err != nil {
			return err
		}

		if i < len(B.sequence)-1 {
			tvCommand.Next = &B.commands[0]
		}

		B.commands = append([]TVCommand{tvCommand},B.commands...)

	}

	return nil
}


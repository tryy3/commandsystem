package commandsystem_test

import (
	"fmt"

	"github.com/tryy3/commandsystem"
)

var (
	sys commandsystem.System
)

func init() {
	sys := commandsystem.NewSystem()
}

func ExampleCommand() {
	CommandTest := &commandsystem.Command{
		Name:            "Test",
		Description:     "Short test description",
		LongDescription: "This is a longer command description",
		SubCommands: []*commandsystem.SubCommand{
			&commandsystem.SubCommand{
				Arguments: []*commandsystem.Argument{
					&commandsystem.Argument{
						Name: "parameter1",
						Type: commandsystem.ArgumentTypeMatch{Match: "hello"},
					},
				},
				Run: func(data *commandsystem.CommandData) error {
					fmt.Println("Hello world.")
					return nil
				},
			},
		},
	}

	sys.RegisterCommands(CommandTest)
}

package commandsystem

import "fmt"
import "strings"

// SimpleErrorHandler is a simple error handler for System
// simply prints out the error
func SimpleErrorHandler(err error) {
	fmt.Println(err)
}

// NewSystem creates a new System with the default settings
func NewSystem() *System {
	return &System{
		Commands:     make([]*Command, 0),
		Prefix:       []string{},
		ErrorHandler: SimpleErrorHandler,
	}
}

// System is the struct that handles all the commands
type System struct {
	Commands []*Command // All registered commands

	Prefix []string // A slice of string for prefixes.

	ErrorHandler func(error)
}

// RegisterCommands adds commands to the system
func (s *System) RegisterCommands(cmds ...*Command) {
	// TODO(tryy3): consider doing some checking here, if commands is properly set and such.
	s.Commands = append(s.Commands, cmds...)
}

// FindCommand checks for a registered command based on
// the commands name and aliases.
func (s *System) FindCommand(name string) *Command {
	for _, cmd := range s.Commands {
		if cmd.Name == name {
			return cmd
		}
		for _, alias := range cmd.Aliases {
			if alias == name {
				return cmd
			}
		}
	}
	return nil
}

// HandleCommand is called whenever a command might have been ran
// and then checks if its a valid command or not
func (s *System) HandleCommand(raw string) {
	for _, p := range s.Prefix {
		if strings.HasPrefix(raw, p) {
			noPrefix := strings.TrimPrefix(raw, p)
			args := strings.Split(noPrefix, " ")
			if len(args) <= 0 {
				s.ErrorHandler(fmt.Errorf("there is no command name"))
				return
			}

			cmd := s.FindCommand(args[0])
			if cmd == nil {
				s.ErrorHandler(fmt.Errorf("can't find the command %s", args[0]))
				return
			}

			data := NewCommandData(cmd)

			for _, sub := range cmd.SubCommands {
				ok, err := s.checkSubCommand(args[1:], sub, data)
				if !ok {
					continue
				}
				if err != nil {
					s.ErrorHandler(err)
					return
				}
			}
			s.ErrorHandler(fmt.Errorf("can't find any subcommands with the arguments supplied"))
		}
	}
}

// checkSubCommand checks for a valid sub command it also runts the SubCommand if
// all arguments are valid
func (s *System) checkSubCommand(args []string, sub *SubCommand, data *CommandData) (bool, error) {
	if len(args) < len(sub.Arguments) {
		return false, nil
	}

	for i, arg := range args {
		count := i
		if count >= len(sub.Arguments) {
			count = len(sub.Arguments) - 1
		}
		subArg := sub.Arguments[count]

		if !subArg.Type.Check(data, subArg.Name, arg) {
			return false, nil
		}
	}

	err := sub.Run(data)
	return true, err
}

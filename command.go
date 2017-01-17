package commandsystem

import (
	"fmt"
	"strings"
)

// CommandHelp is the help struct that is generated from a Command
type CommandHelp struct {
	Name            string   // The name of the command
	Aliases         string   // A joined list of the aliases slice
	Description     string   // The description of the command
	LongDescription string   // The long description of the command
	Subcommands     []string // A slice string of all the subcommands/syntaxes
}

// SubCommand holds the arguments for the command and the Run function if
// all the arguments was passed correctly
type SubCommand struct {
	Arguments []*Argument

	Run func(data *CommandData) error
}

// Command is the struct that holds all the information about the command
// and its subcommands
type Command struct {
	Name            string   // Name of the Command
	Aliases         []string // Aliases of the command name
	Description     string   // Description shown in command list
	LongDescription string   // Long description shown in command information

	HideHelp bool // Hide in help related functions (command list, command information etc.)

	SubCommands []*SubCommand // Slice of subcommands
}

// GenerateHelp generate a CommandHelp struct based on the command
func (c *Command) GenerateHelp() *CommandHelp {
	if c.HideHelp {
		return nil
	}

	help := &CommandHelp{
		Name:            c.Name,
		Aliases:         strings.Join(c.Aliases, ", "),
		Description:     c.Description,
		LongDescription: c.LongDescription,
	}

	subs := make([]string, 0)

	for _, sub := range c.SubCommands {
		str := c.Name + " "
		for _, arg := range sub.Arguments {
			str += arg.GetType() + " "
		}
		subs = append(subs, str)
	}
	help.Subcommands = subs
	return help
}

// ArgumentType is the interface used for all argument types
type ArgumentType interface {
	Check(data *CommandData, name string, raw string) bool
	String() string
}

// Argument is the struct for command arguments
type Argument struct {
	Name    string
	Type    ArgumentType
	Default interface{}
}

// GetType returns the values of the Argument
func (a *Argument) GetType() string {
	if a.Default != nil {
		return fmt.Sprintf("(%s|%s|%v)", a.Name, a.Type.String(), a.Default)
	}
	return fmt.Sprintf("(%s|%s)", a.Name, a.Type.String())
}

// ParsedArgument holds the information of a parsed Argument
type ParsedArgument struct {
	Raw    string
	Parsed interface{}
}

// NewCommandData creates a new command data with the default settings.
func NewCommandData(cmd *Command) *CommandData {
	return &CommandData{
		Command: cmd,
		Args:    make(map[string]*ParsedArgument, 0),
	}
}

// CommandData holds the dat for a command, used when parsing a command
type CommandData struct {
	Command *Command
	Args    map[string]*ParsedArgument
}

// GetArg returns a ParsedArgument from CommandData, if the ParsedArgument
// doesn't exists, it creates one and returns a empty ParsedArgument
func (d *CommandData) GetArg(name string) *ParsedArgument {
	arg, ok := d.Args[name]

	if !ok {
		arg := &ParsedArgument{Raw: "", Parsed: nil}
		d.Args[name] = arg
	}

	return arg
}

// SetArg set the ParsedArgument values
func (d *CommandData) SetArg(name string, raw string, value interface{}) {
	arg := d.GetArg(name)
	arg.Raw = raw
	arg.Parsed = value
}

// AddStringArg appends a string to the ParsedArgument
func (d *CommandData) AddStringArg(name string, value string) {
	arg := d.GetArg(name)

	str, ok := arg.Parsed.(string)

	if !ok {
		return
	}

	if arg.Raw != "" {
		value = " " + value
	}

	arg.Raw += value
	arg.Parsed = str + value
}

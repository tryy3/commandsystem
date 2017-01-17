package commandsystem

import (
	"strconv"
	"strings"
)

// ArgumentTypeString is meant to keep appending arguments as a string
type ArgumentTypeString struct {
}

func (t ArgumentTypeString) String() string {
	return "String"
}

// Check add data to the name, no need to actually check anything
func (t ArgumentTypeString) Check(data *CommandData, name string, raw string) bool {
	data.AddStringArg(name, raw)
	return true
}

// ArgumentTypeInt is meant to check if the argument is a int
type ArgumentTypeInt struct {
}

func (t ArgumentTypeInt) String() string {
	return "Int"
}

// Check converts the string to int and checks if there is an error or not
func (t ArgumentTypeInt) Check(data *CommandData, name string, raw string) bool {
	i, err := strconv.Atoi(raw)
	if err != nil {
		return false
	}
	data.SetArg(name, raw, i)
	return true
}

// ArgumentTypeFloat is meant to check if the argument is a float
type ArgumentTypeFloat struct {
}

func (t ArgumentTypeFloat) String() string {
	return "Float"
}

// Check converts the string to a float64 and checks if there is an error or not
func (t ArgumentTypeFloat) Check(data *CommandData, name string, raw string) bool {
	i, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return false
	}
	data.SetArg(name, raw, i)
	return true
}

// ArgumentTypeMatch is meant to compare the argument with the Match string
// to see if they match or not
// If Strict is true, then it will check with case-sensetive (compare string against string)
type ArgumentTypeMatch struct {
	Match  string
	Strict bool
}

func (t ArgumentTypeMatch) String() string {
	return "Match"
}

// Check matches 2 strings in strict and unstrict mode
func (t ArgumentTypeMatch) Check(data *CommandData, name string, raw string) bool {
	if t.Strict {
		if raw == t.Match {
			return true
		}
	} else {
		if strings.EqualFold(raw, t.Match) {
			return true
		}
	}
	return false
}

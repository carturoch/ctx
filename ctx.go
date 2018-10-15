package main

import (
	"fmt"

	"github.com/carturoch/ctx/listener"
)

type action struct {
	Name     string
	Desc     string
	Flag     string
	Listener func(string) string
}

// Argument wraps a valid received argument
type Argument struct {
	Value string
	Flag  string
}

func registerActions() []action {
	registered := []action{
		action{
			"Default",
			"Status check",
			"",
			listener.Default,
		},
	}
	return registered
}

// ParseArgs converts given args into app valid args
func ParseArgs(args []string) Argument {
	parsed := Argument{"", ""}
	return parsed
}

func main() {
	actions := registerActions()
	fmt.Println(actions[0].Listener(""))
}

package main

import (
	"fmt"
	s "strings"

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
func ParseArgs(args []string) (Argument, error) {
	parsed := Argument{"", ""}
	var cleanArgs []string
	for _, arg := range args {
		val := s.TrimSpace(arg)
		if val == "" {
			continue
		}
		cleanArgs = append(cleanArgs, val)
	}
	if len(cleanArgs) < 2 {
		return parsed, nil
	}
	parsed.Value = s.Join(cleanArgs[1:], " ")
	return parsed, nil
}

func main() {
	actions := registerActions()
	fmt.Println(actions[0].Listener(""))
}

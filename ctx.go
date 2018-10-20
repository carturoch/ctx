package main

import (
	"fmt"
	r "regexp"
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
func ParseArgs(args []string) Argument {
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
		return parsed
	}
	if match, _ := r.MatchString("^\\-{1,2}[a-zA-Z0-9]+", cleanArgs[1]); match {
		parsed.Flag = cleanArgs[1]
		cleanArgs = cleanArgs[1:]
	}
	parsed.Value = s.Join(cleanArgs[1:], " ")
	return parsed
}

func main() {
	actions := registerActions()
	fmt.Println(actions[0].Listener(""))
}

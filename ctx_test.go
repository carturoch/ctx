package main

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want Argument
	}{
		{
			"Returns empty argument when nothing is given",
			args{[]string{"ctx", ""}},
			Argument{"", ""},
		},
		{
			"Discards empty spaces",
			args{[]string{"ctx", "  "}},
			Argument{"", ""},
		},
		{
			"Detects the value",
			args{[]string{"ctx", "Some", "value"}},
			Argument{"Some value", ""},
		},
		{
			"Chars are parsed as value",
			args{[]string{"ctx", "-_-", "Some", "value"}},
			Argument{"-_- Some value", ""},
		},
		{
			"Parses short flags",
			args{[]string{"ctx", "-a", "Some", "value"}},
			Argument{"Some value", "-a"},
		},
		{
			"Parses long flags",
			args{[]string{"ctx", "--version"}},
			Argument{"", "--version"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseArgs(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		name    string
		args    args
		want    Argument
		wantErr bool
	}{
		{
			"Returns empty argument when nothing is given",
			args{[]string{"ctx", ""}},
			Argument{"", ""},
			false,
		},
		{
			"Discards empty spaces",
			args{[]string{"ctx", "  "}},
			Argument{"", ""},
			false,
		},
		{
			"Detects the value",
			args{[]string{"ctx", "Some", "value"}},
			Argument{"Some value", ""},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

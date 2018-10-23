package main

import "testing"

func TestGetQuery(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Returns empty if no argument is present",
			args{[]string{}},
			"",
		},
		{
			"Returns empty if only one arg is present",
			args{[]string{"ctx"}},
			"",
		},
		{
			"Joins with spaces the value",
			args{[]string{"ctx", "one", "two"}},
			"one two",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuery(tt.args.args); got != tt.want {
				t.Errorf("GetQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

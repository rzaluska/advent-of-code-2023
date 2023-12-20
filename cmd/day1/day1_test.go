package main

import "testing"

func TestPreprocessLine(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{input: "two1nine"}, want: "219"},
		{name: "test", args: args{input: "eightwothree"}, want: "823"},
		{name: "test", args: args{input: "abcone2threexyz"}, want: "abc123xyz"},
		{name: "test", args: args{input: "xtwone3four"}, want: "x2134"},
		{name: "test", args: args{input: "4nineeightseven2"}, want: "49872"},
		{name: "test", args: args{input: "zoneight234"}, want: "z18234"},
		{name: "test", args: args{input: "7pqrstsixteen"}, want: "7pqrst6teen"},
		{name: "test", args: args{input: "oneight"}, want: "18"},
	}
	st := PrepareSearchTree()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreprocessLine(st, tt.args.input); got != tt.want {
				t.Errorf("PreprocessLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessLine(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{input: "two1nine"}, want: 29},
		{name: "test", args: args{input: "eightwothree"}, want: 83},
		{name: "test", args: args{input: "abcone2threexyz"}, want: 13},
		{name: "test", args: args{input: "xtwone3four"}, want: 24},
		{name: "test", args: args{input: "4nineeightseven2"}, want: 42},
		{name: "test", args: args{input: "zoneight234"}, want: 14},
		{name: "test", args: args{input: "7pqrstsixteen"}, want: 76},
		{name: "test", args: args{input: "v4"}, want: 44},
		{name: "test", args: args{input: "oneight"}, want: 18},
	}
	st := PrepareSearchTree()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processLine(PreprocessLine(st, tt.args.input)); got != tt.want {
				t.Errorf("PreprocessLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree(t *testing.T) {
	values := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var st SearchTree
	for _, v := range values {
		st.Add(v)
	}
	t.Log(st.MatchPrefix("one"))
	t.Log(st.MatchPrefix("zero"))
	t.Log(st.MatchPrefix("eight"))
	t.Log(st.MatchPrefix("on"))
}

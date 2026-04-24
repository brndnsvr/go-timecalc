package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunWithPositionalArg(t *testing.T) {
	var stdout bytes.Buffer

	err := run([]string{"1.5"}, strings.NewReader(""), &stdout)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	want := "1.5 hours = 90 minutes = 5400 seconds\n"
	if stdout.String() != want {
		t.Fatalf("stdout = %q, want %q", stdout.String(), want)
	}
}

func TestRunPromptsWhenArgOmitted(t *testing.T) {
	var stdout bytes.Buffer

	err := run(nil, strings.NewReader("0.25\n"), &stdout)
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}

	want := "Enter hours: 0.25 hours = 15 minutes = 900 seconds\n"
	if stdout.String() != want {
		t.Fatalf("stdout = %q, want %q", stdout.String(), want)
	}
}

func TestRunRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "text", args: []string{"later"}},
		{name: "negative", args: []string{"-2"}},
		{name: "too many args", args: []string{"1", "2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer

			err := run(tt.args, strings.NewReader(""), &stdout)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}

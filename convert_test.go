package main

import "testing"

func TestConvertHours(t *testing.T) {
	tests := []struct {
		name    string
		hours   float64
		minutes int64
		seconds int64
	}{
		{name: "one hour", hours: 1, minutes: 60, seconds: 3600},
		{name: "one and a half hours", hours: 1.5, minutes: 90, seconds: 5400},
		{name: "quarter hour", hours: 0.25, minutes: 15, seconds: 900},
		{name: "zero hours", hours: 0, minutes: 0, seconds: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertHours(tt.hours)

			if got.Minutes != tt.minutes {
				t.Fatalf("minutes = %d, want %d", got.Minutes, tt.minutes)
			}
			if got.Seconds != tt.seconds {
				t.Fatalf("seconds = %d, want %d", got.Seconds, tt.seconds)
			}
		})
	}
}

func TestParseHours(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      float64
		wantError bool
	}{
		{name: "integer", input: "2", want: 2},
		{name: "decimal", input: "1.5", want: 1.5},
		{name: "trim spaces", input: " 0.25 ", want: 0.25},
		{name: "empty", input: "", wantError: true},
		{name: "spaces", input: "   ", wantError: true},
		{name: "text", input: "soon", wantError: true},
		{name: "negative", input: "-1", wantError: true},
		{name: "nan", input: "NaN", wantError: true},
		{name: "infinity", input: "+Inf", wantError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseHours(tt.input)
			if tt.wantError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("hours = %v, want %v", got, tt.want)
			}
		})
	}
}

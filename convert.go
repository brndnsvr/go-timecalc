package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type conversion struct {
	Hours   float64
	Minutes int64
	Seconds int64
}

func parseHours(input string) (float64, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return 0, fmt.Errorf("hours value is required")
	}

	hours, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hours value %q", input)
	}
	if math.IsNaN(hours) || math.IsInf(hours, 0) {
		return 0, fmt.Errorf("hours value must be a finite number")
	}
	if hours < 0 {
		return 0, fmt.Errorf("hours value must be zero or greater")
	}

	return hours, nil
}

func convertHours(hours float64) conversion {
	seconds := int64(math.Round(hours * 3600))

	return conversion{
		Hours:   hours,
		Minutes: seconds / 60,
		Seconds: seconds,
	}
}

func formatConversion(result conversion) string {
	return fmt.Sprintf("%s hours = %d minutes = %d seconds", strconv.FormatFloat(result.Hours, 'f', -1, 64), result.Minutes, result.Seconds)
}

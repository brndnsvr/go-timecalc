package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args[1:], os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "timecalc: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdin io.Reader, stdout io.Writer) error {
	if len(args) > 1 {
		return fmt.Errorf("expected one hours value, got %d", len(args))
	}

	var input string
	if len(args) == 1 {
		input = args[0]
	} else {
		fmt.Fprint(stdout, "Enter hours: ")

		scanner := bufio.NewScanner(stdin)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("read hours: %w", err)
			}
			return fmt.Errorf("hours value is required")
		}
		input = scanner.Text()
	}

	hours, err := parseHours(input)
	if err != nil {
		return err
	}

	fmt.Fprintln(stdout, formatConversion(convertHours(hours)))

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		// Trim the newline character from the input
		input = strings.TrimSpace(input)

		// Skip empty input
		if input == "" {
			continue
		}

		// Print the error message for invalid command
		fmt.Fprintf(os.Stderr, "%s: command not found\n", input)
	}
}

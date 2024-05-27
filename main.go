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

		// Exit 0 command to quit the shell
		if input == "exit 0" {
			os.Exit(0)
		}

		// Check for the 'echo' command
		if strings.HasPrefix(input, "echo ") {
			// Extract the text after 'echo '
			echoText := strings.TrimPrefix(input, "echo ")
			fmt.Println(echoText)
			continue
		}

		// type command to check if a command is a shell builtin
		if strings.HasPrefix(input, "type ") {
			// Extract the command to check after 'type '
			commandToCheck := strings.TrimPrefix(input, "type ")
			switch commandToCheck {
			case "echo":
				fmt.Fprintln(os.Stdout, "echo is a shell builtin")
			case "type":
				fmt.Fprintln(os.Stdout, "type is a shell builtin")
			case "exit":
				fmt.Fprintln(os.Stdout, "exit is a shell builtin")
			default:
				fmt.Fprintf(os.Stdout, "%s not found\n", commandToCheck)
			}
			continue
		}

		// Print the error message for invalid command
		fmt.Fprintf(os.Stderr, "%s: command not found\n", input)
	}
}

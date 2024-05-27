package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get the PATH environment variable
	pathEnv := os.Getenv("PATH")
	pathDirs := strings.Split(pathEnv, ":")

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
				// Search for the command in the PATH directories
				found := false
				for _, dir := range pathDirs {
					cmdPath := fmt.Sprintf("%s/%s", dir, commandToCheck)
					_, err := os.Stat(cmdPath)
					if err == nil {
						fmt.Printf("%s is %s\n", commandToCheck, cmdPath)
						found = true
						break
					}
				}
				if !found {
					fmt.Printf("%s not found\n", commandToCheck)
				}
			}
			continue
		}

		// Print the error message for invalid command
		fmt.Fprintf(os.Stderr, "%s: command not found\n", input)
	}
}

# go-command-shell

## Features

- **Built-in Commands**:
  - `echo`: Prints the provided text to the console.
  - `exit 0`: Exits the shell.
  - `type`: Identifies shell built-ins and searches for executables in the `PATH`.
- **External Program Execution**: Runs programs available in the system's `PATH` with arguments.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)

### Steps

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/Matusworld/go-command-shell.git
   cd go-command-shell
   ```

### Usage

```sh
./shell.sh
```

# Command Line Examples

This document provides examples of some common command line operations.

- Echo a string:

  ```bash
  $ echo Hello World
  Hello World
  ```

- Check if `echo` is a built-in shell command:

  ```bash
  $ type echo
  echo is a shell builtin
  ```

- Check the location of `ls` command:

  ```bash
  $ type ls
  ls is /bin/ls
  ```

- Try to find a nonexistent command:

  ```bash
  $ type nonexistent
  nonexistent not found
  ```

- Run `program_1234` with `alice` as an argument:

  ```bash
  $ program Joe
  Hello Joe
  ```

- Exit the shell with a status of 0:

  ```bash
  $ exit 0
  ```

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func echo(format string, values ...interface{}) {
	fmt.Printf(format+"\n", values...)
}

func stdinPipe() bool {
	fi, err := os.Stdin.Stat()

	if err != nil {
		panic(err)
	}

	return fi.Mode()&os.ModeNamedPipe != 0
}

func stdinLines() chan string {
	lines := make(chan string)

	go func() {
		defer close(lines)

		if !stdinPipe() {
			echo("No input piped")
			return
		}

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	return lines
}

func stdinInts() chan int64 {
	integers := make(chan int64)

	go func() {
		defer close(integers)

		for line := range stdinLines() {
			parsedInt, err := strconv.ParseInt(line, 10, 64)

			if err == nil {
				integers <- parsedInt
			} else {
				echo("Could not parse as int: %v", line)
			}
		}
	}()

	return integers
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			// os.Exit(1)
		}
		input := strings.TrimSpace(command)
		if input == "exit 0" {
			os.Exit(0)
		}
		if strings.HasPrefix(command, "echo") {
			toPrint := strings.TrimPrefix(input, "echo ")
			fmt.Println(toPrint)
		}
		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}

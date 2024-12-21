package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	force_close_prg_flag := false
	for {
		fmt.Fprint(os.Stdout, "$ ")
		if force_close_prg_flag {
			os.Exit(1)
		}
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading input:", err)
			// os.Exit(1)
		}
		fmt.Println(command[:len(command)-1] + ": command not found")
	}
}

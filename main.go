package main

import (
	"fmt"
	hangman "hangman/Game"
	"os"
)

func main() {
	args := os.Args[1:]
	var arg string
	if len(args) < 1 {
		fmt.Println("\033[31m", "Use words.txt to start a new game or --startWith save.txt to continue your last game.", "\033[0m")
	} else {
		for _, v := range args {
			arg += v
		}
		hangman.Run(arg)
	}
}

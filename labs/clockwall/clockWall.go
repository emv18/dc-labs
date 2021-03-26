package main

import (
	"os"
	"strings"
)

func parser(args []string, ch chan string) {
	for _, arg := range args {
		server := strings.Split(arg, "=")[1]
		ch <- server
	}
	close(ch)
}

func main() {
	args := os.Args[1:]

	ch := make(chan string)
	go parser(args, ch)
}

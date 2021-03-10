package main

import (
	"io"
	"log"
	"net"
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

func dial(ch chan string) {
	for i := range ch {
		server := i
		conn, err := net.Dial("tcp", server)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		client(os.Stdout, conn)
	}
}

func client(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)

	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	args := os.Args[1:]

	ch := make(chan string)
	go parser(args, ch)
	dial(ch)
}

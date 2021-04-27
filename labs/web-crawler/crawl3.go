// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func createFile(name string, urls []string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println("File creation failed")
		return
	}
	for _,url := range len(urls){
		f.Write(url+"\n")
	}
	defer f.Close()
}

func readParam() {
	args := os.Args[1:]
	var url string
	var depth int

	url = args[3]
	number = args[2].split('=')
	depth = number[2]
	return depth
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	depth := readParam()
	urls []string
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				if n > depth{
					return
				}
				go func(link string) {
					worklist <- crawl(link)
				}(link)
				append(urls, link)
			}
		}
	}
	createFile(UrlsList, urls)
}

//!-

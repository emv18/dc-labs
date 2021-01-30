package main

import (
	"fmt"
	"os"
)

var name string

func main() {
	for _,word := range os.Args[1:]{
    name = fmt.Sprintf("%v %v", name, word)
  }
  if len(name) < 1{
    fmt.Println("Error")
  } else{
    fmt.Println("Hello", name, "Welcome to the jungle")
  }
}

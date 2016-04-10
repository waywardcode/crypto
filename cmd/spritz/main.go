package main

import (
        "fmt"
	"sync"
	"os"
)

// ----------------------
// global command args
// ----------------------
var jobs int


// global limiter for parallelism
// -------------------------------
var limiter chan struct{} // limits the number of files we can work on at once
var wg sync.WaitGroup     // this is how we'll make sure all goroutines are done

func usage() {
  fmt.Fprintln(os.Stderr, "Usage:  spritz (hash|crypt) [args...]")
  fmt.Fprintln(os.Stderr, "Commands:  hash   compute the hash of inputs")
  fmt.Fprintln(os.Stderr, "           crypt  encrypt or decrypt inputs")
  fmt.Fprintln(os.Stderr, "  Give '-h' arg for further help on a command")
  os.Exit(2)
}

func main() {
   if len(os.Args) < 2 {
      usage()
   }
   switch os.Args[1] {
   case "hash":
       hashMain()
   case "crypt":
       cryptMain() 
   default:
       usage()
   }
}

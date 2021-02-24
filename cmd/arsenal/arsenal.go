package main

import (
	"flag"
	"fmt"
	"github.com/lukemilby/arsenal/pkg/arsenal"
	"os"
)

func main() {
	// Setup flag options
	// config
	config := flag.String("cfg", "config", "config file for reddits API")
	flag.Parse()
	// payload
	payload := flag.Args()
	if len(payload) <= 0 {
		flag.Usage = func(){
			fmt.Fprintf(os.Stderr, "Usage: ./arsenal <option> <payload> \n")
			flag.VisitAll(func(f *flag.Flag) {
				fmt.Fprintf(os.Stderr, "  -%s %v\n", f.Name, f.Usage) // f.Name, f.Value
			})
			fmt.Fprintf(os.Stderr, "  payload is a collection of words to look for e.g Rifle 9MM AK47\n")
		}
		flag.Usage()
		os.Exit(1)
	}

	arsenal.Run(*config, payload, nil)
}

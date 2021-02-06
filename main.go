package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var opts struct {
		Config string `short:"c" long:"config" description:"Config file (in YAML)" required:"true"`
	}

	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}

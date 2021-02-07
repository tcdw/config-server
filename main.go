package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/tcdw/config-server/config"
	"github.com/tcdw/config-server/server"
	"os"
)

func main() {
	var opts struct {
		Config string `short:"c" long:"config" description:"Config file (in JSON)" required:"true"`
	}

	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	conf, err := config.GetConfig(opts.Config)
	if err != nil {
		panic(err)
	}

	server.Start(*conf)
}

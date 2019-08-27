package main

import (
	"flag"
	"fmt"
	"os"

	app "github.com/cnnrrss/co-pilot/app"
)

func main() {
	flag.Usage = help
	flag.Parse()

	cmds := map[string]func(){
		"start":   app.Start,
		"version": app.PrintVersion,
		"help":    help,
	}

	if cmdFunc, ok := cmds[flag.Arg(0)]; ok {
		cmdFunc()
	} else {
		help()
		os.Exit(2)
	}
}

func help() {
	fmt.Fprintln(os.Stderr, `Usage:
	copilot start                      - start the server
	copilot help                       - show this message
	copilot version                    - display app version`)
}

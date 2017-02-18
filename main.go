package main

// move to CLI folder and file
import (
	"flag"
	"github.com/jonahglover/deplo1/cli"
	"github.com/jonahglover/deplo1/daemon"
)

func main() {
	daemonize := flag.Bool("daemonize", false, "daemoinze this thang")

	flag.Parse()

	if *daemonize {
		daemon.Start()
	} else {
		cli.RunCli()
	}
}

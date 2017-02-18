package cli

// move to CLI folder and file
import (
	"fmt"
	"github.com/jonahglover/deplo1/process"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

func startDaemon() (string, error) {
	b, err := ioutil.ReadFile("/Users/jonah/.deplo1/deplo1.pid")
	if err != nil {
		return err
	}

	pid := string(b)

	fmt.Println(pid)
}

func daemonize(c *cli.Context) error {
	pid, err := startDaemon()
}

func run(c *cli.Context) error {
	//	recipe, err := recipe.New(c.Args().Get(0))
	//
	//	if err != nil {
	//		return err
	//	}
	//
	//	err = process.New(recipe)
	//
	//	if err != nil {
	//		return err
	//	}

	daemonize()

	return nil
}

func list(c *cli.Context) error {
	fmt.Println(process.ProcessList)
	return nil
}

func RunCli() {
	app := cli.NewApp()
	app.Name = "deplo1"
	app.Usage = "deploy stuff"

	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "deplo1 run $file",
			Action: run,
		},
		{
			Name:   "daemonize",
			Usage:  "deplo1 daemonize",
			Action: daemonize,
		},
	}
	app.Run(os.Args)
}

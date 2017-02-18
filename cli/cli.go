package cli

// move to CLI folder and file
import (
	"fmt"
	"github.com/jonahglover/deplo1/client"
	"github.com/urfave/cli"
	"os"
)

func run(c *cli.Context) error {
	reply := client.Run(c.Args().Get(0))
	fmt.Println(reply)
	return nil
}

func list(c *cli.Context) error {
	client.List()
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
			Name:   "list",
			Usage:  "deplo1 list",
			Action: list,
		},
	}
	app.Run(os.Args)
}

package daemon

import (
	"fmt"
	"github.com/jonahglover/deplo1/procedure"
	"github.com/jonahglover/deplo1/process"
	"github.com/jonahglover/deplo1/recipe"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os/exec"
)

type Daemon struct{}

func (t *Daemon) Run(args *procedure.Args, r *procedure.Reply) error {

	recipe, err := recipe.New(args.Filename)

	c := exec.Command(recipe.Filename)
	err = c.Start()

	if err != nil {
		return err
	}

	proc := &process.Process{c.Process.Pid, args.Filename}
	process.ProcessList[c.Process.Pid] = proc

	return nil
}

func (t *Daemon) List(args *procedure.Args, r *procedure.Reply) error {
	for k, v := range process.ProcessList {
		fmt.Println(fmt.Sprintf("%d: %s", k, (*v).Filename))
	}
	return nil
}

func registerCommand(server *rpc.Server, d *Daemon) {
	server.RegisterName("Daemon", d)
}

func Start() {
	serve()
}

func serve() {

	c := new(Daemon)

	server := rpc.NewServer()
	registerCommand(server, c)

	server.HandleHTTP("/", "/debug")

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)
}

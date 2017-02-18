package process

import (
	"fmt"
	"github.com/jonahglover/deplo1/recipe"
	"os/exec"
)

// Process definition.
type Process struct {
	Pid      int
	Filename string
}

var ProcessList = make(map[int]*Process)

func New(r *recipe.Recipe) error {
	cmd := exec.Command(r.Filename)
	err := cmd.Start()
	if err != nil {
		return err
	}

	p := new(Process)

	p.Filename = r.Filename
	p.Pid = cmd.Process.Pid

	ProcessList[p.Pid] = p

	fmt.Println(p)

	return nil
}

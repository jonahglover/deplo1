package process

import (
	//	"os"
	"fmt"
	"github.com/jonahglover/deplo1/recipe"
	"os/exec"
)

// Process definition.
type Process struct {
	Id       int
	Filename string
}

var ProcessList = make(map[int]*Process)

func New(r *recipe.Recipe) error {
	cmd := exec.Command(r.Filename)
	// cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		return err
	}

	p := new(Process)

	p.Filename = r.Filename
	p.Id = cmd.Process.Pid

	ProcessList[p.Id] = p

	fmt.Println(p)

	return nil
}

package client

import (
	"github.com/jonahglover/deplo1/procedure"
	"log"
	"net/rpc"
)

func Run(filename string) procedure.Reply {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &procedure.Args{filename}

	var reply procedure.Reply

	err = client.Call("Daemon.Run", args, &reply)

	if err != nil {
		log.Fatal("daemon error:", err)
	}
	return reply
}

func List() procedure.Reply {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply procedure.Reply

	args := &procedure.Args{}

	err = client.Call("Daemon.List", args, &reply)

	if err != nil {
		log.Fatal("daemon error:", err)
	}
	return reply
}

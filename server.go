package main

import (
	"fmt"
	"log"

	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"os/exec"

	"runtime"
	"strconv"
)

type LatencyArgs struct {
	N int
}

type Llama string

func (c *Llama) Latency(args *LatencyArgs, reply *string) error {
	//fmt.Printf("args: %+v\n", args)

	os := runtime.GOOS
	if os != "darwin" && os != "linux" && os != "freebsd" {
		fmt.Printf("Platform %s not supported\n", runtime.GOOS)
	}

	// the script must be executable (chmod +x)
	out, err := exec.Command(
		"./scripts/latency_"+runtime.GOOS+".sh",
		strconv.Itoa(args.N)).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output from script is: %s\n", out)

	*c = "Latency"
	*reply = "ok"
	return nil
}

func main() {
	llama := new(Llama)

	rpc.Register(llama)

	listener, err := net.Listen("tcp", ":11414")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go jsonrpc.ServeConn(conn)
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/AldieNightStar/netgo"
)

func main() {
	cmds := netgo.NewCommands()
	mem := "0"
	cmds.SetCommand("set", func(s string) string {
		mem = s
		return s
	})
	cmds.SetCommand("get", func(s string) string {
		return mem
	})
	srv := netgo.NewServer(9999, cmds)
	go func() {
		time.Sleep(32 * time.Millisecond)
		cl, err := netgo.Connect("localhost:9999")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		cl.Call("set", "111")
		fmt.Println(cl.Call("get", ""))
		fmt.Println(cl.Call("get", ""))
		srv.Stop()
		fmt.Println(cl.Call("get", ""))
		fmt.Println(cl.Call("get", ""))
		fmt.Println(cl.Call("get", ""))
		srv.Stop()
	}()
	fmt.Println("Serving")
	srv.Serve()
	time.Sleep(6000 * time.Second)
}

func call(c *netgo.Client, name, args string) string {
	resp, _ := c.Call(name, args)
	return resp
}

package main

import (
	"fmt"
	"time"

	"github.com/AldieNightStar/netgo"
)

func main() {
	cmds := netgo.NewCommands()
	cmds.SetInfo("a", "A LETTER")
	cmds.SetInfo("b", "B LETTER")
	cmds.SetInfo("c", "C LETTER")
	srv := netgo.NewServer(9999, cmds)
	go func() {
		time.Sleep(32 * time.Millisecond)
		cl, err := netgo.Connect("localhost:9999")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Resp: " + call(cl, "a", ""))
		fmt.Println("Resp: " + call(cl, "b", ""))
		fmt.Println("Resp: " + call(cl, "c", ""))
		srv.Stop()
	}()
	fmt.Println("Serving")
	srv.Serve()
}

func call(c *netgo.Client, name, args string) string {
	resp, _ := c.Call(name, args)
	return resp
}

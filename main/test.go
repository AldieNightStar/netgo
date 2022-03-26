package main

import (
	"fmt"
	"time"

	"github.com/AldieNightStar/netgo"
)

func main() {
	cmds := netgo.NewCommands()
	cmds.SetInfo("a", "A LETTER")
	srv := netgo.NewServer(9999, cmds)
	go func() {
		time.Sleep(32 * time.Millisecond)
		fmt.Println("Connection...")
		cl, err := netgo.Connect("localhost:9999")
		fmt.Println("Connected")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Calling...")
		resp, err := cl.Call("a", "1")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("Response: ", resp)
		time.Sleep(5 * time.Millisecond)
		srv.Stop()
	}()
	fmt.Println("Serving")
	srv.Serve()
}

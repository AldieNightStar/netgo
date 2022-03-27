package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/AldieNightStar/netgo"
)

func main() {
	go server()
	time.Sleep(1000 * time.Millisecond)
	go client()
	time.Sleep(10 * time.Minute)
}

func server() {
	cmds := netgo.NewCommands()
	cmds.SetCommand("a", func(id int, message string) string {
		return fmt.Sprintf("ECHOED(%d): %s", id, message)
	})

	netgo.NewServer(9999, cmds).HandleOnConnect(func(id int) {
		fmt.Println("Conn", id)
	}).HandleOnDisconnect(func(id int) {
		fmt.Println("Disconn", id)
	}).Serve()
}

func client() {
	cl, err := netgo.Connect("localhost:9999")
	if logError(err) {
		return
	}
	cl2, err := netgo.Connect("localhost:9999")
	if logError(err) {
		return
	}
	for i := 0; i < 100; i++ {
		fmt.Println(cl.CallOrEmpty("a", strconv.Itoa(i)))
		fmt.Println(cl2.CallOrEmpty("a", strconv.Itoa(i)))
	}
	resp, _ := cl.Call("a", "123 123 123")
	fmt.Println("RESP: ", resp)
	resp, _ = cl2.Call("a", "123 123 123")
	cl.Disconnect()
	fmt.Println("RESP: ", resp)
	cl2.Disconnect()
}

func logError(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

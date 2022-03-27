# NetGo - Socket TCP oriented lib

# Import
```go
import "github.com/AldieNightStar/netgo"
```

# Server
```go
// Let's create commands
commands := netgo.NewCommands()

// Echo command on the server
// id - client id to identify which client it is
// req - string message from client
// returns response from the server in string
commands.SetCommand("test", func(id int, req string) string {
    return "ECHO: " + req
})

// Command which returns same string all the time
commands.SetInfo("name", "HaxiDenti Server")

// Create the server
// And call Serve() when needed to start
// Server functions:
//   Serve()                     - starts the server (Run forever)
//   Stop()                      - will stop the server
//   HandleOnConnect(handler)    - handle when someone connects to the server
//   HandleOnDisconnect(handler) - handle when someone disconnects from the server
server := netgo.NewServer(7777, commands).Serve()
```


# Client
```go
// Create the client
// When client created it also joins the server. Otherwise: nil, error
client, err := netgo.Connect("localhost:7777")

// response - (string) Response from the server "test" command. Cand be "" if no such command
// err      - (error) Error message in case of any problems (No connection etc)
response, err := client.Call("test", "Hello!")

// Same as Call(...) but for now it will not return error, just an empty string
response := client.CallOrEmpty("test", "Hello!")

// Stop the connection
// To reconnect you need to create another client
client.Disconnect()
```
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
commands.SetCommand(func(req string) string {
    return "ECHO: " + req
})

// Command which returns same string all the time
commands.SetInfo("name", "HaxiDenti Server")

// Create the server
// And call Serve() when needed to start
// Server functions:
//   Serve() - starts the server (Run forever)
//   Stop()  - will stop server
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
```
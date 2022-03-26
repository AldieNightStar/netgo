package netgo

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Server struct {
	port     int
	commands Commands
	toStop   bool
}

func NewServer(port int, commands Commands) *Server {
	return &Server{
		port:     port,
		commands: commands,
		toStop:   false,
	}
}

func (s *Server) Stop() {
	s.toStop = true
}

func (s *Server) Serve() error {
	ls, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.port))
	if err != nil {
		return err
	}
	for !s.toStop {
		conn, err := ls.Accept()
		if err != nil {
			continue
		}
		go serveConn(s, conn)
		go func() {
			for !s.toStop {
				time.Sleep(time.Second)
			}
			ls.Close()
		}()
	}
	return nil
}

func (s *Server) callCommand(name, args string) string {
	cmd, ok := s.commands[name]
	if !ok {
		return ""
	}
	return cmd(args)
}

func serveConn(s *Server, conn net.Conn) {
	buf := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	for !s.toStop {
		time.Sleep(time.Millisecond)
		str, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		if strings.HasSuffix(str, "\n") {
			str = str[0 : len(str)-1]
		}
		cmd := ""
		args := ""
		{
			arr := strings.SplitN(str, " ", 2)
			if len(arr) == 1 {
				cmd = str
			} else if len(arr) >= 2 {
				cmd = arr[0]
				args = arr[1]
			}
		}
		response := s.callCommand(cmd, args)
		buf.WriteString(response + "\n")
		buf.Flush()
		if s.toStop {
			break
		}
	}
}

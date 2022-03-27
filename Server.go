package netgo

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Server struct {
	port         int
	commands     Commands
	toStop       bool
	onDisconnect Handler
	onConnect    Handler
	globCnt      int
}

func NewServer(port int, commands Commands) *Server {
	return &Server{
		port:     port,
		commands: commands,
		toStop:   false,
		globCnt:  0,
	}
}

func (s *Server) Stop() {
	s.toStop = true
}

func (s *Server) HandleOnDisconnect(f Handler) *Server {
	s.onDisconnect = f
	return s
}

func (s *Server) HandleOnConnect(f Handler) *Server {
	s.onConnect = f
	return s
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

func (s *Server) callCommand(id int, name, args string) string {
	cmd, ok := s.commands[name]
	if !ok {
		return ""
	}
	return cmd(id, args)
}

func serveConn(s *Server, conn net.Conn) {
	buf := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	clientId := s.globCnt
	s.globCnt += 1
	if s.onConnect != nil {
		s.onConnect(clientId)
	}
	for !s.toStop {
		time.Sleep(time.Millisecond)
		str, err := buf.ReadString('\n')
		if err != nil {
			if s.onDisconnect != nil {
				s.onDisconnect(clientId)
			}
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
		response := s.callCommand(clientId, cmd, args)
		buf.WriteString(response + "\n")
		buf.Flush()
	}
}

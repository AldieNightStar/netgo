package netgo

import (
	"bufio"
	"net"
	"strings"
	"time"
)

type Client struct {
	conn net.Conn
	buf  bufio.ReadWriter
}

func Connect(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
		buf:  *bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)),
	}, nil
}

func (c *Client) Call(name string, args string) (string, error) {
	c.conn.SetDeadline(time.Now().Add(5 * time.Second))
	_, err := c.buf.WriteString(name + " " + args + "\n")
	if err != nil {
		return "", err
	}
	err = c.buf.Flush()
	if err != nil {
		return "", err
	}
	time.Sleep(time.Millisecond)
	resp, err := c.buf.ReadString('\n')
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(resp, "\n") {
		resp = resp[0 : len(resp)-1]
	}
	return resp, nil
}

func (c *Client) CallOrEmpty(name string, args string) string {
	resp, err := c.Call(name, args)
	if err != nil {
		return ""
	}
	return resp
}

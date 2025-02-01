package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func (c *client) Connect() (err error) {
	if c.conn != nil {
		slog.Error("already connected")
		return nil
	}
	c.conn, err = net.DialTimeout("tcp", c.address, c.timeout)
	slog.Info(fmt.Sprintf("...connected to %s\n", c.address))
	return err
}

func (c *client) Close() error {
	if c.conn == nil {
		slog.Error("not connected")
		return nil
	}
	return c.conn.Close()
}

func (c *client) Send() error {
	_, err := io.Copy(c.conn, c.in)
	if err != nil {
		slog.Error(fmt.Sprintf("error during sending: %s", err.Error()))
	}
	return err
}

func (c *client) Receive() error {
	_, err := io.Copy(c.out, c.conn)
	if err != nil {
		slog.Error(fmt.Sprintf("error during receving: %s", err.Error()))
	}
	return err
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{address: address, timeout: timeout, in: in, out: out}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// set logger.
func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))
}

func readFlags() (time.Duration, string) {
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal(`The "host" and "port" parameters are not specified`)
	}
	port := flag.Arg(1)
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatal(`Incorrect value for the "port" parameter`)
	}
	address := net.JoinHostPort(flag.Arg(0), port)
	return *timeout, address
}

func main() {
	timeout, address := readFlags()
	telnetClient := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	if err := telnetClient.Connect(); err != nil {
		log.Fatal(fmt.Errorf("connection was failed. error: %w", err))
	}
	defer func() {
		if err := telnetClient.Close(); err != nil {
			log.Fatal("problem with client closing.", err)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// send goroutine.
	go func() {
		slog.Info("sending start")
		err := telnetClient.Send()
		if err != nil {
			slog.Error("sending problem", slog.Any("error", err))
		}
		slog.Info("sending break")
		cancel()
	}()

	// receive goroutine.
	go func() {
		slog.Info("receiving start")
		err := telnetClient.Receive()
		if err != nil {
			slog.Error("receiving problem", slog.Any("error", err))
		}
		slog.Info("receiving break")
		cancel()
	}()

	<-ctx.Done() // wait here.
}

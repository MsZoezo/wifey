package socket

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/buildkite/shellwords"
	"github.com/fatih/color"
)

type Server struct {
	socket net.Listener
}

func InitServer() (*Server, error) {
	socket, err := net.Listen("unix", "/tmp/wifey.sock")

	if err != nil {
		return nil, err
	}

	return &Server{socket}, nil
}

func (server Server) InitSignalHandler() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove("/tmp/wifey.sock")
		os.Exit(1)
	}()
}

func (server Server) Handle() {
	for {
		conn, err := server.socket.Accept()

		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			reader := bufio.NewReader(conn)
			str, err := reader.ReadString('\n')

			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}

			words, err := shellwords.Split(str)

			if err != nil {
				slog.Error(err.Error())
				os.Exit(1)
			}

			slog.Info("Received new request on socket.", slog.Any("Request", words))

			switch words[0] {
			case "send":
				fmt.Fprintln(conn, "Send command received")
			case "receive":
				fmt.Fprintln(conn, "Receive command received")
			default:
				fmt.Fprintln(conn, color.RedString("૮⸝⸝> ﻌ <⸝⸝ა Unknown command, ignoring..."))
			}

		}(conn)
	}
}

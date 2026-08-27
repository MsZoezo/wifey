package core

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Socket struct {
	socket net.Listener
}

func InitSocket() (*Socket, error) {
	socket, err := net.Listen("unix", "/tmp/wifey.sock")

	if err != nil {
		return nil, err
	}

	return &Socket{socket}, nil
}

func (socket Socket) InitSignalHandler() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove("/tmp/wifey.sock")
		os.Exit(1)
	}()
}

func (socket Socket) Handle() {
	for {
		conn, err := socket.socket.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			defer conn.Close()

			buf := make([]byte, 4096)

			n, err := conn.Read(buf)

			if err != nil {
				log.Fatal(err)
			}

			_, err = conn.Write(buf[:n])

			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}

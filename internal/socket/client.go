package socket

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	socket net.Conn
}

func InitClient() (*Client, error) {
	socket, err := net.Dial("unix", "/tmp/wifey.sock")

	if err != nil {
		return nil, err
	}

	return &Client{socket}, nil
}

func (client Client) Send(str string) (res string, err error) {
	_, err = fmt.Fprintln(client.socket, str)

	if err != nil {
		return
	}

	reader := bufio.NewReader(client.socket)
	res, err = reader.ReadString('\n')

	return
}

func (client Client) Close() {
	client.socket.Close()
}

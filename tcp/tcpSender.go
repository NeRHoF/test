package tcp

import (
	"net"
	"os"
	test_entities "test/testEntities"
)

const (
	protocolName = "tcp"
	listenerHost = "127.0.0.1:8081"
)

type tcpSender struct {
	TcpConnection net.Conn
}

func NewTcpSender() test_entities.TcpSendRpc {
	conn, err := net.Dial(protocolName, listenerHost)
	if err != nil {
		os.Exit(1)
	}
	return tcpSender{TcpConnection: conn}
}

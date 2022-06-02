package tpuService

import (
	"errors"
	"net"
)

type Sender struct {
	Addr *net.UDPAddr
	conn *net.UDPConn
}

func (sender *Sender) SendTransaction(transaction []byte) (int, error) {
	n, err := sender.conn.WriteTo(transaction, sender.Addr)
	if err != nil {
		return 0, errors.New("Send transaction error: " + err.Error())
	}
	return n, nil
}

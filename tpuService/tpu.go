package tpuService

import (
	"github.com/gagliardetto/solana-go/rpc"
	tpuService "github.com/karneges/validators-proto"
	"log"
	"net"
)

type ClusterNodeWithSchedule struct {
	*rpc.GetClusterNodesResult
	CountOfSchedules int
}

type Tpu struct {
	sortedTpu []*tpuService.Validator
	conn      *net.UDPConn
}

func New(validators []*tpuService.Validator) *Tpu {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 1234})
	if err != nil {
		log.Printf(err.Error())
	}
	return &Tpu{
		sortedTpu: validators,
		conn:      conn,
	}
}
func (tpu *Tpu) GetValidators() []Sender {
	var senders []Sender
	for _, validator := range tpu.sortedTpu {
		addr, _ := net.ResolveUDPAddr("udp4", validator.Tpu)
		senders = append(senders, Sender{
			Addr: addr,
			conn: tpu.conn,
		})
	}
	return senders
}

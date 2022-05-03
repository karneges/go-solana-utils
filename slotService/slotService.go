package slotService

import (
	"context"
	"github.com/gagliardetto/solana-go/rpc"
	"time"
)

type SlotIdx struct {
	CurrentIdx      uint64
	CurrentIdxCh    <-chan uint64
	client          *rpc.Client
	pollingInterval time.Duration
}

func New(client *rpc.Client, pollingInterval time.Duration) *SlotIdx {

	slotIdx := &SlotIdx{
		client:          client,
		CurrentIdx:      0,
		pollingInterval: pollingInterval,
	}
	slotIdx.runIndexWatcher()
	return slotIdx
}

func (slotIdx *SlotIdx) runIndexWatcher() {
	ch := make(chan uint64)
	go func() {
		for {
			idx, err := slotIdx.client.GetEpochInfo(context.TODO(), rpc.CommitmentRecent)
			if err != nil {
				println(err.Error())
				continue
			}
			ch <- idx.SlotIndex
			time.Sleep(slotIdx.pollingInterval)
		}
	}()
	slotIdx.CurrentIdxCh = ch
	slotIdx.CurrentIdx = <-ch
	go func() {
		for {
			slotIdx.CurrentIdx = <-ch
		}
	}()

}

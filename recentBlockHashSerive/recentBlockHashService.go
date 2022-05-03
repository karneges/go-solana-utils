package recentBlockHashSerive

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"time"
)

type RecentBlockHash struct {
	Hash  solana.Hash
	HasCh chan solana.Hash
}

func New(client *rpc.Client, pollingInterval time.Duration) *RecentBlockHash {
	recentBlockHash := &RecentBlockHash{}
	recentBlockHashCh := make(chan solana.Hash)
	go func() {
		for {
			res, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentRecent)
			if err != nil {
				panic(err)
			}
			recentBlockHashCh <- res.Value.Blockhash
			time.Sleep(pollingInterval)
		}

	}()
	recentBlockHash.HasCh = recentBlockHashCh
	recentBlockHash.Hash = <-recentBlockHashCh
	go func() {
		for newBlockHash := range recentBlockHashCh {
			recentBlockHash.Hash = newBlockHash
		}
	}()
	return recentBlockHash
}

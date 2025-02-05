package app

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	basincapnp "github.com/tablelandnetwork/basin-cli/pkg/capnp"
)

// DealInfo represents information about a deal.
type DealInfo struct {
	CID         string `json:"cid"`
	Timestamp   int64  `json:"timestamp"`
	Size        uint32 `json:"size"`
	IsArchived  bool   `json:"is_archived"`
	CacheExpiry string `json:"cache_expiry"`
}

// BasinProvider ...
type BasinProvider interface {
	Create(context.Context, string, string, basincapnp.Schema, common.Address, int64) (bool, error)
	List(context.Context, common.Address) ([]string, error)
	Deals(context.Context, string, string, uint32, uint64, Timestamp, Timestamp) ([]DealInfo, error)
	LatestDeals(context.Context, string, string, uint32, Timestamp, Timestamp) ([]DealInfo, error)
	Reconnect() error
}

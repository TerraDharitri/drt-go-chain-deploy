package data

import (
	"github.com/TerraDharitri/drt-go-chain/genesis/data"
	"github.com/TerraDharitri/drt-go-chain/sharding"
)

// GeneratorOutput represents the structure that will contain aggregated generated data
type GeneratorOutput struct {
	ValidatorBlsKeys []*BlsKey
	ObserverBlsKeys  []*BlsKey
	WalletKeys       []*WalletKey
	AdditionalKeys   []*WalletKey
	InitialAccounts  []data.InitialAccount
	InitialNodes     []*sharding.InitialNode
	DelegatorKeys    []*WalletKey
}

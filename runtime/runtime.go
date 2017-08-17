package runtime

import (
	"math/big"

	"github.com/end-r/firevm/state"
)

// Config for FireVM instance
type Config struct {
	ChainConfig *ChainConfig
	Difficulty  *big.Int
	Origin      Address
	Coinbase    Address
	BlockNumber *big.Int
	Time        *big.Int
	FuelLimit   uint64
	FuelPrice   *big.Int
	Value       *big.Int
	DisableJit  bool // "disable" so it's enabled by default
	Debug       bool

	State     *state.State
	GetHashFn func(n uint64) Hash
}

func (c *Config) setDefaults() {

}

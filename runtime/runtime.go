package runtime

import (
	"demos/ansible/go-ethereum-master/common"
	"demos/ansible/go-ethereum-master/core/vm"
	"demos/ansible/params"
	"ender/ans/state"
	"math/big"

	"github.com/end-r/firevm"
)

// Config for FireVM instance
type Config struct {
	ChainConfig *params.ChainConfig
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
	EVMConfig   vm.Config

	State     *state.StateDB
	GetHashFn func(n uint64) common.Hash
}

func (c *Config) setDefaults() {

}

// Call executes the code given by the contract's address. It will return the
// EVM's return value or an error if it failed.
//
// Call, unlike Execute, requires a config and also requires the State field to
// be set.
func Call(address firevm.Address, input []byte, cfg *Config) ([]byte, uint64, error) {
	setDefaults(cfg)

	vmenv := NewEnv(cfg)

	sender := cfg.State.GetOrNewStateObject(cfg.Origin)
	// Call the code with the given configuration.
	ret, leftOverGas, err := vmenv.Call(
		sender,
		address,
		input,
		cfg.GasLimit,
		cfg.Value,
	)

	return ret, leftOverGas, err
}

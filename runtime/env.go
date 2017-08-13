package runtime

import (
	"demos/ansible/go-ethereum-master/common"
	"demos/ansible/go-ethereum-master/core"
	"demos/ansible/go-ethereum-master/core/vm"
	"math/big"
)

func NewEnv(cfg *Config) *vm.EVM {
	context := vm.Context{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		GetHash:     func(uint64) common.Hash { return common.Hash{} },

		Origin:      cfg.Origin,
		Coinbase:    cfg.Coinbase,
		BlockNumber: cfg.BlockNumber,
		Time:        cfg.Time,
		Difficulty:  cfg.Difficulty,
		GasLimit:    new(big.Int).SetUint64(cfg.GasLimit),
		GasPrice:    cfg.GasPrice,
	}

	return vm.NewEVM(context, cfg.State, cfg.ChainConfig, cfg.EVMConfig)
}

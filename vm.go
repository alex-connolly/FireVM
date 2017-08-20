package firevm

import (
	"github.com/end-r/firevm/ops"
	"github.com/end-r/vmgen"
)

const (
	vmPath = "firevm.vm"
)

var (
	fuels    = map[string]vmgen.FuelFunction{}
	executes = map[string]vmgen.ExecuteFunction{
		// arithmetic operations
		"ADD":    ops.Add,
		"SUB":    ops.Sub,
		"DIV":    ops.Div,
		"MUL":    ops.Mul,
		"MOD":    ops.Mod,
		"ADDMOD": ops.AddMod,
		"MULMOD": ops.MulMod,
		// comparison operations
		"LT":     ops.Lt,
		"GT":     ops.Gt,
		"SLT":    ops.SLt,
		"SGT":    ops.SGt,
		"EQ":     ops.Eq,
		"ISZERO": ops.IsZero,
		// bitwise operations
		"AND": ops.And,
		"OR":  ops.Or,
		"NOT": ops.Not,
		"XOR": ops.Xor,
		// block operations
		"HASH":       ops.Blockhash,
		"COINBASE":   ops.Coinbase,
		"TIMESTAMP":  ops.Timestamp,
		"NUMBER":     ops.Number,
		"DIFFICULTY": ops.Difficulty,
		"FUELLIMIT":  ops.FuelLimit,
		// environment operations
		"ADDRESS":      ops.Address,
		"ORIGIN":       ops.Origin,
		"BALANCE":      ops.Balance,
		"CALLER":       ops.Caller,
		"CALLVALUE":    ops.CallValue,
		"CALLDATALOAD": ops.CallDataLoad,
		"CALLDATASIZE": ops.CallDataSize,
		"CALLDATACOPY": ops.CallDataCopy,
		"CODESIZE":     ops.CodeSize,
		"CODECOPY":     ops.CodeCopy,
		"FUELPRICE":    ops.FuelPrice,
		"EXTCODESIZE":  ops.ExtCodeSize,
		"EXTCODECOPY":  ops.ExtCodeCopy,
		// crypto operations
		"SHA3": ops.SHA3,
		//stack operations
		"POP.": ops.Pop,
		"PUSH": ops.Push,
		"DUP":  ops.Dup,
		"SWAP": ops.Swap,
		// external operations
		"LOG": ops.Log,
		// flow operations
		"JUMP":     ops.Jump,
		"JUMPI":    ops.JumpI,
		"PC":       ops.PC,
		"FUEL":     ops.Fuel,
		"MEMSIZE":  ops.MemSize,
		"JUMPDEST": ops.JumpDest,
		// memory operations
		"SET":   ops.Set,
		"GET":   ops.Get,
		"LOAD":  ops.Load,
		"STORE": ops.Store,
	}
)

// NewVM returns a new FireVM instance
func NewVM() *vmgen.VM {
	return vmgen.CreateVM(vmPath, nil, executes, fuels, nil)
}

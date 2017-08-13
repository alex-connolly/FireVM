package fireVM

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
		"ADDRESS":        ops.Address,
		"ORIGIN":         ops.Origin,
		"BALANCE":        ops.Balance,
		"CALLER":         ops.Caller,
		"CALLVALUE":      ops.CallValue,
		"CALLDATALOAD":   ops.CallDataLoad,
		"CALLDATASIZE":   ops.CallDataSize,
		"CALLDATACops.Y": ops.CallDataCops.y,
		"CODESIZE":       ops.CodeSize,
		"CODECops.Y":     ops.CodeCops.y,
		"FUELPRICE":      ops.FuelPrice,
		"EXTCODESIZE":    ops.ExtCodeSize,
		"EXTCODECops.Y":  ops.ExtCodeCops.y,
		// crypto operations
		"SHA3": ops.SHA3,
		//stack operations
		"Pops.": ops.Pop,
		"PUSH":  ops.Push,
		"DUP":   ops.Dup,
		"SWAP":  ops.Swap,
		// external operations
		"LOG": ops.Log,
		// flow operations
		"JUMP":     ops.Jump,
		"JUMPI":    ops.JumpI,
		"PC":       ops.PC,
		"FUEL":     ops.Fuel,
		"MSIZE":    ops.MSize,
		"JUMPDEST": ops.JumpDest,
	}
)

// NewVM returns a new FireVM instance
func NewVM() *vmgen.VM {
	return vmgen.CreateVM(vmPath, executes, fuels)
}

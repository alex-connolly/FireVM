package fireVM

import (
	"axia/fireVM/ops"

	"github.com/end-r/vmgen"
)

const (
	vmPath = "firevm.vm"
)

var (
	fuels    = map[string]vmgen.FuelFunction{}
	executes = map[string]vmgen.ExecuteFunction{
		// arithmetic ops.Operations
		"ADD":    ops.OpAdd,
		"SUB":    ops.OpSub,
		"DIV":    ops.OpDiv,
		"MUL":    ops.OpMul,
		"MOD":    ops.OpMod,
		"ADDMOD": ops.OpAddMod,
		"MULMOD": ops.OpMulMod,
		// comparison ops.Operations
		"LT":     ops.OpLt,
		"GT":     ops.OpGt,
		"SLT":    ops.OpSLt,
		"SGT":    ops.OpSGt,
		"EQ":     ops.OpEq,
		"ISZERO": ops.OpIsZero,
		// bitwise ops.Operations
		"AND": ops.OpAnd,
		"OR":  ops.OpOr,
		"NOT": ops.OpNot,
		"XOR": ops.OpXor,
		// block ops.Operations
		"HASH":       ops.OpBlockhash,
		"COINBASE":   ops.OpCoinbase,
		"TIMESTAMP":  ops.OpTimestamp,
		"NUMBER":     ops.OpNumber,
		"DIFFICULTY": ops.OpDifficulty,
		"FUELLIMIT":  ops.OpFuelLimit,
		// environment ops.Operations
		"ADDRESS":          ops.OpAddress,
		"ORIGIN":           ops.OpOrigin,
		"BALANCE":          ops.OpBalance,
		"CALLER":           ops.OpCaller,
		"CALLVALUE":        ops.OpCallValue,
		"CALLDATALOAD":     ops.OpCallDataLoad,
		"CALLDATASIZE":     ops.OpCallDataSize,
		"CALLDATACops.OpY": ops.OpCallDataCops.Opy,
		"CODESIZE":         ops.OpCodeSize,
		"CODECops.OpY":     ops.OpCodeCops.Opy,
		"FUELPRICE":        ops.OpFuelPrice,
		"EXTCODESIZE":      ops.OpExtCodeSize,
		"EXTCODECops.OpY":  ops.OpExtCodeCops.Opy,
		// crypto ops.Operations
		"SHA3": ops.OpSHA3,
		//stack ops.Operations
		"Pops.Op": ops.OpPop,
		"PUSH":    ops.OpPush,
		"DUP":     ops.OpDup,
		"SWAP":    ops.OpSwap,
		// external ops.Operations
		"LOG": ops.OpLog,
		// flow ops.Operations
		"JUMP":     ops.OpJump,
		"JUMPI":    ops.OpJumpI,
		"PC":       ops.OpPC,
		"FUEL":     ops.OpFuel,
		"MSIZE":    ops.OpMSize,
		"JUMPDEST": ops.OpJumpDest,
	}
)

// NewVM returns a new FireVM instance
func NewVM() *vmgen.VM {
	vm := vmgen.CreateVM(vmPath, executes, fuels)
	return vm
}

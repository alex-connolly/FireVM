package fireVM

import (
	"github.com/end-r/vmgen"
)

const (
	vmPath = "firevm.vm"
)

var (
	fuels    = map[string]vmgen.FuelFunction{}
	executes = map[string]vmgen.ExecuteFunction{
		// arithmetic operations
		"ADD":    opAdd,
		"SUB":    opSub,
		"DIV":    opDiv,
		"MUL":    opMul,
		"MOD":    opMod,
		"ADDMOD": opAddMod,
		"MULMOD": opMulMod,
		// comparison operations
		"LT":     opLt,
		"GT":     opGt,
		"SLT":    opSLt,
		"SGT":    opSGt,
		"EQ":     opEq,
		"ISZERO": opIsZero,
		// bitwise operations
		"AND": opAnd,
		"OR":  opOr,
		"NOT": opNot,
		"XOR": opXor,
		// block operations
		"HASH":       opBlockhash,
		"COINBASE":   opCoinbase,
		"TIMESTAMP":  opTimestamp,
		"NUMBER":     opNumber,
		"DIFFICULTY": opDifficulty,
		"FUELLIMIT":  opFuelLimit,
		// environment operations
		"ADDRESS":      opAddress,
		"ORIGIN":       opOrigin,
		"BALANCE":      opBalance,
		"CALLER":       opCaller,
		"CALLVALUE":    opCallValue,
		"CALLDATALOAD": opCallDataLoad,
		"CALLDATASIZE": opCallDataSize,
		"CALLDATACOPY": opCallDataCopy,
		"CODESIZE":     opCodeSize,
		"CODECOPY":     opCodeCopy,
		"FUELPRICE":    opFuelPrice,
		"EXTCODESIZE":  opExtCodeSize,
		"EXTCODECOPY":  opExtCodeCopy,
		// crypto operations
		"SHA3": opSHA3,
		//stack operations
		"POP":  opPop,
		"PUSH": opPush,
		"DUP":  opDup,
		"SWAP": opSwap,
		// external operations
		"LOG": opLog,
		// flow operations
		"JUMP":     opJump,
		"JUMPI":    opJumpI,
		"PC":       opPC,
		"FUEL":     opFuel,
		"MSIZE":    opMSize,
		"JUMPDEST": opJumpDest,
	}
)

// NewVM returns a new FireVM instance
func NewVM() *vmgen.VM {
	vm := vmgen.CreateVM(vmPath, executes, fuels)
	return vm
}

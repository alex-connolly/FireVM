package ops

import (
	"math/big"

	"github.com/end-r/vmgen"
)

// Lt ...
func Lt(vm *vmgen.VM) {
	executeLt(vm.Stack)
}

func executeLt(s *vmgen.Stack) {
	x := createBigInt(s.Pop(1))
	y := createBigInt(s.Pop(1))
	if x.Cmp(y) < 0 {
		// in EVM this comes from the int pool
		// TODO: why?
		s.Push(new(big.Int).SetUint64(1).Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

// Gt ...
func Gt(vm *vmgen.VM) {
	executeGt(vm.Stack)
}

func executeGt(s *vmgen.Stack) {
	x := createBigInt(s.Pop(1))
	y := createBigInt(s.Pop(1))
	if x.Cmp(y) > 0 {
		// in EVM this comes from the int pool
		// TODO: why?
		s.Push(new(big.Int).SetUint64(1).Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

// SLt ...
func SLt(vm *vmgen.VM) {

}

// SGt ...
func SGt(vm *vmgen.VM) {

}

// Eq ...
func Eq(vm *vmgen.VM) {
	executeEq(vm.Stack)
}

func executeEq(s *vmgen.Stack) {
	x := createBigInt(s.Pop(1))
	y := createBigInt(s.Pop(1))
	if x.Cmp(y) == 0 {
		// in EVM this comes from the int pool
		// TODO: why?
		s.Push(new(big.Int).SetUint64(1).Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

// IsZero ...
func IsZero(vm *vmgen.VM) {
	executeIsZero(vm.Stack)
}

func executeIsZero(s *vmgen.Stack) {
	x := createBigInt(s.Pop(1))
	if x.Sign() > 0 {
		s.Push(new(big.Int).Bytes())
	} else {
		// in EVM this comes from the int pool
		// TODO: why?
		s.Push(new(big.Int).SetUint64(1).Bytes())
	}
}

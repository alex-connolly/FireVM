package ops

import (
	"math/big"

	"github.com/end-r/vmgen"
)

// Pop ...
func Pop(vm *vmgen.VM) {
	executePop(vm.Stack)
}

func executePop(s *vmgen.Stack) {
	s.Pop()
}

// Push ...
func Push(vm *vmgen.VM) {
	executePush(vm.Stack, vm.Input)
}

func executePush(s *vmgen.Stack, i vmgen.Input) {
	size := i.Code().Next(1)
	pushSize := int(new(big.Int).SetBytes(size).Int64())
	s.Push(i.Code().Next(pushSize))
}

// Dup ...
func Dup(vm *vmgen.VM) {
	executeDup(vm.Stack, vm.Input)
}

func executeDup(s *vmgen.Stack, i vmgen.Input) {
	size := int(new(big.Int).SetBytes(i.Code().Next(1)).Int64())
	dupSize := int(new(big.Int).SetBytes(i.Code().Next(size)).Int64())
	s.Dup(dupSize)
}

// Swap ...
func Swap(vm *vmgen.VM) {
	executeSwap(vm.Stack, vm.Input)
}

func executeSwap(s *vmgen.Stack, i vmgen.Input) {
	size := int(new(big.Int).SetBytes(i.Code().Next(1)).Int64())
	swapSize := int(new(big.Int).SetBytes(i.Code().Next(size)).Int64())
	s.Swap(swapSize)
}

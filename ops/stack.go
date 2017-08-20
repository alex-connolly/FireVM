package ops

import "github.com/end-r/vmgen"

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
	//size := i.Code().Next(1)
	//vm.Stack.Push(vm.Next(size.Int64()))
}

// Dup ...
func Dup(vm *vmgen.VM) {
	executeDup(vm.Stack, vm.Input)
}

func executeDup(s *vmgen.Stack, i vmgen.Input) {
	//size := i.Code().Next(1)
	//s.Dup(size)
}

// Swap ...
func Swap(vm *vmgen.VM) {
	executeSwap(vm.Stack, vm.Input)
}

func executeSwap(s *vmgen.Stack, i vmgen.Input) {
	//size := i.Code().Next(1)
	//s.Swap(size)
}

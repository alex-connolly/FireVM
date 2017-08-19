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
	executePush(vm.Stack, vm.Contract)
}

func executePush(s *vmgen.Stack, c *vmgen.Contract) {
	size := vm.Next(1)
	vm.Stack.Push(vm.Next(size.Int64()))
}

// Dup ...
func Dup(vm *vmgen.VM) {
	executeDup(vm.Stack, vm.Contract)
}

func executeDup(s *vmgen.Stack, c *vmgen.Contract) {
	size := c.Next(1)
	s.Dup(size)
}

// Swap ...
func Swap(vm *vmgen.VM) {
	executeSwap(vm.Stack, vm.Contract)
}

func executeSwap(s *vmgen.Stack, c *vmgen.Contract) {
	size := c.Next(1)
	s.Swap(size)
}

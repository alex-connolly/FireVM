package ops

import "github.com/end-r/vmgen"

// Pop ...
func Pop(vm *vmgen.VM) {
	vm.Stack.Pop(1)
}

// Push ...
func Push(vm *vmgen.VM) {
	executePush(vm)
}

func executePush(vm *vmgen.VM) {

}

// Dup ...
func Dup(vm *vmgen.VM) {
	executeDup(vm.Stack, vm.Input)
}

func executePush(vm *vmgen.VM) {
	size := createBigInt(vm.Next(1))
	vm.Stack.Push(vm.Next(size))
}

// Swap ...
func Swap(vm *vmgen.VM) {

}

func executeSwap(vm *vmgen.VM) {
	size := createBigInt(vm.Next(1))
	vm.Stack.Push(vm.Next(size))
}

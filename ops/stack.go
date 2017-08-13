package ops

import "github.com/end-r/vmgen"

// Pop ...
func Pop(vm *vmgen.VM, params []byte) {
	vm.Stack.Pop(1)
}

// Push ...
func Push(vm *vmgen.VM, params []byte) {

}

// Dup ...
func Dup(vm *vmgen.VM, params []byte) {

}

// Swap ...
func Swap(vm *vmgen.VM, params []byte) {

}

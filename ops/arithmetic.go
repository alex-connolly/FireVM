package ops

import "github.com/end-r/vmgen"

func OpAdd(vm *vmgen.VM, params []byte) {
	a := vm.Stack.Pop(1)
	b := vm.Stack.Pop(1)
	c := a + b
	vm.Stack.Push(c)
}

func OpSub(vm *vmgen.VM, params []byte) {

}

func OpMul(vm *vmgen.VM, params []byte) {

}

func OpDiv(vm *vmgen.VM, params []byte) {

}

func OpMod(vm *vmgen.VM, params []byte) {

}

func OpAddMod(vm *vmgen.VM, params []byte) {

}

func OpMulMod(vm *vmgen.VM, params []byte) {

}

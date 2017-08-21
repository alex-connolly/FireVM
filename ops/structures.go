package ops

import "github.com/end-r/vmgen"

// Array turns the previous offset byte arrays into an array
func Array(vm *vmgen.VM) {
	executeArray(vm.Stack)
}

func executeArray(s *vmgen.Stack) {

}

func Map(vm *vmgen.VM) {
	executeMap(vm.Stack)
}

func executeMap(s *vmgen.Stack) {

}

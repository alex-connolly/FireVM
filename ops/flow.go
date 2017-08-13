package ops

import "github.com/end-r/vmgen"

// Jump ...
func Jump(vm *vmgen.VM) {

}

// JumpI ...
func JumpI(vm *vmgen.VM) {

}

// PC ...
func PC(vm *vmgen.VM) {

}

// Fuel ...
func Fuel(vm *vmgen.VM) {

}

func executeFuel(s *vmgen.Stack) {

}

// MSize ...
func MSize(vm *vmgen.VM) {
	executeMSize(vm.Stack)
}

func executeMSize(s *vmgen.Stack) {

}

// JumpDest marks a location as a valid jump destination
func JumpDest(vm *vmgen.VM) {
	// do nothing
}

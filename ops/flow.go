package ops

import (
	"demos/ansible/go-ethereum-master/core/vm"
	"math/big"

	"github.com/end-r/vmgen"
)

// Jump ...
func Jump(vm *vmgen.VM) {

}

// JumpI ...
func JumpI(vm *vmgen.VM) {

}

// PC ...
func PC(vm *vmgen.VM) {
	executePC(vm.Stack, vm.PC)
}

func executePC(s *vmgen.Stack, c uint64) {
	a := new(big.Int).setUInt64(c)
	s.Push(a)
}

// Fuel ...
func Fuel(vm *vmgen.VM) {
	executeFuel(vm.Stack, vm.Contract)
}

func executeFuel(s *vmgen.Stack, c vmgen.Contract) {
	a := new(big.Int).setInt64(c.Fuel)
	s.Push(a)
}

// MSize ...
func MSize(vm *vmgen.VM) {
	executeMSize(vm.Stack)
}

func executeMSize(s *vmgen.Stack, m vm.Memory) {
	a := new(big.Int).setInt64(int64(m.Length()))
	s.Push(a)
}

// JumpDest marks a location as a valid jump destination
func JumpDest(vm *vmgen.VM) {
	// do nothing
}

// Stop ...
func Stop(vm *vmgen.VM) {

}

// Return ...
func Return(vm *vmgen.VM) {

}

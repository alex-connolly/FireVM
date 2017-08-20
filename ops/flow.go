package ops

import (
	"math/big"

	"github.com/end-r/vmgen"
)

// Jump ...
func Jump(vm *vmgen.VM) {
	executeJump(vm.Stack)
}

func executeJump(s *vmgen.Stack) {
	//pos := s.Pop()

}

// JumpI ...
func JumpI(vm *vmgen.VM) {
	executeJumpI()
}

func executeJumpI() {

}

// PC ...
func PC(vm *vmgen.VM) {
	//	executePC(vm.Stack, vm.PC)
}

func executePC(s *vmgen.Stack, c uint64) {
	a := new(big.Int).SetUint64(c)
	s.Push(a.Bytes())
}

// Fuel ...
func Fuel(vm *vmgen.VM) {
	executeFuel(vm.Stack, vm.Input)
}

func executeFuel(s *vmgen.Stack, i vmgen.Input) {
	//a := new(big.Int).setInt64(i.Fuel())
	//s.Push(a)
}

// MSize ...
func MSize(vm *vmgen.VM) {
	//executeMSize(vm.Stack)
}

func executeMSize(s *vmgen.Stack, m vmgen.Memory) {
	/*a := new(big.Int).SetInt64(int64(m.Length()))
	s.Push(a)*/
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

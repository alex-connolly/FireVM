package ops

import (
	"github.com/end-r/vmgen"
)

// Get ...
func Get(vm *vmgen.VM) {
	executGet(vm.Stack, vm.Memory["memory"])
}

func executeGet(s *vmgen.Stack, m vmgen.Memory) {
	offset := s.Pop(1)
	val := bigInt(m.Get(offset.Int64(), 32))
	s.Push(val)
}

// MStore ...
func Set(vm *vmgen.VM) {
	executeSet(vm.Stack, vm.Memory["memory"])
}

func executeSet(s *vmgen.Stack, m vmgen.Memory) {
	start, val := s.Pop(1), s.Pop(1)
	m.Set(start.Uint64(), 32, math.PaddedBigBytes(val, 32))
}

// Load ...
func Load(vm *vmgen.VM) {
	executeLoad(vm.Stack, vm.Contract, vm.State)
}

func executeLoad(s *vmgen.Stack, c *vmgen.Contract, state vmgen.State) {
	s.Push(val)
}

// Store ...
func Store(vm *vmgen.VM) {
	executeStore(vm.Stack, vm.Contract, vm.State)
}

func executeStore(s *vmgen.Stack, c *vmgen.Contract, state vmgen.State) {

}

package ops

import (
	"github.com/end-r/vmgen"
)

// Get ...
func Get(vm *vmgen.VM) {
	//executeGet(vm.Stack, vm.Memory["memory"])
}

func executeGet(s *vmgen.Stack, m vmgen.Memory) {
	/*offset := s.Pop()
	val := bigInt(m.Get(offset.Int64(), 32))
	s.Push(val)*/
}

// Set ...
func Set(vm *vmgen.VM) {
	//executeSet(vm.Stack, vm.Memory["memory"])
}

func executeSet(s *vmgen.Stack, m vmgen.Memory) {
	//start, val := s.Pop(), s.Pop()
	//m.Set(start.Uint64(), 32, math.PaddedBigBytes(val, 32))
}

// Load ...
func Load(vm *vmgen.VM) {
	executeLoad(vm.Stack, vm.Input, vm.State)
}

func executeLoad(s *vmgen.Stack, i vmgen.Input, state vmgen.State) {
	//s.Push(val)
}

// Store ...
func Store(vm *vmgen.VM) {
	executeStore(vm.Stack, vm.Input, vm.State)
}

func executeStore(s *vmgen.Stack, i vmgen.Input, state vmgen.State) {

}

package ops

import (
	"math/big"

	"github.com/end-r/firevm/memory"
	"github.com/end-r/vmgen"
)

func MemSize(vm *vmgen.VM) {
	executeMemSize(vm.Stack, vm.Memory["memory"])
}

func executeMemSize(s *vmgen.Stack, m vmgen.Memory) {
	fm := m.(memory.FireMemory)
	size := new(big.Int).SetInt64(int(fm.Size()))
	s.Push(size.Bytes())
}

// Get from memory
func Get(vm *vmgen.VM) {
	executeGet(vm.Stack, vm.Memory["memory"])
}

func executeGet(s *vmgen.Stack, m vmgen.Memory) {
	offset := bigInt(s.Pop())
	fm := m.(memory.FireMemory)
	val := fm.Get(offset.Int64(), 32)
	s.Push(val)
}

// Set data to memory
func Set(vm *vmgen.VM) {
	executeSet(vm.Stack, vm.Memory["memory"])
}

func executeSet(s *vmgen.Stack, m vmgen.Memory) {
	start := s.Pop()
	val := s.Pop()
	fm := m.(memory.FireMemory)
	fm.Set(new(big.Int).SetBytes(start).Uint64(), 32, val)
	//m.Set(start.Uint64(), 32, math.PaddedBigBytes(val, 32))
}

// Load ...
func Load(vm *vmgen.VM) {
	executeLoad(vm.Stack, vm.Input, vm.State)
}

func executeLoad(s *vmgen.Stack, i vmgen.Input, state vmgen.State) {
	/*key := axia.HashBytes(s.Pop())
	val := state.Get(i.Address(), key)
	s.Push(val)*/
}

// Store data in state
func Store(vm *vmgen.VM) {
	executeStore(vm.Stack, vm.Input, vm.State)
}

func executeStore(s *vmgen.Stack, i vmgen.Input, state vmgen.State) {
	/*key := axia.HashBytes(s.Pop())
	val := axia.HashBytes(s.Pop())
	state.Set(i.Address(), key, val)*/
}

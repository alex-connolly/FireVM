package ops

import (
	"demos/ansible/go-ethereum-master/common"
	"math/big"

	"github.com/end-r/vmgen"
)

// MLoad ...
func MLoad(vm *vmgen.VM) {
	executeMLoad(vm.Stack, vm.Memory["memory"])
}

func executeMLoad(s *vmgen.Stack, m vmgen.Memory) {
	offset := s.Pop(1)
	val := new(big.Int).SetBytes(m.Get(offset.Int64(), 32))
	s.Push(val)
}

// MStore ...
func MStore(vm *vmgen.VM) {
	executeMStore(vm.Stack, vm.Memory["memory"])
}

func executeMStore(s *vmgen.Stack, m vmgen.Memory) {
	start, val := s.Pop(1), s.Pop(1)
	m.Set(start.Uint64(), 32, math.PaddedBigBytes(val, 32))
}

// SLoad ...
func SLoad(vm *vmgen.VM) {
	executeSLoad(vm.Stack, vm.Contract, vm.State)
}

func executeSLoad(s *vmgen.Stack, c *vmgen.Contract, state vmgen.State) {
	loc := common.BigToHash(s.Pop(1))
	val := state.GetState(c.Address(), loc).Big()
	s.Push(val)
}

// SStore ...
func SStore(vm *vmgen.VM) {
	executeSStore(vm.Stack, vm.Contract, vm.State)
}

func executeSStore(s *vmgen.Stack, c *vmgen.Contract, state vmgen.State) {
	loc := util.BigToHash(stack.pop())
	val := stack.pop()
	evm.StateDB.SetState(contract.Address(), loc, common.BigToHash(val))
}

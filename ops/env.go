package ops

import (
	"github.com/end-r/go-axia/axia"

	"math/big"

	"github.com/end-r/vmgen"
)

// Address ...
func Address(vm *vmgen.VM) {
	executeAddress(vm.Stack, vm.Contract)
}

func executeAddress(s *vmgen.Stack, c *vmgen) {
	s.Push(c.Address.Bytes())
}

// Balance ...
func Balance(vm *vmgen.VM) {
	executeBalance(vm.Stack, vm.State)
}

func executeBalance(s *vmgen.Stack, state vmgen.State) {
	a := axia.BytesToAddress(s.Pop())
	balance := state.GetBalance(a)
	s.Push(balance.Bytes())
}

// Origin ...
func Origin(vm *vmgen.VM) {
	executeOrigin(vm.Stack, vm.Environment)
}

func executeOrigin(s *vmgen.Stack, a vmgen.Address) {
	s.Push(a.Bytes())
}

// Caller ...
func Caller(vm *vmgen.VM) {
	executeCaller(vm.Stack, vm.Contract)
}

func executeCaller(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(c.Caller.Bytes())
}

// CallValue ...
func CallValue(vm *vmgen.VM) {
	executeCallValue(vm.Stack, vm.Contract)
}

func executeCallValue(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(c.Value())
}

// CallDataLoad ...
func CallDataLoad(vm *vmgen.VM) {
	executeCallDataLoad(vm.Stack, vm.Contract)
}

func executeCallDataLoad(s *vmgen.Stack, c *vmgen.Contract) {
	cd := getData(c.Input, s.Pop(), new(big.Int).SetInt64(32))
	s.Push(cd)
}

// CallDataSize ...
func CallDataSize(vm *vmgen.VM) {
	executeCallDataSize(vm.Stack, vm.Contract)
}

func executeCallDataSize(s *vmgen.Stack, c *vmgen.Contract) {
	a := new(big.Int).SetInt64(int64(len(c.Input)))
	s.Push(a.Bytes())
}

// CallDataCopy ...
func CallDataCopy(vm *vmgen.VM) {
	//executeCallDataCopy(vm.Stack, vm.Memory["memory"], vm.Contract)
}

func executeCallDataCopy(s *vmgen.Stack, m vmgen.Memory, c vmgen.Contract) {
	mOff := s.Pop()
	cOff := s.Pop()
	l := s.Pop()
	m.Set(mOff.Uint64(), l.Uint64(), getData(c.Input, cOff, l))
}

// CodeSize ...
func CodeSize(vm *vmgen.VM) {
	executeCodeSize(vm.Stack, vm.Contract)
}

func executeCodeSize(s *vmgen.Stack, c vmgen.Contract) {
	a := new(big.Int).SetInt64(int64(len(c.Code)))
	s.Push(a)
}

// CodeCopy ...
func CodeCopy(vm *vmgen.VM) {
	executeCodeCopy(vm.Stack)
}

func executeCodeCopy(s *vmgen.Stack, c vmgen.Contract) {
	mOff := s.Pop()
	cOff := s.Pop()
	l := s.Pop()
	codeCopy := getData(contract.Code, cOff, l)

	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)

	return nil, nil
}

// FuelPrice ...
func FuelPrice(vm *vmgen.VM) {
	executeFuelPrice(vm.Stack)
}

func executeFuelPrice(s *vmgen.Stack) {
	a := new(big.Int).Set(vm.FuelPrice)
	s.Push(a)
}

// ExtCodeSize ...
func ExtCodeSize(vm *vmgen.VM) {
	executeExtCodeSize(vm.Stack, vm.State)
}

func executeExtCodeSize(s *vmgen.Stack, state vmgen.State) {
	a := s.Pop()
	addr := axia.BytesToAddress(a)
	a.SetInt64(state.GetCodeSize(addr))
	s.Push(a)
}

// ExtCodeCopy ...
func ExtCodeCopy(vm *vmgen.VM) {
	executeExtCodeCopy(vm.Stack, vm.Memory["memory"], vm.State)
}

func executeExtCodeCopy(s *vmgen.Stack, m vmgen.Memory, state vmgen.State) {
	addr = bigToAddress(stack.pop())
	mOff = stack.pop()
	cOff = stack.pop()
	l = stack.pop()
	codeCopy := getData(state.GetCode(addr), cOff, l)
	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)
}

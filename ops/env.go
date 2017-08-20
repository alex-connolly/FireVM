package ops

import (
	"github.com/end-r/vmgen"
)

// Address ...
func Address(vm *vmgen.VM) {
	executeAddress(vm.Stack, vm.Input)
}

func executeAddress(s *vmgen.Stack, i vmgen.Input) {
	//s.Push(c.Address.Bytes())
}

// Balance ...
func Balance(vm *vmgen.VM) {
	executeBalance(vm.Stack, vm.State)
}

func executeBalance(s *vmgen.Stack, state vmgen.State) {
	/*a := axia.BytesToAddress(s.Pop())
	balance := state.GetBalance(a)
	s.Push(balance.Bytes())*/
}

// Origin ...
func Origin(vm *vmgen.VM) {
	//executeOrigin(vm.Stack, vm.Environment["origin"])
}

func executeOrigin(s *vmgen.Stack, a vmgen.Address) {
	s.Push(a.Bytes())
}

// Caller ...
func Caller(vm *vmgen.VM) {
	executeCaller(vm.Stack, vm.Input)
}

func executeCaller(s *vmgen.Stack, i vmgen.Input) {
	//s.Push(c.Caller.Bytes())
}

// CallValue ...
func CallValue(vm *vmgen.VM) {
	executeCallValue(vm.Stack, vm.Input)
}

func executeCallValue(s *vmgen.Stack, i vmgen.Input) {
	//s.Push(c.Value())
}

// CallDataLoad ...
func CallDataLoad(vm *vmgen.VM) {
	executeCallDataLoad(vm.Stack, vm.Input)
}

func executeCallDataLoad(s *vmgen.Stack, i vmgen.Input) {
	//cd := getData(c.Input, s.Pop(), new(big.Int).SetInt64(32))
	///s.Push(cd)
}

// CallDataSize ...
func CallDataSize(vm *vmgen.VM) {
	executeCallDataSize(vm.Stack, vm.Input)
}

func executeCallDataSize(s *vmgen.Stack, i vmgen.Input) {
	/*a := new(big.Int).SetInt64(int64(len(c.Input)))
	s.Push(a.Bytes())*/
}

// CallDataCopy ...
func CallDataCopy(vm *vmgen.VM) {
	//executeCallDataCopy(vm.Stack, vm.Memory["memory"], vm.Input)
}

func executeCallDataCopy(s *vmgen.Stack, m vmgen.Memory, i vmgen.Input) {
	/*mOff := s.Pop()
	cOff := s.Pop()
	l := s.Pop()*/
	//	m.Set(mOff.Uint64(), l.Uint64(), getData(c.Input, cOff, l))
}

// CodeSize ...
func CodeSize(vm *vmgen.VM) {
	executeCodeSize(vm.Stack, vm.Input)
}

func executeCodeSize(s *vmgen.Stack, i vmgen.Input) {
	/*a := new(big.Int).SetInt64(int64(len(c.Code)))
	s.Push(a.Bytes())*/
}

// CodeCopy ...
func CodeCopy(vm *vmgen.VM) {
	executeCodeCopy(vm.Stack, vm.Input)
}

func executeCodeCopy(s *vmgen.Stack, i vmgen.Input) {
	/*mOff := s.Pop()
	cOff := s.Pop()
	l := s.Pop()
	codeCopy := getData(contract.Code, cOff, l)

	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)*/

}

// FuelPrice ...
func FuelPrice(vm *vmgen.VM) {
	executeFuelPrice(vm.Stack)
}

func executeFuelPrice(s *vmgen.Stack) {
	/*a := new(big.Int).Set(vm.FuelPrice)
	s.Push(a)*/
}

// ExtCodeSize ...
func ExtCodeSize(vm *vmgen.VM) {
	executeExtCodeSize(vm.Stack, vm.State)
}

func executeExtCodeSize(s *vmgen.Stack, state vmgen.State) {
	/*a := s.Pop()
	addr := axia.BytesToAddress(a)
	a.SetInt64(state.GetCodeSize(addr))
	s.Push(a)*/
}

// ExtCodeCopy ...
func ExtCodeCopy(vm *vmgen.VM) {
	//executeExtCodeCopy(vm.Stack, vm.Memory["memory"], vm.State)
}

func executeExtCodeCopy(s *vmgen.Stack, m vmgen.Memory, state vmgen.State) {
	/*addr := bigToAddress(s.Pop())
	mOff := s.Pop()
	cOff := s.Pop()
	l := s.Pop()
	//codeCopy := getData(state.GetCode(addr), cOff, l)
	//memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)*/
}

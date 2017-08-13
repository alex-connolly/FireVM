package ops

import (
	"demos/ansible/go-ethereum-master/common"
	"math/big"

	"github.com/end-r/vmgen"
)

// Address ...
func Address(vm *vmgen.VM) {
	executeAddress(vm.Stack, vm.Contract)
}

func executeAddress(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(c.Address)
}

// Balance ...
func Balance(vm *vmgen.VM) {
	executeBalance(vm.Stack, vm.Contract)
}

func executeBalance(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(c.Address)
}

// Origin ...
func Origin(vm *vmgen.VM) {
	executeOrigin(vm.Stack, vm.Environment)
}

func executeOrigin(s *vmgen.Stack, e vmgen.Environment) {
	s.Push(e["origin"])
}

// Caller ...
func Caller(vm *vmgen.VM) {
	executeCaller(vm.Stack, vm.Contract)
}

func executeCaller(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(createBigInt(c.Caller().Bytes()))
}

// CallValue ...
func CallValue(vm *vmgen.VM) {
	executeCallValue(vm.Stack)
}

func executeCallValue(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(createBigInt(c.value))
}

// CallDataLoad ...
func CallDataLoad(vm *vmgen.VM) {
	executeCallDataLoad(vm.Stack)
}

func executeCallDataLoad(s *vmgen.Stack) {
	stack.push(createBigInt(getData(contract.Input, s.Pop(1), common.Big32)))
}

// CallDataSize ...
func CallDataSize(vm *vmgen.VM) {
	executeCallDataSize(vm.Stack)
}

func executeCallDataSize(s *vmgen.Stack, c *vmgen.Contract) {
	s.Push(new(big.Int).SetInt64(int64(len(c.Input))))
}

// CallDataCopy ...
func CallDataCopy(vm *vmgen.VM) {
	executeCallDataCopy(vm.Stack)
}

func executeCallDataCopy(s *vmgen.Stack) {
}

// CodeSize ...
func CodeSize(vm *vmgen.VM) {
	executeCodeSize(vm.Stack)
}

func executeCodeSize(s *vmgen.Stack) {
}

// CodeCopy ...
func CodeCopy(vm *vmgen.VM) {
	executeCodeCopy(vm.Stack)
}

func executeCodeCopy(s *vmgen.Stack) {
}

// FuelPrice ...
func FuelPrice(vm *vmgen.VM) {
	executeFuelPrice(vm.Stack)
}

func executeFuelPrice(s *vmgen.Stack) {
}

// ExtCodeSize ...
func ExtCodeSize(vm *vmgen.VM) {
	executeExtCodeSize(vm.Stack)
}

func executeExtCodeSize(s *vmgen.Stack) {
}

// ExtCodeCopy ...
func ExtCodeCopy(vm *vmgen.VM) {
	executeExtCodeCopy(vm.Stack)
}

func executeExtCodeCopy(s *vmgen.Stack) {
}

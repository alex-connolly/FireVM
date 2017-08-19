package ops

import (
	"math/big"

	"github.com/end-r/vmgen"
)

// And ...
func And(vm *vmgen.VM) {
	executeAnd(vm.Stack)
}

func executeAnd(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).And(a, b)
	s.Push(c.Bytes())
}

// Or ...
func Or(vm *vmgen.VM) {
	executeOr(vm.Stack)
}

func executeOr(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Or(a, b)
	s.Push(c.Bytes())
}

// Xor ..
func Xor(vm *vmgen.VM) {
	executeXor(vm.Stack)
}

func executeXor(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Xor(a, b)
	s.Push(c.Bytes())
}

// Not ...
func Not(vm *vmgen.VM) {
	executeNot(vm.Stack)
}

func executeNot(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := new(big.Int).Not(a)
	s.Push(b.Bytes())
}

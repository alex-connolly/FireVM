package ops

import (
	"math/big"

	"github.com/end-r/vmgen"
)

var (
	bigZero = new(big.Int)
)

// Add ...
func Add(vm *vmgen.VM) {
	executeAddition(vm.Stack)
}

func executeAddition(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Add(a, b)
	s.Push(c.Bytes())
}

// Sub ...
func Sub(vm *vmgen.VM) {
	executeAddition(vm.Stack)
}

func executeSubtraction(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Sub(a, b)
	s.Push(c.Bytes())
}

// Mul ...
func Mul(vm *vmgen.VM) {
	executeMultiplication(vm.Stack)
}

func executeMultiplication(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Mul(a, b)
	s.Push(c.Bytes())
}

// Div ...
func Div(vm *vmgen.VM) {
	executeDivision(vm.Stack)
}

func executeDivision(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Div(a, b)
	s.Push(c.Bytes())
}

// Mod ...
func Mod(vm *vmgen.VM) {
	executeModulo(vm.Stack)
}

func executeModulo(s *vmgen.Stack) {
	a := bigInt(s.Pop())
	b := bigInt(s.Pop())
	c := new(big.Int).Mod(a, b)
	s.Push(c.Bytes())
}

// AddMod ...
func AddMod(vm *vmgen.VM) {
	executeAddMod(vm.Stack)
}

func executeAddMod(s *vmgen.Stack) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	z := bigInt(s.Pop())
	if z.Cmp(bigZero) > 0 {
		a := x.Add(x, y)
		a.Mod(a, z)
		s.Push(a.Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

// MulMod ...
func MulMod(vm *vmgen.VM) {
	executeMulMod(vm.Stack)
}

func executeMulMod(s *vmgen.Stack) {
	x := bigInt(s.Pop())
	y := bigInt(s.Pop())
	z := bigInt(s.Pop())
	if z.Cmp(bigZero) > 0 {
		m := x.Add(x, y)
		m.Mod(m, z)
		s.Push(m.Bytes())
	} else {
		s.Push(new(big.Int).Bytes())
	}
}

// Concat ...
func Concat(vm *vmgen.VM) {
	executeConcat(vm.Stack)
}

func executeConcat(s *vmgen.Stack) {
	a := s.Pop()
	b := s.Pop()
	c := append(a, b...)
	s.Push(c)
}

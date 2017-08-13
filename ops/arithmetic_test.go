package ops

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/vmgen"
)

func TestOpAdd(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeAddition(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop(1))
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1+o2))) == 0, "wrong add value")
}

func TestOpSub(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeSubtraction(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop(1))
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1-o2))) == 0, "wrong sub value")
}

func TestOpMul(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeMultiplication(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop(1))
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1*o2))) == 0, "wrong mul value")
}

func TestOpDiv(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeDivision(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop(1))
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1/o2))) == 0, fmt.Sprintf("wrong div value: %s", c.String()))
}

func TestOpMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeModulo(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop(1))
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1%o2))) == 0, "wrong mod value")
}

func TestOpAddMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeAddMod(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	//c := new(big.Int).SetBytes(s.Pop(1))

}

func TestOpMulMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeMulMod(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	//c := new(big.Int).SetBytes(s.Pop(1))
}

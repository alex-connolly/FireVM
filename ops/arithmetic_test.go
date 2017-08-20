package ops

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/vmgen"
)

func TestAdd(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeAddition(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1+o2))) == 0, "wrong add value")
}

func TestSub(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeSubtraction(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1-o2))) == 0, "wrong sub value")
}

func TestMul(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeMultiplication(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1*o2))) == 0, "wrong mul value")
}

func TestDiv(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeDivision(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1/o2))) == 0, fmt.Sprintf("wrong div value: %s", c.String()))
}

func TestMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeModulo(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetInt64(int64(o1%o2))) == 0, "wrong mod value")
}

func TestAddMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2, o3 := 4, 2, 1
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o3)).Bytes())
	executeAddMod(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	//c := new(big.Int).SetBytes(s.Pop())

}

func TestMulMod(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2, o3 := 4, 2, 1
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o3)).Bytes())
	executeMulMod(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	//c := new(big.Int).SetBytes(s.Pop())
}

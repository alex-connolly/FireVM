package ops

import (
	"math/big"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/vmgen"
)

func TestLt(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeLt(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "wrong lt value")
}

func TestGt(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 2, 4
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeGt(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "wrong gt value")
}

func TestSLt(t *testing.T) {

}

func TestSGt(t *testing.T) {

}

func TestEq(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 4
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeEq(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "wrong eq value")
}

func TestIsZero(t *testing.T) {
	s := new(vmgen.Stack)
	o1 := 0
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	executeIsZero(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	c := new(big.Int).SetBytes(s.Pop())
	goutil.Assert(t, c.Cmp(new(big.Int).SetUint64(1)) == 0, "wrong isZero value")
}

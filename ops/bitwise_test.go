package ops

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/end-r/vmgen"

	"github.com/end-r/goutil"
)

func TestOpAnd(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeAnd(s)
	goutil.Assert(t, s.Size() == 1, fmt.Sprintf("wrong stack size: %d", s.Size()))
}

func TestOpOr(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeOr(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
}

func TestOpXor(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	executeXor(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
}

func TestOpNot(t *testing.T) {
	s := new(vmgen.Stack)
	o1 := 0
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	executeNot(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
}

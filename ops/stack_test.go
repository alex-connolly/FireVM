package ops

import (
	"math/big"
	"testing"

	"github.com/end-r/vmgen"

	"github.com/end-r/goutil"
)

func TestPop(t *testing.T) {
	s := new(vmgen.Stack)
	o1, o2 := 4, 2
	s.Push(new(big.Int).SetInt64(int64(o1)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(o2)).Bytes())
	goutil.Assert(t, s.Size() == 2, "wrong stack size")
	executePop(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	executePop(s)
	goutil.Assert(t, s.Size() == 0, "wrong stack size")
}

func TestPush(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size")
	//c := vmgen.NewInput("01AA")
	//executePush(s, c)
	//goutil.Assert(t, s.Size() == 1, "wrong stack size")
}

func TestDup(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size")
	//	c := vmgen.NewInput("01AA")
	//	executePush(s, c)
	//	goutil.Assert(t, s.Size() == 1, "wrong stack size")
}

func TestSwap(t *testing.T) {

}

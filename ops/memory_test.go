package ops

import (
	"axia/fireVM/memory"
	"math/big"
	"testing"

	"github.com/end-r/goutil"
	"github.com/end-r/vmgen"
)

func TestMemSize(t *testing.T) {
	s := new(vmgen.Stack)
	size := 64
	m := memory.NewFireMemory(size)
	executeMemSize(s, m)
	goutil.Assert(t, s.Size() == 1, "wrong stack size")
	msize := int(new(big.Int).SetBytes(s.Pop()).Int64())
	goutil.Assert(t, size == msize, "wrong msize")
}

func TestLoad(t *testing.T) {
	s := new(vmgen.Stack)
}

func TestStore(t *testing.T) {
	s := new(vmgen.Stack)
}

func TestSet(t *testing.T) {
	s := new(vmgen.Stack)

}

func TestGet(t *testing.T) {
	s := new(vmgen.Stack)

}

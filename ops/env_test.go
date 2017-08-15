package ops

import (
	"axia/fireVM/memory"
	"axia/vmgen"
	"math/big"
	"testing"

	"github.com/end-r/goutil"
)

func TestAddress(t *testing.T) {
	s := new(vmgen.Stack)
	s.Push()
}

func TestBalance(t *testing.T) {
	s := new(vmgen.Stack)
	a := axia.ParseAddress("x0000000000000000001")
	s.Push(a.Bytes())

}

func TestOrigin(t *testing.T) {

}

func TestCaller(t *testing.T) {

}

func TestCallValue(t *testing.T) {

}

func TestCallDataLoad(t *testing.T) {

}

func TestCallDataSize(t *testing.T) {

}

func TestCallDataCopy(t *testing.T) {

}

func TestCodeSize(t *testing.T) {
	s := new(vmgen.Stack)
	c := contractWithCode()
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeCodeSize(s, c)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestCodeCopy(t *testing.T) {
	s := new(vmgen.Stack)
	c := vmgen.ContractWithCode()
	m := new(memory.FireMemory)
	mOff, cOff, l := 4, 2, 1
	s.Push(new(big.Int).SetInt64(int64(mOff)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(cOff)).Bytes())
	s.Push(new(big.Int).SetInt64(int64(l)).Bytes())
	executeCodeCopy(s, c, m)
}

func TestFuelPrice(t *testing.T) {
	s := new(vmgen.Stack)
	price := new(big.Int).SetInt64(int64(1000))
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeFuelPrice(s, price)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestExtCodeSize(t *testing.T) {
	s := new(vmgen.Stack)
	a := axia.NewAddress()
	state := new(vmgen.State)
	c := contractWithCode("010101")
	state.Store(a, c)
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	s.Push(a.Bytes())
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
	executeExtCodeSize(s, state)
}

func TestExtCodeCopy(t *testing.T) {
	s := new(vmgen.Stack)
	a := axia.NewAddress()
	s.Push(a.Bytes())
	mOff = stack.pop()
	cOff = stack.pop()
	l = stack.pop()
	codeCopy := getData(state.GetCode(addr), cOff, l)
	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)
	executeExtCodeCopy(s, state, m)
}

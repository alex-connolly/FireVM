package ops

/*
import (
	"math/big"
	"testing"

	"github.com/end-r/vmgen"

	"github.com/end-r/goutil"
)

func TestAddress(t *testing.T) {
	s := new(vmgen.Stack)
	c := createContract()
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeAddress(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestBalance(t *testing.T) {
	s := new(vmgen.Stack)
	a := axia.ParseAddress("x0000000000000000001")
	state := new(vmgen.State)
	balance := 100
	state.SetBalance(a, new(big.Int).SetInt64(int64(balance)))
	s.Push(a.Bytes())
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
	executeBalance(s, state)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
	goutil.Assert(t, s.Pop().Cmp(new(big.Int).SetInt64(int64(balance))), "wrong balance")
}

func TestOrigin(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	a := axia.NewAddress("x0000000000000000001")
	executeOrigin(s, a)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestCaller(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeCaller(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestCallValue(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeCallValue(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestCallDataLoad(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeCallDataLoad(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestCallDataSize(t *testing.T) {
	s := new(vmgen.Stack)
	goutil.Assert(t, s.Size() == 0, "wrong stack size (not 0)")
	executeCallDataSize(s)
	goutil.Assert(t, s.Size() == 1, "wrong stack size (not 1)")
}

func TestCallDataCopy(t *testing.T) {
	s := new(vmgen.Stack)
	executeCallDataCopy(s)
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
	state := new(vmgen.State)
	c := contractWithCode("010101")
	a := axia.NewAddress()
	state.Store(a, c)
	s.Push(a.Bytes())
	mOff = s.Pop()
	cOff = s.Pop()
	l = s.Pop()
	codeCopy := getData(state.GetCode(addr), cOff, l)
	memory.Set(mOff.Uint64(), l.Uint64(), codeCopy)
	executeExtCodeCopy(s, state, m)
}

*/

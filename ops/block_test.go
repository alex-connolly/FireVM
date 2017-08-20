package ops

import (
	"testing"

	"github.com/end-r/vmgen"

	"github.com/end-r/goutil"
)

func TestBlockhash(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["blockhash"] = []byte("random")
	pushBytes(s, e["blockhash"])
}

func TestCoinbase(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["coinbase"] = []byte("random")
	pushBytes(s, e["coinbase"])
}

func TestTimestamp(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["timestamp"] = []byte("random")
	pushBytes(s, e["timestamp"])
}

func TestNumber(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["blockNumber"] = []byte("random")
	pushBytes(s, e["blockNumber"])
}

func TestDifficulty(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["difficulty"] = []byte("random")
	pushBytes(s, e["difficulty"])
}

func TestFuelLimit(t *testing.T) {

}

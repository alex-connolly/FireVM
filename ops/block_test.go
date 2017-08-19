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
	executeBlockhash(s, e)
}

func TestCoinbase(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["coinbase"] = []byte("random")
	executeCoinbase(s, e)
}

func TestTimestamp(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["coinbase"] = []byte("random")
	executeCoinbase(s, e)
}

func TestNumber(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["blockNumber"] = []byte("random")
	executeNumber(s, e)
}

func TestDifficulty(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["difficulty"] = []byte("random")
	executeDifficulty(s, e)
}

func TestFuelLimit(t *testing.T) {
	s := new(vmgen.Stack)
	e := make(vmgen.Environment)
	goutil.Assert(t, s.Size() == 0, "wrong starting size")
	e["fuelLimit"] = []byte("random")
	executeFuelLimit(s, e)
}

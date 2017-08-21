package memory

import (
	"axia/vmgen"
	"testing"
)

func TestMemory(t *testing.T) {

}

func TestNewFireMemory(t *testing.T) {

}

func TestAssertion(t *testing.T) {
	m := NewFireMemory(10)
	validMemory(m)
}

func validMemory(m vmgen.Memory) {

}

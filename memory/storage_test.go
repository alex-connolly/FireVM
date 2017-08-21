package memory

import (
	"fmt"
	"testing"

	"github.com/end-r/go-axia/axia"

	"github.com/end-r/goutil"
)

func TestStorageSetGet(t *testing.T) {
	s := NewFireStorage()
	val := axia.HashString("me")
	s.Set(axia.HashString("hello"), val)
	goutil.Assert(t, s.Get(axia.HashString("hello")) == val, "")

}

func TestStorageCopy(t *testing.T) {
	s := NewFireStorage()
	s.Set(axia.HashString("hello"), axia.HashString("me"))
	goutil.Assert(t, s.Size() == 1, "wrong original size")
	f := s.Copy()
	goutil.Assert(t, s.Size() == 1, "wrong original size post-copy")
	goutil.Assert(t, f.Size() == 1, "wrong copy size")
}

func TestStorageSize(t *testing.T) {
	s := NewFireStorage()
	goutil.Assert(t, s.Size() == 0, "wrong 0 size")
	s.Set(axia.HashString("hello"), axia.HashString("me"))
	goutil.Assert(t, s.Size() == 1, "wrong 1 size")
	s.Set(axia.HashString("hello"), axia.HashString("me"))
	goutil.Assert(t, s.Size() == 1, "wrong duplicate size")
	s.Set(axia.HashString("goodbye"), axia.HashString("friend"))
	goutil.Assert(t, s.Size() == 2, fmt.Sprintf("wrong 2 size: %d", s.Size()))
}

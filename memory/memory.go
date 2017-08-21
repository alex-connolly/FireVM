package memory

// FireMemory ...
type FireMemory struct {
	capacity    int
	store       []byte
	lastGasCost uint64
	lastReturn  []byte
}

// NewFireMemory ...
func NewFireMemory(capacity int) *FireMemory {
	m := new(FireMemory)
	m.capacity = capacity
	return m
}

// Size ...
func (m *FireMemory) Size() int {
	return m.capacity
}

// Set sets offset + size to value
func (m *FireMemory) Set(offset, size uint64, value []byte) {
	// length of store may never be less than offset + size.
	// The store should be resized PRIOR to setting the memory
	if size > uint64(len(m.store)) {
		panic("INVALID memory: store empty")
	}

	// It's possible the offset is greater than 0 and size equals 0. This is because
	// the calcMemSize (common.go) could potentially return 0 when size is zero (NO-OP)
	if size > 0 {
		copy(m.store[offset:offset+size], value)
	}
}

// Get returns offset + size as a new slice
func (m *FireMemory) Get(offset, size int64) (cpy []byte) {
	if size == 0 {
		return nil
	}
	if len(m.store) > int(offset) {
		cpy = make([]byte, size)
		copy(cpy, m.store[offset:offset+size])

		return
	}
	return
}

func (s *FireMemory) DisplayContents() {

}

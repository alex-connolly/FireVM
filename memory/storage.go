package memory

// FireStorage ...
type FireStorage map[firevm.Hash]firevm.Hash

func (s FireStorage) Set(key, value []byte) {
	s.data[key] = value
}

func (s FireStorage) Get(key []byte) []byte {
	return s.data[key]
}

func (s FireStorage) Size() int {
	return len(s.data)
}

func (s FireStorage) Copy() (f FireStorage) {
	for k, v := range s {
		f[k] = v
	}
	return f
}

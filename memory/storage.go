package memory

// FireStorage ...
type FireStorage struct {
	data map[[]byte][]byte
}

func (s FireStorage) Set(key, value []byte) {
	s.data[key] = value
}

func (s FireStorage) Get(key []byte) []byte {
	return s.data[key]
}

func (s FireStorage) Size() int {
	return len(s.data)
}

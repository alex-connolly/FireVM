package memory

import "github.com/end-r/go-axia/axia"

// FireStorage ...
type FireStorage struct {
	data map[axia.Hash]axia.Hash
}

// NewFireStorage ...
func NewFireStorage() *FireStorage {
	s := new(FireStorage)
	s.data = make(map[axia.Hash]axia.Hash)
	return s
}

// Set ...
func (s FireStorage) Set(key, value axia.Hash) {
	s.data[key] = value
}

// Get ...
func (s FireStorage) Get(key axia.Hash) axia.Hash {
	return s.data[key]
}

// Size ...
func (s FireStorage) Size() int {
	return len(s.data)
}

// Copy ...
func (s FireStorage) Copy() (f FireStorage) {
	f.data = make(map[axia.Hash]axia.Hash)
	for k, v := range s.data {
		f.data[k] = v
	}
	return f
}

func (s FireStorage) DisplayContents() {

}

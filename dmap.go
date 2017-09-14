package firevm

// DMap = Deterministic Map
type DMap struct {
	data  map[interface{}]interface{}
	array []interface{}
}

// Size ...
func (d *DMap) Size() int {
	return len(d.array)
}

// Assign ...
func (d *DMap) Assign(key, value interface{}) {
	d.data[key] = value
}

// Retrieve ...
func (d *DMap) Retrieve(key interface{}) interface{} {
	return d.data[key]
}

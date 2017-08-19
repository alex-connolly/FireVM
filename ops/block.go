package ops

import "github.com/end-r/vmgen"

func pushBytes(s *vmgen.Stack, bytes []byte) {
	s.Push(bytes)
}

// Blockhash ...
func Blockhash(vm *vmgen.VM) {
	pushBytes(vm.Stack, vm.Environment["blockhash"])
}

// Coinbase ...
func Coinbase(vm *vmgen.VM) {
	pushBytes(vm.Stack, vm.Environment["coinbase"])
}

// Timestamp ...
func Timestamp(vm *vmgen.VM) {
	pushBytes(vm.Stack, vm.Environment["timestamp"])
}

// Number ...
func Number(vm *vmgen.VM) {
	pushBytes(vm.Stack, vm.Environment["number"])
}

// Difficulty ...
func Difficulty(vm *vmgen.VM) {
	pushBytes(vm.Stack, vm.Environment["difficulty"])
}

// FuelLimit ...
func FuelLimit(vm *vmgen.VM) {
	pushBytes(vm.Stack, vm.Environment["fuelLimit"])
}

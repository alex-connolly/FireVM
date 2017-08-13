package ops

import "github.com/end-r/vmgen"

// Blockhash ...
func Blockhash(vm *vmgen.VM) {
	executeBlockhash(vm.Stack, vm.Environment)
}

func executeBlockhash(s *vmgen.Stack, e vmgen.Environment) {

}

// Coinbase ...
func Coinbase(vm *vmgen.VM) {
	executeCoinbase(vm.Stack, vm.Environment)
}

func executeCoinbase(s *vmgen.Stack, e vmgen.Environment) {
	s.Push(e["coinbase"])
}

// Timestamp ...
func Timestamp(vm *vmgen.VM) {
	executeTimestamp(vm.Stack, vm.Environment)
}

func executeTimestamp(s *vmgen.Stack, e vmgen.Environment) {
	s.Push(e["timestamp"])
}

// Number ...
func Number(vm *vmgen.VM) {
	executeNumber(vm.Stack, vm.Environment)
}

func executeNumber(s *vmgen.Stack, e vmgen.Environment) {
	s.Push(e["blockNumber"])
}

// Difficulty ...
func Difficulty(vm *vmgen.VM) {
	executeDifficulty(vm.Stack, vm.Environment)
}

func executeDifficulty(s *vmgen.Stack, e vmgen.Environment) {
	s.Push(e["difficulty"])
}

// FuelLimit ...
func FuelLimit(vm *vmgen.VM) {
	executeFuelLimit(vm.Stack, vm.Environment)
}

func executeFuelLimit(s *vmgen.Stack, e vmgen.Environment) {
	s.Push(e["fuelLimit"])
}

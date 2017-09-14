package firevm

import (
	"log"

	"github.com/end-r/firevm/ops"
	"github.com/end-r/vmgen"
)

const vm = `
name = "FireVM"
author = "ender"
description = "A stack-based virtual machine for executing smart contracts."

category("ARITHMETIC"){

    description = "Basic arithmetic operations."

    instruction("ADD", "01"){
        description = "Addition"
        validate = 2
        fuel = 10
    }

    instruction("SUB", "02"){
        description = "Subtraction"
        validate = 2
        fuel = 10
    }

    instruction("MUL", "03"){
        description = "Multiplication"
        validate = 2
        fuel = 10
    }

    instruction("DIV", "04"){
        description = "Multiplication"
        validate = 2
        fuel = 10
    }

    instruction("MOD", "05"){
        validate = 2
    }

    instruction("ADDMOD", "06"){
        validate = 3
    }

    instruction("MULMOD", "07"){
        validate = 3
    }

    instruction("CONCAT", "0A"){
        description = "Concatenates two byte arrays"
        validate = 2
    }
}

category("Crypto"){

    description = "Cryptographic functions."

    instruction("SHA3", "30"){

    }
}

category("Comparison"){

    description = "Comparison operations."

    instruction("LT"){

    }

    instruction("GT"){

    }

    instruction("SLT"){

    }

    instruction("SGT"){

    }

    instruction("EQ"){

    }

    instruction("ISZERO"){

    }
}

category("BITWISE"){

    instruction("AND"){

    }

    instruction("OR"){

    }

    instruction("XOR"){

    }

    instruction("NOT"){

    }

}

category("ENV"){

    instruction("ADDRESS"){

    }

    instruction("BALANCE"){

    }

    instruction("ORIGIN"){

    }

    instruction("CALLER"){

    }

    instruction("CALLVALUE"){

    }

    instruction("CALLDATALOAD"){

    }

    instruction("CALLDATASIZE"){

    }

    instruction("CALLDATACOPY"){

    }

    instruction("CODESIZE"){

    }

    instruction("CODECOPY"){

    }

    instruction("FUELCOST"){

    }

    instruction("EXTCODESIZE"){

    }

    instruction("EXTCODECOPY"){

    }

}

category("BLOCK"){

    description = "Operations which relate to the current block."

    instruction("BLOCKHASH"){

    }

    instruction("COINBASE"){

    }

    instruction("TIMESTAMP"){

    }

    instruction("NUMBER"){

    }

    instruction("DIFFICULTY"){

    }

    instruction("FUELLIMIT"){

    }

}

category("FLOW"){

    instruction("JUMP", "70"){

    }

    instruction("JUMPI", "71"){

    }

    instruction("PC", "72"){

    }

    instruction("FUEL", "73"){

    }

    instruction("JUMPDEST", "74"){

    }

    instruction("STOP", "75"){

    }

    instruction("RETURN", "76"){

    }
}

category("MEMORY"){
    instruction("GET", "40"){

    }

    instruction("SET", "41"){

    }

    instruction("STORE", "42"){

    }

    instruction("LOAD", "43"){

    }
    instruction("MSIZE", "44"){

    }
}

category("Stack"){
    instruction("POP", "60"){

    }
    instruction("PUSH", "61"){

    }
    instruction("DUP", "62"){

    }
    instruction("SWAP", "63"){

    }
}

category("Structures"){

    description = "Creating structures."

    instruction("ARRAY", "A0"){

    }

    instruction("MAP", "A1"){

    }
}
`

var (
	fuels    = map[string]vmgen.FuelFunction{}
	executes = map[string]vmgen.ExecuteFunction{
		// arithmetic operations
		"ADD":    ops.Add,
		"SUB":    ops.Sub,
		"DIV":    ops.Div,
		"MUL":    ops.Mul,
		"MOD":    ops.Mod,
		"ADDMOD": ops.AddMod,
		"MULMOD": ops.MulMod,
		"CONCAT": ops.Concat,
		// comparison operations
		"LT":     ops.Lt,
		"GT":     ops.Gt,
		"SLT":    ops.SLt,
		"SGT":    ops.SGt,
		"EQ":     ops.Eq,
		"ISZERO": ops.IsZero,
		// bitwise operations
		"AND": ops.And,
		"OR":  ops.Or,
		"NOT": ops.Not,
		"XOR": ops.Xor,
		// block operations
		"HASH":       ops.Blockhash,
		"COINBASE":   ops.Coinbase,
		"TIMESTAMP":  ops.Timestamp,
		"NUMBER":     ops.Number,
		"DIFFICULTY": ops.Difficulty,
		"FUELLIMIT":  ops.FuelLimit,
		// environment operations
		"ADDRESS":      ops.Address,
		"ORIGIN":       ops.Origin,
		"BALANCE":      ops.Balance,
		"CALLER":       ops.Caller,
		"CALLVALUE":    ops.CallValue,
		"CALLDATALOAD": ops.CallDataLoad,
		"CALLDATASIZE": ops.CallDataSize,
		"CALLDATACOPY": ops.CallDataCopy,
		"CODESIZE":     ops.CodeSize,
		"CODECOPY":     ops.CodeCopy,
		"FUELPRICE":    ops.FuelPrice,
		"EXTCODESIZE":  ops.ExtCodeSize,
		"EXTCODECOPY":  ops.ExtCodeCopy,
		// crypto operations
		"SHA3": ops.SHA3,
		//stack operations
		"POP.": ops.Pop,
		"PUSH": ops.Push,
		"DUP":  ops.Dup,
		"SWAP": ops.Swap,
		// external operations
		"LOG": ops.Log,
		// flow operations
		"JUMP":     ops.Jump,
		"JUMPI":    ops.JumpI,
		"PC":       ops.PC,
		"FUEL":     ops.Fuel,
		"MEMSIZE":  ops.MemSize,
		"JUMPDEST": ops.JumpDest,
		// memory operations
		"SET":   ops.Set,
		"GET":   ops.Get,
		"LOAD":  ops.Load,
		"STORE": ops.Store,
	}
)

// NewVM returns a new FireVM instance
func NewVM() *vmgen.VM {
	vm, errs := vmgen.CreateVM(vm, nil, executes, fuels, nil)
	if errs != nil {
		log.Println(errs)
		return nil
	}
	return vm
}

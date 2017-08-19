package ops

import "math/big"

func bigInt(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}

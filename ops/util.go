package ops

import "math/big"

func createBigInt(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}

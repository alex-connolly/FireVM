package ops

import "math/big"

func bigInt(bytes []byte) *big.Int {
	return new(big.Int).SetBytes(bytes)
}

func intAsBig(val int) *big.Int {
	return new(big.Int).SetInt64(int64(val))
}

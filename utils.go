package main

import (
	"fmt"
	"math/big"
)

func GetEthAmount(amount *big.Int) string {
	value := new(big.Int)
	value = value.Div(amount, new(big.Int).SetUint64(1000000000000000))
	intval := value.Uint64()
	floatval := float32(intval) / 1000.0
	return fmt.Sprintf("%.3f", floatval)
}

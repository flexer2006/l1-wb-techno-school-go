package twentytwo

import "math/big"

func parseBig(s string) *big.Int {
	v, _ := new(big.Int).SetString(s, 10)
	return v
}

func parseOperands(aStr, bStr string) (*big.Int, *big.Int) {
	return parseBig(aStr), parseBig(bStr)
}

func BigAdd(aStr, bStr string) *big.Int {
	a, b := parseOperands(aStr, bStr)
	return new(big.Int).Add(a, b)
}

func BigSub(aStr, bStr string) *big.Int {
	a, b := parseOperands(aStr, bStr)
	return new(big.Int).Sub(a, b)
}

func BigMul(aStr, bStr string) *big.Int {
	a, b := parseOperands(aStr, bStr)
	return new(big.Int).Mul(a, b)
}

func BigDiv(aStr, bStr string) *big.Int {
	a, b := parseOperands(aStr, bStr)
	return new(big.Int).Quo(a, b)
}

func BigQuoRem(aStr, bStr string) (*big.Int, *big.Int) {
	a, b := parseOperands(aStr, bStr)
	quot := new(big.Int)
	rem := new(big.Int)
	quot.QuoRem(a, b, rem)
	return quot, rem
}

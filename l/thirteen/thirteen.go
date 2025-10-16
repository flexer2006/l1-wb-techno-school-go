package thirteen

func swapAS(a, b int) {
	a += b
	b = a - b
	a -= b
}

func swapBitwise(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
}

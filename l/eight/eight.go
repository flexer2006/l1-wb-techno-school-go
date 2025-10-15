package eight

func setBit(num, idx, bit int) int {
	return (num &^ (1 << idx)) | (bit << idx)
}

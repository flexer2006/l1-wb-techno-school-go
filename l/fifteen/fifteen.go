package fifteen

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	b := make([]byte, 100)
	copy(b, v)
	justString = string(b)
}

func createHugeString(s int) string {
	return string(make([]byte, s))
}

func main() {
	someFunc()
}

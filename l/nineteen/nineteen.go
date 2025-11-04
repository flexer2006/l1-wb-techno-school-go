package nineteen

import "fmt"

func reverseString(s string) string {
	res := []rune(s)
	for l, r := 0, len(res)-1; l < r; {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

func nineTask() {
	fmt.Println(reverseString("главрыба"))
	fmt.Println(reverseString("qwerty)йцукен"))
	fmt.Println(reverseString("1234567890"))
}

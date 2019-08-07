package main

import (
	"bytes"
	"fmt"
)

/*
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

*/

func intToRoman(num int) string {
	// b := make([]int, 13)
	var base [13]int = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var str [13]string = [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var div int
	var buf bytes.Buffer
	for i := 0; i < 13; i++ {
		div = num / base[i]
		if div > 0 {
			// fmt.Println(div)
			num = num % base[i]
			for div > 0 {
				buf.WriteString(str[i])
				div--
			}
		}
	}
	return buf.String()
}

func main() {

	num := 1994
	fmt.Println(intToRoman(num))
	return
}

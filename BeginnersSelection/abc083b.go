package BeginnersSelection

import (
	"fmt"
)

func Abc083b() {
	var n, a, b int
	sum := 0

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	for i := 1; i <= n; i++ {
		tmp := 0
		for j := i; j > 0; j /= 10 {
			tmp += j % 10
		}
		if tmp >= a && tmp <= b {
			sum += i
		}
	}

	fmt.Println(sum)

}

package BeginnersSelection

import (
	"fmt"
)

func PracticeA() {
	var a, b, c int
	var s string

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%s", &s)

	fmt.Printf("%d %s\n", a+b+c, s)
}
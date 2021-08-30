package BeginnersSelection

import "fmt"

func Abc086a() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	if a*b%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

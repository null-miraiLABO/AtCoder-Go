package BeginnersSelection

import "fmt"

func Abc087b() {
	var a, b, c, x int
	cn := 0
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				sum := 500*i + 100*j + 50*k
				if x == sum {
					cn++
				}

			}
		}
	}

	fmt.Println(cn)

}

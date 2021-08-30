package BeginnersSelection

import (
	"fmt"
)

func Abc081b() {
	resurt := 0
	var le, tmp int

	fmt.Scanf("%d", &le)

	for i := 0; i < le; i++ {
		fmt.Scanf("%d", &tmp)

		cnt := 0
		for tmp > 0 && tmp%2 == 0 {
			tmp /= 2
			cnt++
		}

		if i == 0 || cnt < resurt {
			resurt = cnt
		}

	}

	fmt.Println(resurt)

}

package practice

import (
	"fmt"
)

func Anotoverflow() {
	var n int
	num := 2
	flag_higth := 30
	flag_low := 10

	fmt.Scanf("%d", &n)

	for i:=0;i<=n;i++ {
		num = num * 2
	}

	fmt.Printf("%d\n", num)

	if n < flag_higth && n > flag_low{
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}
package BeginnersSelection

import (
	"fmt"
	"strings"
)

func Abc081a() {
	var str string
	count := 0
	fmt.Scanf("%s", &str)

	arr := strings.Split(str, "")
	fmt.Println(arr)

	for i := 0; i < 3; i++ {
		if arr[i] == "1" {
			count++
		}
	}

	fmt.Println(count)

}

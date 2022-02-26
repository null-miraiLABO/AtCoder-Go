package BeginnersSelection

import "fmt"

func Abc085c(){
    var sum_cn, sum_value int
	var ten_thousand, five_thousand, thousand int

    fmt.Scanf("%d", &sum_cn)
    fmt.Scanf("%d", &sum_value)

	ten_thousand, five_thousand, thousand = count_money(sum_cn, sum_value)

	fmt.Printf("%d %d %d \n", ten_thousand, five_thousand, thousand)
}

func count_money(sum_cn,sum_value int) (int,int,int) {
	var x, y, z int
	var i, j, k int

	for i=0;i<=sum_cn;i++ {
		for j=0;j<=sum_cn;j++ {
			k = sum_cn-(i+j)
			if k >= 0 {
				if 10000*i + 5000*j + 1000*k == sum_value{
					x = i
					y = j
					z = k
					return x, y, z
				}
			}
		}
	}

	x = -1
	y = -1
	z = -1
	return x, y, z
}
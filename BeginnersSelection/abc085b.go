package BeginnersSelection

import (
    "fmt"
    "sort"
)

func Abc085b(){
    var n, scan_tmp int
	count := 0

    fmt.Scanf("%d", &n)

    d := make([]int, 0)

    for i:=0;i<n;i++{
        fmt.Scanf("%d", &scan_tmp)
        d = append(d, scan_tmp)
    }

	// 昇順ソート
    sort.Ints(d)

    for i:=0;i<n-1;i++{
		if d[i] != d[i+1] { // 同じ直径はカウントしない
			count++
		}
    }

	count++ // スライスの最後の要素分

    fmt.Println(count)

}
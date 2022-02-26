package BeginnersSelection

import "fmt"

func Abc049c(){
	var s, ans string

	fmt.Scanf("%s", &s)

	ans = judge(s)

	fmt.Println(ans)
}

func judge(s string) string {
	ans := "NO"

	for len(s) < 5 {
		if s[:5] == "dream" {
			s == s[5:]
		}else if s[:5] == "dreamer" {
			s == s[5:]
		}else{
			break
		}

		if len(s) == 0 {
			ans = "YES"
		}
	}

	return ans
}
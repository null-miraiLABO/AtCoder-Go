package main

import (
	"atCoder/BeginnersSelection"
	"fmt"
)

func main() {
	var sel string
	fmt.Println("問題を選択して下さい。")
	fmt.Scanf("%s", &sel)

	switch sel {
	case "practiceA":
		fmt.Println("this is practiceA")
		BeginnersSelection.PracticeA()
	case "abc086a":
		fmt.Println("this is abc086a")
		BeginnersSelection.Abc086a()
	case "abc081a":
		fmt.Println("this is abc086a")
		BeginnersSelection.Abc081b()
	}

}

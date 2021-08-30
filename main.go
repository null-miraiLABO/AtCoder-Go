package main

import (
	"atCoder/BeginnersSelection"
	"fmt"
	"os/exec"
)

func main() {
	var sel string
	cmd := exec.Command("ls", "BeginnersSelection/")
	var result, err = cmd.Output()

	fmt.Printf(">ファイル一覧\n%s\n", result)
	fmt.Println("err : ", err)
	fmt.Println(">問題を選択して下さい。")
	fmt.Scanf("%s", &sel)

	switch sel {
	case "practiceA":
		fmt.Println(">this is practiceA")
		BeginnersSelection.PracticeA()
	case "abc086a":
		fmt.Println(">this is abc086a")
		BeginnersSelection.Abc086a()
	case "abc081a":
		fmt.Println(">this is abc081a")
		BeginnersSelection.Abc081a()
	case "abc081b":
		fmt.Println(">this is abc081b")
		BeginnersSelection.Abc081b()
	}

}

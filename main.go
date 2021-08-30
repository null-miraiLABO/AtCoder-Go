package main

import (
	"atCoder/BeginnersSelection"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var sel string

	fmt.Println(">ファイル一覧")
	list("BeginnersSelection/")

	fmt.Println(">問題を選択して下さい。")
	fmt.Scanf("%s", &sel)
	selection(sel)

}

func list(path string) {
	cmd := exec.Command("ls", path)
	var ls_result, err = cmd.Output()

	var str string = string(ls_result)

	slice := strings.Split(str, "\n")
	//printじゃなくて変数に入れるよう改良したい
	for _, str := range slice {
		if str != "" {
			str = strings.Replace(str, ".go", "", -1)
			fmt.Println(strings.Replace(str, str[:1], strings.ToUpper(str[:1]), -1))
		}
	}
	fmt.Println("\nerr : ", err)
}

func selection(sel string) {
	switch sel {
	case "PracticeA":
		fmt.Println(">this is practiceA")
		BeginnersSelection.PracticeA()
	case "Pbc086a":
		fmt.Println(">this is abc086a")
		BeginnersSelection.Abc086a()
	case "Abc081a":
		fmt.Println(">this is abc081a")
		BeginnersSelection.Abc081a()
	case "Abc081b":
		fmt.Println(">this is abc081b")
		BeginnersSelection.Abc081b()
	}
}

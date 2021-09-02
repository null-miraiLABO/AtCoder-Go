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
	list_view("BeginnersSelection/")

	fmt.Println(">問題を選択して下さい。")
	fmt.Scanf("%s", &sel)
	run(sel)

}

func list_view(path string) {
	cmd := exec.Command("ls", path)
	var ls_result, err = cmd.Output()

	var str string = string(ls_result)

	slice := strings.Split(str, "\n")
	for _, str := range slice {
		if str != "" {
			str = strings.Replace(str, ".go", "", -1)
			fmt.Println(str)
		}
	}
	fmt.Println("\nerr : ", err)
}

func run(sel string) {
	fmt.Printf(">this is %s \n", sel)

	switch sel {
	case "practiceA":
		BeginnersSelection.PracticeA()
	case "abc086a":
		BeginnersSelection.Abc086a()
	case "abc081a":
		BeginnersSelection.Abc081a()
	case "abc081b":
		BeginnersSelection.Abc081b()
	case "abc087b":
		BeginnersSelection.Abc087b()
	default:
		fmt.Printf(">%s does not exist  \n", sel)
	}

}

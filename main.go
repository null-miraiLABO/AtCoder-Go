package main

import (
	"atCoder/BeginnersSelection"
	"fmt"
	"os/exec"
	"strings"
)

var dir_path string = "BeginnersSelection/"

func main() {
	var sel string
	var command, pkg_path string

	fmt.Println(">ファイル一覧")
	list_view(dir_path)

	fmt.Println(">問題を選択して下さい。")
	fmt.Printf("name: ")
	fmt.Scanf("%s", &sel)
	if sel != "q" {
		run(sel)
	}

	fmt.Printf("cmd: ")
	fmt.Scan(&command, &pkg_path)
	command_run(command, pkg_path)

}

func command_run(cmd, str string) {
	switch cmd {
	case "run":
		fmt.Println("this is run")
	case "cat":
		cat := exec.Command("cat", dir_path+str+".go")
		var cat_result, err = cat.Output()
		fmt.Printf("\n%s\nerr : %s\n", cat_result, err)
	case "list":
		list_view(dir_path)
	case "exit":
		fmt.Println("this is exit")
	default:
		fmt.Printf("command not found: %s \n", cmd+str)
	}

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
	case "abc083b":
		BeginnersSelection.Abc083b()
	case "abc088b":
		BeginnersSelection.Abc088b()
	default:
		fmt.Printf(">%s does not exist  \n", sel)
	}

}

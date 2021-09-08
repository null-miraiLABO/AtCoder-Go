package main

import (
	"atCoder/BeginnersSelection"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var dir_path string = "BeginnersSelection/"
var fl bool = true

func main() {
	var command, pkg_path string

	fmt.Println("[AtCoder code control tool with Go]")

	for i := 0; fl && i < 10; i++ {
		fmt.Printf("cmd: ")

		var sc = bufio.NewScanner(os.Stdin)
		sc.Scan()
		st := sc.Text()
		arr := strings.Split(st, " ")

		command = arr[0]
		if len(arr) > 1 {
			pkg_path = arr[1]
		} else {
			pkg_path = ""
		}

		command_run(command, pkg_path)
	}

}

func command_run(cmd, str string) {
	switch cmd {
	case "run":
		run(str)
	case "cat":
		cat := exec.Command("cat", dir_path+str+".go")
		var cat_result, err = cat.Output()
		fmt.Printf("\n%s\nerr : %s\n", cat_result, err)
	case "list":
		list_view(dir_path)
	case "exit":
		fl = false
	default:
		fmt.Printf("command not found: %s \n", cmd+str)
	}

}

func list_view(path string) {
	fmt.Println(">ファイル一覧")
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
	fmt.Printf(">running : %s\n", sel)

	switch sel {
	case "q":
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

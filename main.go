package main

import (
	"atCoder/BeginnersSelection"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// パッケージディレクトリのパス
var dir_path string = "BeginnersSelection/"

// ループフラグ
var fl bool = true

func main() {
	fmt.Println("[hakadoru!!]")
	fmt.Println("this is AtCoder code control tool with Go")

	// 動作部分。flがfalseで終了(100回の制限あり)
	for i := 0; fl && i < 100; i++ {
		fmt.Printf("hakadoru-command> ")

		var sc = bufio.NewScanner(os.Stdin)
		sc.Scan()
		st := sc.Text()
		arr := strings.Split(st, " ")

		command_run(arr)
	}
}

// 各コマンドの選択
func command_run(arr []string) {
	cmd := arr[0]
	switch cmd {
	case "run":
		run(arr)
	case "cat":
		cat(arr)
	case "list":
		list_view(dir_path)
	case "exit":
		fl = false
	default:
		fmt.Printf("command not found: %s \n", arr)
	}
}

// `cat`コマンド動作部分
func cat(arr []string) {
	var result string
	// sedオプション時
	if arr[1] == "-sed" {
		cat := exec.Command("cat", dir_path+arr[2]+".go")
		var cat_result, err = cat.Output()
		tmp := strings.Replace(dir_path, "/", "", -1)
		result = string(cat_result)
		result = strings.Replace(result, tmp, "main", -1)
		tmp = strings.Replace(arr[2], arr[2][:1], strings.ToUpper(arr[2][:1]), -1)
		result = strings.Replace(result, tmp, "main", -1)

		if err != nil {
			fmt.Println("err: ", err)
		}

	} else {
		// catオプション無い時
		cat := exec.Command("cat", dir_path+arr[1]+".go")
		var cat_result, err = cat.Output()
		result = string(cat_result)

		if err != nil {
			fmt.Println("err: ", err)
		}
	}
	fmt.Println(result)
}

// `list`コマンド動作部分
func list_view(path string) {
	cmd := exec.Command("ls", path)
	var ls_result, err = cmd.Output()
	var str string = string(ls_result)

	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("<ファイル一覧>")
	slice := strings.Split(str, "\n")
	for _, str := range slice {
		if str != "" {
			str = strings.Replace(str, ".go", "", -1)
			fmt.Println(str)
		}
	}
}

// `run`コマンド動作部分
// が増える度に追記する。
func run(arr []string) {
	sel := arr[1]
	fmt.Printf("running : %s\n", sel)

	switch sel {
	case "q": //quit
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

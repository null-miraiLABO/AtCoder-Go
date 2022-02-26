package main

import (
	"atCoder/BeginnersSelection"
	"atCoder/practice"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var contestDirPath string = "BeginnersSelection/"

// ループ終了フラグ
var exitFlag bool = true

func main() {
	fmt.Println(" hakadoru へようこそ")
	fmt.Println(" hakadoru はAtCoderコード管理ツールです。")
	fmt.Println(" 試しに `man` コマンドでマニュアルを見てみましょう！ ")
	fmt.Printf("\n\n")

	// コマンド入力動作部分。exitFlagが`exit`コマンドでfalseで終了(100回の制限あり)
	for i := 0; exitFlag && i < 100; i++ {
		fmt.Printf("hakadoru-command > ")

		var sc = bufio.NewScanner(os.Stdin)
		sc.Scan()
		st := sc.Text()
		// arr[1]にコマンド,arr[2]に問題名またはオプション,arr[3]に問題名または空
		arr := strings.Split(st, " ")

		command_run(arr)
	}
}

// 各コマンドの選択
func command_run(arr []string) {
	cmd := arr[0]
	switch cmd {
	case "man":
		manual()
	case "run":
		run(arr)
	case "cat":
		cat(arr)
	case "list":
		list_view(contestDirPath)
	case "exit":
		exitFlag = false
	default:
		fmt.Printf("command not found: %s \n", arr)
	}
}

// `cat`コマンド動作部分。指定した問題のコードをcatする。
func cat(arr []string) {
	var result string
	// sedオプション時
	if arr[1] == "-sed" {
		cat := exec.Command("cat", contestDirPath+arr[2]+".go")
		var cat_result, err = cat.Output()
		tmp := strings.Replace(contestDirPath, "/", "", -1)
		result = string(cat_result)
		result = strings.Replace(result, tmp, "main", -1)
		tmp = strings.Replace(arr[2], arr[2][:1], strings.ToUpper(arr[2][:1]), -1)
		result = strings.Replace(result, tmp, "main", -1)

		if err != nil {
			fmt.Println("err: ", err)
		}

	} else {
		// catオプション無い時
		cat := exec.Command("cat", contestDirPath+arr[1]+".go")
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
// 問題が増える度にcaseを追記しなければならない。文字リテラルの中で関数実行出来ない為
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
	case "abc085b":
		BeginnersSelection.Abc085b()
	case "abc085c":
		BeginnersSelection.Abc085c()
	case "abc049c":
		BeginnersSelection.Abc049c()
	case "anotoverflow":
		practice.Anotoverflow()
	default:
		fmt.Printf(">%s does not exist  \n", sel)
	}
}

func manual() {
	var result string

	// manual.txtをstrに入れる
	str, err := os.Open("manual.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer str.Close()

	// strに"\n"を入れ、整えたものをresurtに格納
	tmp_slice := make([]byte, 1024)
	for {
		n, err := str.Read(tmp_slice)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		result += string(tmp_slice[:n]) + "\n"
	}
	fmt.Println(result)
}

# AtCoder-Go
Goで書いたコードを保存しておくやつ

## AtCoderが捗る！なんちゃってcliツールを作る(途中)
### 前書き
問題別パッケージを管理するツールを作りました。

リポジトリ:[https://github.com/null-miraiLABO/AtCoder-Go](https://github.com/null-miraiLABO/AtCoder-Go)

### 仕様
#### 流れ
1.atCoderディレクトリ上で `main.go` を実行。

2.stockディレクトリの中のパッケージ一覧を表示

3.マニュアルを表示

**4.コマンドの実行**
パッケージ名は、「stockディレクトリの中のパッケージ一覧を表示」を参考に
```diff
//選択したパッケージを実行
run パッケージ名

//選択したパッケージのコードを表示
cat パッケージ名 

//選択したパッケージをコードを文字列置換し表示
sed パッケージ名
  `package パッケージ名 -> package main` //読み込みパッケージ
  `func パッケージ名() -> func main()` //関数名

//終了。
exit
```

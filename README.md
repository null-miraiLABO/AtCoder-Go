# AtCoder-Go -hakadoru-

Go で書いたコードを保存しておくやつ

## AtCoder が捗る！なんちゃって cli ツールを作る

### 前書き

問題別パッケージを管理するツールを作りました。

[実行例:Youtubeリンク](https://www.youtube.com/watch?v=AfbT2hg7_p0)

### 仕様

#### 流れ

1.atCoder ディレクトリ上で `go run main.go` をする。

2.各コマンドの実行

3.`exit`で終了

#### コマンドマニュアル

```diff
//実装済みの全コマンド

  man //マニュアルの表示

  list //パッケージ内ファイル一覧

  run [パッケージ名] //選択したパッケージを実行

  cat [オプション] [パッケージ名] //選択したパッケージのコードを表示
    -sed //選択したパッケージをコードをAtCoder提出用に文字列置換し表示

  exit //終了。

//記入例
  cat -sed abc083b

```

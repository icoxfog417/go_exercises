# Go exercises

Go言語の基礎的な処理を学ぶためのリポジトリです。

## Setup

以下のソフトウェアをインストールしておいてください

* [Git](https://git-scm.com/)
* [Go](http://golang-jp.org/)
  * `GOPATH`などの設定を含め、[こちらの手順](http://www.golangbootcamp.com/book/get_setup)が一番確実です
* プログラムを開発するための、好きなエディタ
  * 普段IntelliJ Ideaを使っているなら、[go-lang-idea-plugin](https://github.com/go-lang-plugin-org/go-lang-idea-plugin)の導入がお手軽です

Gitの細かい使い方については触れないため、困ったときは以下のサイトなどを参考にしてみてください。

* Git
  * [Git チュートリアル](https://www.atlassian.com/ja/git/tutorial/git-basics)
  * [使い始める Git](http://qiita.com/icoxfog417/items/617094c6f9018149f41f)

## How to begin

1. Starを押してから(m(_ _)m)、右上にある「fork」というボタンから、本リポジトリをfork(=コピー)してください。
2. forkしたリポジトリを、git cloneによって手元の端末に取ってきます。これで準備は完了です。
3. 各フォルダの中にあるREADMEには、そのエクササイズで達成すべき課題が書かれています。Pointで示されているヒントをもとに、処理を完成させてください。
4. 処理は、各フォルダ内に`ex01.go`というような`ex`で始まる名前のファイルを作成してその中に実装してください(`ex`始まりのファイルは`.gitignore`されています)
5. 各フォルダには、解答例が配置されています。わからない場合は、そちらを参考にしてください。

## Exercises

* [Let's write Golang](https://github.com/icoxfog417/go_exercises/tree/master/01_golang_basic)
  * Go言語の基本的な文法を理解しよう
* Collections
  * Go言語におけるコレクションの種類と使いかたを理解しよう
* Find Pattern
  * 文字列の中から特定のパターンを見つけてみよう
* Interface
  * Go言語におけるクラスの考え方とも言える、型を理解しよう
* User Packages
  * 便利なライブラリを使って処理を実装してみよう
* Access Web API
  * Go言語でHTTPリクエストを送る方法を理解しよう

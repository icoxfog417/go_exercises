# 02.Find Pattern

## Exercise(ex03.go)

このファイル(`README.md`)を読み込んで、'##'で始まっている見出しを抽出してください。上手くいけば、以下が得られるはずです。

* Exercise(ex03.go)
* Point

## Point

**File Read**

* [`os.Open`](https://golang.org/pkg/os/#Open)を利用することでファイルにアクセスできます。
* [`os.OpenFile`](https://golang.org/pkg/os/#OpenFile)を利用することで、ファイルアクセス時のパーミッションをより厳密に指定できます。
  * Permissionの設定は、[こちら](http://stackoverflow.com/questions/14249467/os-mkdir-and-os-mkdirall-permission-value)が参考になります。
* [bufio](https://golang.org/pkg/bufio/)や[io/ioutil](https://golang.org/pkg/io/ioutil/)に用意されているReader/Writerを利用することで、より効率的な読み込み/書き込みが可能です。
  * 各Reader/Writerの違いについては、[こちら](http://www.yunabe.jp/docs/golang_io.html)によくまとまっています。

**Pattern Matching**

* 正規表現によるマッチングを行う際は、[`regexp`](https://golang.org/pkg/regexp/)を利用します。
* パターンのコンパイル(`MustCompile`)にはそれなりに時間がかかるようなので、繰り返し処理では複数回行わないよう注意しましょう

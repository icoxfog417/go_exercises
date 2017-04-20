# 01. Let's write Golang

## Exercise(ex01.go)

引数から受け取った数だけ、`GOGO!`と出力する処理を実装してください(引数がない場合は、`YOU NEED TO SET COUNT!`と出力します)。

```
>01_golang_basic/ex01.go -count 3
GOGO!
GOGO!
GOGO!
```

```
>01_golang_basic/ex01.go
YOU NEED TO SET COUNT!
```

ただし、引数が10より大きい場合はその値が`TOO MATCH!`と出力します。

```
>01_golang_basic/ex01.go -count 100
100 is TOO MATCH!
```

## Point

**Go Style**

Go言語ではフォーマットは結構厳しくチェックされ、そもそもコンパイルをする際(`go run`や`go build`)にもある程のチェックが行われます。  
さらに、[`go vet`](http://golang-jp.org/pkg/code.google.com/p/go.tools/cmd/vet/)を利用することでより厳密なチェックをかけることができます。また、[gofmt](http://golang-jp.org/cmd/gofmt/)を利用することでソースコードを整形することができます。  
多くの統合開発環境ではこれらのツールをサポートしているので、Go言語を書く際は有効化して綺麗なソースを書きましょう！

**Package**

Go言語はpackageの宣言から始まります。packageはプログラムをモジュール化する単位となります。`main`は特別なパッケージで、プログラム実行時の起点となります(`go run`や`go build`などで実行、実行ファイルを作成するときには必ず必要になります)。

同じpackage内の関数や変数には自由にアクセス可能です。  
外部のpackageを利用する際は、`import`で利用を宣言して使用します。このとき外部からアクセスできるのは、[**頭文字が大文字の変数・関数のみ**](https://go-tour-jp.appspot.com/basics/3)になります。頭文字を大文字にする=publicにする、という理解で良いと思います。

**Execution & Compile**

`.go`の拡張子のプログラムを、`go run`により実行でききます。ただ、先述の通り、実行できるのは`main`パッケージであるプログラムのみです。　　
`go build`によりコンパイル、また`go intall`によりコンパイル及びインストールを行うことができます。インストールは、`go build`によって生成された実行ファイルを`$GOPATH/bin`(`GOBIN`)へ配置します。

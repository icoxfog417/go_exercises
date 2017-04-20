# 02.Collections in Golang

## Exercise(ex02.go)

引数で受け取った年齢の人の名前を変えすプログラムを作成してください。名前と年齢のリストは以下の通りです。

```go
var nameAges = map[string]int{
    "Andy":  20,
    "Bob":   21,
    "Cony":  21,
    "Debit": 20,
    "Elphy": 21,
    "Fox":   23,
}
```



## Point

Go言語におけるコレクションは、以下3つになります。

* Array: 長さの決まったコレクション(ex: `var x[10]int`)
* Slice: 長さの決まっていないコレクション(リスト)(ex: `[]int{1, 2}`)
* map: key-valueによるデータの格納(ex: `map[string]int{"age", 20}`)

そして、長さを取得するときは`len`、繰り返し処理を行うときは`range`を使用します。`range`は、Array/Sliceの場合はインデックスとその中の要素、mapの場合はkey/valueのペアを返します。

また、sliceに要素を追加する場合は`append`を利用します。


```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```

from [GO BOOTCAMMP Chapter4 Collection Types](http://www.golangbootcamp.com/book/collection_types)

なお、変数の代入に`=`や`:=`が混じっていて気になると思いますが、`:=`は変数の宣言と型の指定を省略し、代入される値から推定させることができます。

```go
var i int
j := i // j is an int
```

from [A Tour of Go: Type inference](https://tour.golang.org/basics/14)

また、出力時は`%q`を使うとフォーマットをしてくれます。

```go
package main

import "fmt"

func main() {
	a := [...]string{"hello", "world!"}
	fmt.Printf("%q", a)
}
```

from [GO BOOTCAMMP Chapter4 Collection Types](http://www.golangbootcamp.com/book/collection_types)

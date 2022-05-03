package mylib

import "fmt"

var Public string = "public!"   // main.goなど外部パッケージで使用したい場合は、先頭大文字。（structや関数名も同様）
var private string = "private!" // 先頭小文字だと外部パッケージでは参照できない。

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human!")
}

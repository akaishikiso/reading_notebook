package main

import (
	"fmt"
)

type NumberBox struct {
	value int
}

func main() {
	number := &NumberBox{value: 10}
	updateValue(number)
	fmt.Println(number.value) // 20

	replaceBox(number)
	fmt.Println(number.value) // 20（replaceBox の中で代入されたが元の number は変わっていない）

	number = replaceBoxForInsert(number)
	fmt.Println(number.value) // 30（replaceBoxForInsert の中で新しい構造体が作られ、返り値として受け取った）

	var emptyBox *NumberBox = nil
	replaceBox(emptyBox)
	fmt.Println(emptyBox == nil) // true（replaceBox 内で代入されても元の emptyBox は変わらない）

	// ポインタを通して値を変更する例
	var x int = 10
	fmt.Println(x) // 10
	failedUpdate(&x)
	fmt.Println(x) // 10（値は変更されない）

	successUpdate(&x)
	fmt.Println(x) // 20（値が変更される）
}

// ポインタを通して value を直接変更する
func updateValue(box *NumberBox) {
	box.value = 20
}

// 引数として受け取ったポインタ変数に新しい構造体を代入（呼び出し元には影響しない）
func replaceBox(box *NumberBox) {
	if box != nil {
		fmt.Println(*box) 
	} else {
		fmt.Println("box is nil")
	}
	box = &NumberBox{value: 30}
	fmt.Println(*box)
}

// 引数として受け取ったポインタ変数に新しい構造体を代入し、返り値として新しいポインタを返す
func replaceBoxForInsert(box *NumberBox) *NumberBox {
	if box != nil {
		fmt.Println(*box) 
	} else {
		fmt.Println("box is nil")
	}
	box = &NumberBox{value: 30}
	fmt.Println(*box)
	return box // 返り値として新しいポインタを返す
}

func failedUpdate(x *int) {
	x2 := 20 // 新しい変数を作成
	x = &x2 // 引数のポインタを新しい変数のポインタに変更
	// しかし、呼び出し元のポインタは変更されない
	// つまり、x の値は変更されない
	// これは、ポインタの値を変更しているのではなく、ポインタ自体を新しいアドレスに変更しているため
	// 呼び出し元のポインタは依然として元のアドレスを指している
}

func successUpdate(x *int){
	*x = 20 // ポインタを通して値を変更
}

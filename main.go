package main

import "fmt"

func main() {
	var aaa int = 1_000_000
	var bbb = 2
	fmt.Printf("%d, %d\n", aaa, bbb)
	// aaa = 2
	fmt.Printf("Hello, World\n")
	var price int
	fmt.Print("値段は")
	fmt.Scan(&price)
	fmt.Printf("%d円dayo-\n", price)
	const message string = "Hello, World!"
	fmt.Println(message)
	sample1()
	sample2()
	sample3()
}

func sample1() {
	const (
		c = 1 << iota
		d
		e
		f
	)
	fmt.Println("サンプルだよ")
	fmt.Printf("%d, %d, %d, %d\n", c, d, e, f)
}

func sample2() {
	num := 3
	fmt.Println("サンプルだよ")
	fmt.Printf("%d\n", num)
}

func sample3() {
	// ifの中でのみ使える値「num」（変数名はなんでもOK）を宣言できる
	if num := 3; num > 0 {
		println("numは整数！")
	} else {
		println("numは負の値！")
	}
	// コンパイルエラー
	// undefined: num
	// println(num)
}

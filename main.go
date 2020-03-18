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

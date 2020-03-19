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
	sample5()
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

func sample4() {
	num := 5
	switch num {
	case 3, 5:
		fmt.Println("3か5")
		// breakいらない
		// 連鎖させたい場合は、「fallthrough」を使う
	default:
		fmt.Println("3じゃないし、５じゃない")
	}
}

func sample5() {
	fmt.Println("ーーーーーーーーーーーーーーー")
	fmt.Println("少し長いforが流れます")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}
}

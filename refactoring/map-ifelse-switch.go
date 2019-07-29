package main

import "fmt"

func on1() {
	fmt.Println("1,2")
	fmt.Println("1,2")
	fmt.Println("1,2")
}

func on3() {
	fmt.Println("3")
	fmt.Println("3")
	fmt.Println("3")
}

func on4() {
	fmt.Println("4")
	fmt.Println("4")
	fmt.Println("4")
}

func bad1(n int) {
	switch n {
	case 1, 2:
		on1()
	case 3:
		on3()
	case 4:
		on4()
	}
}

func good1(n int) {
	inHandlers := map[int]func(){
		1: on1,
		2: on1,
		3: on3,
		4: on4,
	}
	if handler, ok := inHandlers[n]; ok {
		handler()
	}
}

func main() {
	bad1(2)
	fmt.Println("--------------")
	good1(2)
	fmt.Println("--------------")
}

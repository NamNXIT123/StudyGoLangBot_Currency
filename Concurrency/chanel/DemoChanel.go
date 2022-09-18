package chanel

import "fmt"

func DemoChanel() {
	var kenh1 chan int
	fmt.Println("Kenh1:", kenh1)

	if kenh1 == nil {
		fmt.Println("channel a is nil, going to define it")
		kenh1 = make(chan int)
		fmt.Printf("Type of a is %T", kenh1)
		fmt.Println()
		fmt.Println("Kenh1 after Make(chan T):", kenh1)
	}

	//
	kenh2 := make(chan string)
	fmt.Println("kenh2", kenh2)
}

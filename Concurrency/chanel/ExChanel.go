package chanel

import "fmt"

func calcSquares(number int, chanel1 chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	chanel1 <- sum
}

func calcCubes(number int, chanel2 chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	chanel2 <- sum
}

func readChanelDealock(chanel3 chan int) {
	//fmt.Println(<-chanel3)
	chanel3 <- 19
	fmt.Println("AAAAAAAAAAAAAAAA")
}

func EXChanel() {
	/*number := 589
	chanel1 := make(chan int) // Create Chanel1
	chanel2 := make(chan int) // Create Chanel2
	//
	go calcSquares(number, chanel1)
	go calcCubes(number, chanel2)
	//
	//dataRecevei1 := <-chanel1
	//fmt.Println("dataRecevei1", dataRecevei1)
	squares, cubes := <-chanel1, <-chanel2
	fmt.Println("Final output", squares+cubes)*/

	ch := make(chan int)
	//ch <- 5
	fmt.Println("AAAAAAAAAAAAAAA")
	go readChanelDealock(ch)
	fmt.Println(<-ch)
}

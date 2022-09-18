package chanel

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		fmt.Println("aaaaaaa")
		number /= 10
		fmt.Println("number:", number)
	}
	close(dchnl)
}
func calcSquares1(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
	fmt.Println("calcSquares1")
}

func calcCubes1(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
	fmt.Println("calcCubes1")
}

func EXCloseChanel() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)

	///
	go calcSquares1(number, sqrch)
	go calcCubes1(number, cubech)
	//
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

package chanel

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	fmt.Println("Ghi thành công") // Không có giá trị vì hàm dừng lại tại  sendch <- 10
}

func DemoChanel1Chieu() {
	sendch := make(chan<- int) // Kênh chỉ ghi
	go sendData(sendch)
	//fmt.Println(<-sendch)  // Không thể đọc được do Chanel chỉ ghi / không cho đọc
}

func sendData1(sendch chan<- int) { // Tham số đầu vào 1 chanel có thể 1 chiều
	sendch <- 10
}

func DemoChanelChange() {
	chnl := make(chan int) // chanel 2 chiều
	go sendData1(chnl)     //
	fmt.Println(<-chnl)
}

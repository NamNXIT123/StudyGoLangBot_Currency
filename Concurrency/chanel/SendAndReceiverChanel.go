package chanel

import (
	"fmt"
	"time"
)

func SendAndReceiver() {
	chanelBool := make(chan bool)
	// Goroutine
	go hello(chanelBool)
	//
	fmt.Println("Truowcs khi kênh bị chặn")
	//
	data := <-chanelBool // Dòng mã này đang chặn có nghĩa là cho đến khi một số Goroutine ghi dữ liệu vào chanelBool, điều khiển sẽ không chuyển sang dòng mã tiếp theo.
	//Do đó, điều này loại bỏ sự cần thiết phải time.Sleepcó trong chương trình ban đầu để ngăn Goroutine chính thoát ra.
	//
	/*for i := 0; i < 100; i++ {
		fmt.Println("main:", i)
	}*/

	//
	fmt.Println("data1 read Chanel:", data)
	data2 := <-chanelBool
	fmt.Println("data2 read Chanel:", data2)

	time.Sleep(5 * time.Second)
	data3 := <-chanelBool
	fmt.Println("data3 read Chanel:", data3)
	//data4 := <-chanelBool
	//fmt.Println("data3 read Chanel:", data4)  // DEALOCK

	//
}

func hello(chanelBool chan bool) {
	fmt.Println("Hello world goroutine")
	//time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	//chanelBool <- true // Write data to Chanel
	//for i := 0; i < 100; i++ {
	//fmt.Println("hello", i)
	chanelBool <- true // Write data to Chanel
	//}
	//fmt.Println("truyen Data vao kenh xong")
	chanelBool <- false
	chanelBool <- true

	fmt.Println("truyen Data vao kenh xong")

}

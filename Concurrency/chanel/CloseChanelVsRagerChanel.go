package chanel

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl) // Close kênh // Sử dụng cho vòng lặp khi đọc từ Kênh
}
func ChanelCloseVsRager() {
	ch := make(chan int)
	go producer(ch)
	// Duyệt qua kênh
	/*for {
		v, ok := <-ch    // Câu lệnh đọc dữ liệu ở Kênh // Nếu kênh không được đóng thì sẽ dẫn gawpj dealLock
		if ok == false { //Nếu oksai có nghĩa là kênh bị đóng Close(chanel)
			break
		}
		fmt.Println("Received ", v, ok)
	}*/

	// Sử dụng for Range
	for v := range ch {
		fmt.Println("Received ", v)
	}

}

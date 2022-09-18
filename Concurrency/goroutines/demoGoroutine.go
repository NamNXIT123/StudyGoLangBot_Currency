package goroutines

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

/*
Khi một Goroutine mới được bắt đầu, lệnh gọi goroutine sẽ trả về ngay lập tức.
Không giống như các hàm, điều khiển không đợi Goroutine kết thúc việc thực thi.
Điều khiển trả về ngay lập tức dòng mã tiếp theo sau lệnh gọi Goroutine và
mọi giá trị trả về từ Goroutine đều bị bỏ qua.

Goroutine chính sẽ được chạy cho bất kỳ Goroutines nào khác để chạy.
Nếu chương trình Goroutine chính kết thúc thì chương trình sẽ bị kết thúc và không có chương trình Goroutine
nào khác sẽ chạy.
*/
func GoroutinesDemo() {
	go hello()
	//time.Sleep(time.Second * 2)
	fmt.Println("main function")
	time.Sleep(time.Second * 2)
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
func MutilGotroutines() {
	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}

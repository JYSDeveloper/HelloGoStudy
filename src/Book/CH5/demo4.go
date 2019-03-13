package main

import "runtime"

//模拟消息队列
func main() {
	stop := make(chan struct{})
	memory := make(chan string,100)
	go GetMessage(stop,memory)
	println("Number of Goroutine = ",runtime.NumGoroutine())
	<-stop
	for v := range memory {
		println(v)
	}
}

func GetMessage(stop chan struct{},memory chan string )  {
	for i := 1; i<=100;i++  {
		s := "Message  i"
		memory <- s
	}
	close(memory)
	stop <- struct{}{}
}

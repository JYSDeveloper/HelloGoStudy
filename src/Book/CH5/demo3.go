package main

import (
	"runtime"
	"time"
)

func main() {
	c := make(chan struct{})
	go func(t chan struct{}) {
		sum := 0
		for  i:=0;i<100000;i++{
			sum+=i
		}
		time.Sleep(1*time.Second)
		println(sum)
		t <- struct{}{}
	}(c)
	println("当前线程数",runtime.NumGoroutine())
	//执行结束后，才能取出结果，否则会阻塞在这里，达到同步等待
	<-c
}

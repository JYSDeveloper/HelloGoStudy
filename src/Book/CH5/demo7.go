package main

import (
	"math/rand"
	"runtime"
)

//Exit 模式

func main() {
	exit := make(chan struct{})
	ch := GeneratorIntValue(exit)
	for i := 0; i <= 100; i++ {
		println(<-ch)
	}
	exit <- struct{}{} // 等价于 close(exit)
	println("number of goroutine = " ,runtime.NumGoroutine())
}

func GeneratorIntValue(exit chan struct{}) chan int{
	ch := make(chan int,5) // 每次生产5个
	go func() {
		Label:
			for {
				select {
					case ch <- rand.Intn(100):
					case <- exit:
						break Label
				}
			}
		close(ch)
	}()
	return ch
}

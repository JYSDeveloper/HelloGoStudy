package main

import (
	"fmt"
	"time"
)
var storage = make(chan string,100)
var consumer1 = make(chan string,1)
var consumer2 = make(chan string,1)
func main() {
	go MessageGenerator(storage) //生产100个消息
	go Consumer(consumer1) // c1 开始消费
	go Consumer(consumer2) // c2 开始消费
	go Work(consumer2,consumer1)
	time.Sleep(5*time.Second)
}

func Work(a,b chan string)  {
	for{
		v := <- storage
		select {
			case a <- v:
			case b <- v:
		}
/* 代码换成这段，就发现消息丢失了，因为case的时候，storage取出的操作先执行，
	然后判断ab哪个可以接受，会导致取了两个消息，但是只有一个分配给ab了。
		select {
		case a <- <- storage:
		case b <- <- storage:
		}
*/
	}
}

func Consumer(temp chan string)  {
	for {
		println(<- temp,temp)
	}
}

func MessageGenerator(storage chan string)  {
	for i := 0; i < 10;i++ {
		storage <- fmt.Sprintf("Message %d",i)
	}
}

package main
//select

func main() {
	ch := make(chan int ,1)
	go func(c chan int) {
		for {
			select {
				//这里会随机放入一个数字
				case ch <-1:
				case ch <-0:
				case ch<-2:
			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		println(<-ch)
	}
}


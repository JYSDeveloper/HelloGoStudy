package main

import "time"

func main()  {
	//异步执行方法
	go sum()
	println("End of Main func")
	time.Sleep(5* time.Second)
}
func sum()  {
	sum := 0
	for i := 0; i<=100	;i++  {
		sum += i
	}
	println(sum)
}

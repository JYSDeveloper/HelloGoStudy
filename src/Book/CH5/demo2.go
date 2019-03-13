package main

import "runtime"

//获取线程数的内容

func main() {
	println("当前线程数：",runtime.GOMAXPROCS(0))
	//大于等于0的参数，就是修改线程数量
	runtime.GOMAXPROCS(2)
	println("当前线程数：",runtime.GOMAXPROCS(0))
}

package main

import "HelloGoStudy/src/SelfDrive/DesignPattern"

func main() {
	ch := DesignPattern.GetCacheInstance()
	cho := DesignPattern.GetCacheInstanceOnce()
	println(ch)
	ch2 := DesignPattern.GetCacheInstanceOnce()
	println(ch)
	println(ch2)
	println(cho)

}

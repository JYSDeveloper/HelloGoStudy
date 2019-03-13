package main

import (
	"net/http"
	"sync"
)

//Wait

var wg sync.WaitGroup
var visitUrls = []string{
	"https://www.baidu.com",
	"https://www.google.com",
	"https://www.sina.com",
	"https://www.qq.com",
}
func main() {
	for _,v := range visitUrls {
		wg.Add(1)
		go visit(v)
	}
	wg.Wait()
}

func visit(url string)  {
	defer wg.Done() //等价于 wg.Add(-1)
	resp,error := http.Get(url)
	if error == nil{
		println(resp.Status)
	}else {
		println(error.Error())
	}
}

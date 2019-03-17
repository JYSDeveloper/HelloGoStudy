package main

import "time"

// future 模式

type Query struct {
	sql chan string
	result chan string
}

func execQuery(q Query)  {
	for  {
		go func() {
			sql := <-q.sql
			//这里应该访问数据库
			q.result <- "result of :" + sql
		}()
	}

}

func main() {
	q := Query{make(chan string ,2), make(chan string ,2)}
	go execQuery(q)

	q.sql <- "select * from student"
	q.sql <- "select * from teacher"

	//Do some thins
	time.Sleep(1 * time.Second)

	println(<-q.result)
	println(<-q.result)
}



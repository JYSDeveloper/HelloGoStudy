package main

import (
	"runtime"
	"sync"
)

const ResolveNumber = 100
//模拟拆分一批任务
type Task struct {
	Begin int
	End int
	Result chan<- int
}

func (t *Task) do()  {
	sum := 0
	for i := t.Begin; i <=t.End ;i++  {
		sum +=i
	}
	t.Result <- sum
}

func main() {
	taskChain := make(chan Task,100)
	resultChain := make(chan int ,100)
	sync := &sync.WaitGroup{}
	go InitTask(taskChain,resultChain,10000)
	go DistributeTask(taskChain,sync,resultChain)
	println("num:",runtime.NumGoroutine())

	sum := ProcessResult(resultChain)
	println(sum)
}

func InitTask(taskChain chan Task, resultChain chan int, count int)  {
	num := count / ResolveNumber
	mod := count % ResolveNumber
	for i := 0; i < num; i++ {
		b := ResolveNumber * i + 1
		e := ResolveNumber * (i + 1)
		task := Task{
			Begin:b,
			End:e,
			Result:resultChain,
		}
		taskChain <- task
	}
	if mod != 0{
		last := count * 10
		task := Task{
			Begin:last+1,
			End:count,
			Result:resultChain,
		}
		taskChain <- task
	}
	close(taskChain)
}

func DistributeTask(taskChain <-chan Task, sync *sync.WaitGroup, resultChain chan int){

	for v := range taskChain {
		sync.Add(1)
		go Process(v,sync)
	}
	sync.Wait()
	close(resultChain)
}

func Process(t Task, sync *sync.WaitGroup){
	t.do()
	sync.Done()
}

func ProcessResult(resultChain chan int) int{
	sum := 0
	for e := range resultChain {
		println(e)
		sum += e
	}
	return sum
}

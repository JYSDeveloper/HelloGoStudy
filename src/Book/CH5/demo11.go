package main

//固定10个线程处理任务，完成之后推出，可以设计成无线循环的模式
const threadNums  = 10
const ResolveNumbers = 100

type task struct {
	Begin int
	End int
	Result chan<- int
}

func (t *task) do()  {
	sum := 0
	for i := t.Begin; i <=t.End ;i++  {
		sum +=i
	}
	t.Result <- sum
}

func main() {

	taskChain := make(chan task,100)
	resultChain := make(chan int ,100)
	done := make(chan struct{},100)
	go InitTask2(taskChain,resultChain,10000)
	go DistributeTaskWithLimit(taskChain,threadNums,done)
	go CloseResult(done,resultChain,threadNums)
	rs := 0
	for e := range resultChain {
		rs+=e
	}
	println(rs)
}
func InitTask2(taskChain chan task, resultChain chan int, count int)  {
	num := count / ResolveNumbers
	mod := count % ResolveNumbers
	for i := 0; i < num; i++ {
		b := ResolveNumbers * i + 1
		e := ResolveNumbers * (i + 1)
		task := task{
			Begin:b,
			End:e,
			Result:resultChain,
		}
		taskChain <- task
	}
	if mod != 0{
		last := count * 10
		task := task{
			Begin:last+1,
			End:count,
			Result:resultChain,
		}
		taskChain <- task
	}
	close(taskChain)
}

func DistributeTaskWithLimit(task chan task, workers int, done chan struct{})  {
	for i := 0; i <= workers; i++ {
		go ProcessTask(task,done)
	}
}

func ProcessTask(task chan task, done chan struct{})  {
	for e := range task {
		e.do()
	}
	done <- struct{}{}
}

func CloseResult(done chan struct{},result chan int, workers int)  {
	for i := 0; i < workers; i++ {
		<- done
	}
	close(result)
	close(done)
}

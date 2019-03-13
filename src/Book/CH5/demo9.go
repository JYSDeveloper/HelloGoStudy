package main

func chain(in chan int)chan int  {
	out := make(chan int, cap(in))
	go func() {
		for v := range in {
			println("out")
			out <- v +1
		}
		close(out)
	}()
	return out
}

func main() {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			println("in")
		in <- i

	}
		close(in)}()
	r := chain(chain(chain(in)))
	for v := range r {
		println(v)
	}

}

package Context

import (
	"context"
	"fmt"
)

func CRun()  {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n:= 1
		go func() {
			for {
				select {
					case <-ctx.Done():
						return
						case dst <- n:
							n++
				}
			}
		}()
		return dst
	}
	ctx ,cancel := context.WithCancel(context.Background())
	defer cancel()
	for e := range gen(ctx) {
		fmt.Println(e)
		if e == 5{
			break
		}
	}
}

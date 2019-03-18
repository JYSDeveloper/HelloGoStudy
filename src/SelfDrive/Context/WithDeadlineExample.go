package Context

import (
	"context"
	"time"
)

func DRun()  {
	d := time.Now().Add(50 * time.Millisecond)
	ctx ,cancel := context.WithDeadline(context.Background(),d)
	defer cancel()

	select {
		case <-time.After(1 * time.Second):
			println("overslept")
		case <-ctx.Done():
				println(ctx.Err())
	}
}

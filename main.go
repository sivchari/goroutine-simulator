package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"time"
)

func pschan(ctx context.Context) {
	tick := time.NewTicker(500 * time.Millisecond)
	fmt.Println("start pschan")
	for {
		select {
		case <-tick.C:
			ps := runtime.AllPsSnapshot()
			fmt.Println("===== dump ps =====")
			for _, p := range ps {
				fmt.Printf("m: %+v\n", p.Machine)
				fmt.Printf("p: %+v\n", p)
			}
		case <-ctx.Done():
			return
		}
	}
}

func globalchan(ctx context.Context) {
	tick := time.NewTicker(2 * time.Second)
	fmt.Println("start globalchan")
	for {
		select {
		case <-tick.C:
			gs := runtime.GlobalRunq()
			fmt.Println("===== dump global goroutine =====")
			for _, g := range gs {
				fmt.Printf("goroutine: %+v\n", g)
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	go globalchan(ctx)
	go pschan(ctx)

	for range 1000 {
		go func() {
			time.Sleep(5 * time.Second)
		}()
	}

	<-ctx.Done()

	println("exit main")
}

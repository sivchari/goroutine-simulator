package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Second)
		}()
	}

	fmt.Println("===== dump ms =====")
	ms := runtime.AllMsSnapshot()
	for i, m := range ms {
		if i == 1 {
			fmt.Println("break")
			break
		}
		fmt.Printf("memstats: %+v\n", m)
	}

	fmt.Println("===== dump ps =====")
	ps := runtime.AllPsSnapshot()
	for i, p := range ps {
		if i == 1 {
			fmt.Println("break")
			break
		}
		fmt.Printf("p: %+v\n", p)
	}

	fmt.Println("===== dump global goroutine =====")
	gs := runtime.GlobalRunq()
	for _, g := range gs {
		fmt.Printf("goroutine: %+v\n", g)
	}

	fmt.Println("===== dump local goroutine =====")
	ls := runtime.LocalRunq()
	for _, l := range ls {
		fmt.Printf("local: %+v\n", l)
	}

	fmt.Printf("all len:\n global gs: %v\n local gs: %v\n ms: %v\n ps: %v\n", len(gs), len(ls), len(ms), len(ps))

	wg.Wait()
}

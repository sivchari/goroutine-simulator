package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func mschan() {
	tick := time.NewTicker(2 * time.Second)
	fmt.Println("start mschan")
	for {
		select {
		case <-tick.C:
			ms := runtime.AllMsSnapshot()
			fmt.Println("===== dump ms =====")
			for _, m := range ms {
				fmt.Printf("memstats: %+v\n", m)
			}
		}
	}
}

func pschan() {
	tick := time.NewTicker(2 * time.Second)
	fmt.Println("start pschan")
	for {
		select {
		case <-tick.C:
			ps := runtime.AllPsSnapshot()
			fmt.Println("===== dump ps =====")
			for _, p := range ps {
				fmt.Printf("p: %+v\n", p)
			}
		}
	}
}

func localchan() {
	tick := time.NewTicker(5 * time.Second)
	fmt.Println("start localchan")
	for {
		select {
		case <-tick.C:
			ls := runtime.LocalRunq()
			fmt.Println("===== dump local goroutine =====")
			for _, l := range ls {
				fmt.Printf("local: %+v\n", l)
			}
		}
	}
}

func globalchan() {
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
		}
	}
}

func main() {
	go localchan()
	go globalchan()
	go mschan()
	go pschan()

	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Second)
		}()
	}
	wg.Wait()
}

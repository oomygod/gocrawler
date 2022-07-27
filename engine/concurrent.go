package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		go CreateWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for o := range out {
		for _, item := range o.Items {
			itemCount++
			log.Printf("got item #%d %v", itemCount, item)
		}
		for _, r := range o.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func CreateWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	for {
		ready.WorkerReady(in)
		r := <-in
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		out <- parseResult
	}
}

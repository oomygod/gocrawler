package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)
	e.Scheduler.ConfigureWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		go CreateWorker(in, out)
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

func CreateWorker(in chan Request, out chan ParserResult) {
	for r := range in {
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		out <- parseResult
	}
}

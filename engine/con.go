package engine

import "log"

func ConRun(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParserResult)

	for i := 0; i < 50; i++ {
		go ConWorker(in, out)
	}

	for _, r := range seeds {
		in <- r
	}

	itemCount := 0
	for o := range out {
		for _, item := range o.Items {
			itemCount++
			log.Printf("got item #%d %v", itemCount, item)
		}
		result := o
		go func() {
			for _, r := range result.Requests {
				in <- r
			}
		}()
	}

}

func ConWorker(in chan Request, out chan ParserResult) {
	for r := range in {
		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		out <- parseResult
	}
}

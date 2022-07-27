package engine

import (
	"learn/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (s SimpleEngine) Run(seeds ...Request) {
	var reqQueue []Request
	for _, r := range seeds {
		reqQueue = append(reqQueue, r)
	}
	for len(reqQueue) > 0 {
		r := reqQueue[0]
		reqQueue = reqQueue[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		reqQueue = append(reqQueue, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("got item %v", item)
		}
	}
}

func Worker(r Request) (ParserResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url, 0)
	if err != nil {
		log.Printf("fetcher: error"+"fetching url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	parseResult := r.ParserFunc(body)
	return parseResult, nil
}

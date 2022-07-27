package scheduler

import "learn/crawler/engine"

type SimplerScheduler struct {
	workerChan chan engine.Request
}

func (s *SimplerScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimplerScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimplerScheduler) WorkerReady(requests chan engine.Request) {
}

func (s *SimplerScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

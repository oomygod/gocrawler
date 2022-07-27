package scheduler

import "learn/crawler/engine"

type SimplerScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimplerScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}

func (s *SimplerScheduler) ConfigureWorkerChan(requests chan engine.Request) {
	s.WorkerChan = requests
}

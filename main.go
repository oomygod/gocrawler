package main

import (
	"learn/crawler/engine"
	"learn/crawler/scheduler"
	"learn/crawler/zhenai/parser"
)

func main() {
	//1.单任务
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	//并发无调度器
	//engine.ConRun(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	//2.并发简单调度器
	//concurrentEngine := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.SimplerScheduler{},
	//	WorkerCount: 10,
	//}
	//concurrentEngine.Run(
	//	engine.Request{
	//		Url:        "http://www.zhenai.com/zhenghun",
	//		ParserFunc: parser.ParseCityList,
	//	},
	//)

	//3.并发队列调度器
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(
		engine.Request{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		},
	)

}

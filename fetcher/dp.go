package fetcher

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//const url = "https://album.zhenai.com/u/1677398857"
const remoteWs = "ws://127.0.0.1:9222/devtools/browser/28497aaf-f9a5-4bbf-92b0-e88b60ce5500"

func GetAllCookies(url string) string {
	//本地
	//opts := append(
	//	chromedp.DefaultExecAllocatorOptions[:],
	//	chromedp.Flag("headless", false),
	//)
	//
	//allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	//defer cancel()

	//远程
	allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), remoteWs)
	defer cancel()

	// 创建chrome实例
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()
	file, err := os.OpenFile("./cookies.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer file.Close()
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		//chromedp.WaitVisible(`app`, chromedp.ByID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			j, err := json.Marshal(cookies)
			if err != nil {
				return err
			}

			//【write】写入[]byte类型数据
			_, err = file.Write(j)
			if err != nil {
				return err
			}

			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	//time.Sleep(2 * time.Second)
	// 创建用于超时退出的上下文管理器
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	return file.Name()
}

func ReadCookie() []network.CookieParam {
	file, err := os.Open("./cookies.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	// 读取文件数据
	jsonBlob, _ := ioutil.ReadAll(file)
	//var cookies []*network.CookieParam
	var cookies []network.CookieParam
	// Json解码
	err = json.Unmarshal(jsonBlob, &cookies)
	if err != nil {
		panic(err)
	}
	//fmt.Println(cookies)
	return cookies
}
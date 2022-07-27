package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//var rateLimiter = time.Tick(100*time.Millisecond)
//var proxyUrl = "SOCKS5://127.0.0.1:1080"

func Fetch(target string, retry int) ([]byte, error) {
	//<- rateLimiter
	//proxy, _ := url.Parse(proxyUrl)
	//tr := &http.Transport{
	//	Proxy:           http.ProxyURL(proxy),
	//}
	client := &http.Client{
		//Transport: tr,
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		log.Fatalln(err)
	}
	//浏览器中的User-Agent
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	//
	if isProfileUrl(target) {
		req.Header.Add("cookie", GetCookie())
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Printf("[0]wrong status code %d n", resp.StatusCode,req.Cookies(),resp.Cookies())

		//202
		if resp.StatusCode == http.StatusAccepted {
			if isProfileUrl(target) {
				if retry < 0 { //重试一次
					SetCookie(target)
					retry++
					return Fetch(target, retry)
				}
			}
		}
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	//utf8Reader := transform.NewReader(resp.Body,charmap.ISO8859_1.NewDecoder())
	//return ioutil.ReadAll(utf8Reader)

	return ioutil.ReadAll(resp.Body)
}

func isProfileUrl(url string) bool {
	return strings.Contains(url, "album.zhenai.com/u/")
}

//func determineEncoding(r io.Reader) encoding.Encoding {
//	bytes,err := bufio.NewReader(r).Peek(1024)
//	if err!=nil {
//		panic(err)
//	}
//	e,_,_ := charset.DetermineEncoding(bytes,"")
//	return e
//}

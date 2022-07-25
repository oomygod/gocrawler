package parser

import (
	"fmt"
	"learn/crawler/fetcher"
	"strings"
	"testing"
)

func TestParseProfile(t *testing.T) {
	//url := "https://album.zhenai.com/u/1677398857"
	//url := "https://album.zhenai.com/u/1700301913"
	//url := "https://www.zhenai.com/zhenghun/aba"
	//url := "https://www.zhenai.com/zhenghun"
	for _,url := range data(){
		content, err := fetcher.Fetch(url,0)
		if err != nil {
			t.Logf("fetcher: error" + "fetching url %s: %v",url,err)
		}
		//t.Log(string(content))
		res := ParseProfile(content,"")
		fmt.Printf("%+v",res)
		break
	}
}

func data() []string {
	str := `https://album.zhenai.com/u/1700301913
https://album.zhenai.com/u/1162957428
https://album.zhenai.com/u/70043057
https://album.zhenai.com/u/1993147507
https://album.zhenai.com/u/1297598047
https://album.zhenai.com/u/1649645699
https://album.zhenai.com/u/1312469493
https://album.zhenai.com/u/1894144891
https://album.zhenai.com/u/1979306698
https://album.zhenai.com/u/1346722275
https://album.zhenai.com/u/1376996864
https://album.zhenai.com/u/1602294364
https://album.zhenai.com/u/1376989027
https://album.zhenai.com/u/1016698699
https://album.zhenai.com/u/1413454179
https://album.zhenai.com/u/1173578750
https://album.zhenai.com/u/1002457370
https://album.zhenai.com/u/80190751
https://album.zhenai.com/u/1045633034
https://album.zhenai.com/u/1982028238
https://album.zhenai.com/u/1783091548
https://album.zhenai.com/u/1677398857
https://album.zhenai.com/u/101658152
https://album.zhenai.com/u/103132583
https://album.zhenai.com/u/1596207608
https://album.zhenai.com/u/1544748367`
	arr := strings.Split(str,"\n")
	return arr
}

func TestCookie(t *testing.T) {
	url := "https://album.zhenai.com/u/1677398857"
	fetcher.SetCookie(url)
	cookie := fetcher.GetCookie()
	fmt.Println(cookie)
}

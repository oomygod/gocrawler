package parser

import (
	"learn/crawler/engine"
	"regexp"
	"strings"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		//变量拷贝一份，解决直接使用m[2]传入ParseProfile方法后值都一样的问题
		nickName := string(m[2])

		profileUrl := strings.Replace(string(m[1]),"http:","https:",1)
		result.Items = append(result.Items, "User " + nickName)
		result.Requests = append(result.Requests, engine.Request{
			Url:        profileUrl,
			ParserFunc: func(c []byte) engine.ParserResult {
					return ParseProfile(c,nickName)
			}})
	}
	return result
}

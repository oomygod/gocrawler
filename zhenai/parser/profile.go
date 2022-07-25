package parser

import (
	"github.com/PuerkitoBio/goquery"
	"learn/crawler/engine"
	"learn/crawler/model"
	"log"
	"strconv"
	"strings"
)

func ParseProfile(content []byte, name string) engine.ParserResult {
	profile := model.Profile{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		log.Fatalln(err)
	}

	//name := dom.Find(".nickName").Text()
	var purple []string
	dom.Find(".m-userInfoDetail .m-content-box .purple").Each(func(i int, selection *goquery.Selection) {
		purple = append(purple, selection.Text())
	})
	var pink []string
	dom.Find(".m-userInfoDetail .m-content-box .pink").Each(func(i int, selection *goquery.Selection) {
		pink = append(pink, selection.Text())
	})

	if len(purple) == 0 {
		return engine.ParserResult{}
	}

	//fmt.Println("profile：", name, purple[:3], pink[:3])
	profile.Name = name
	profile.Gender = getGender(dom)
	age, err := strconv.Atoi(strings.Replace(purple[1], "岁", "", 1))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(strings.Replace(purple[3], "cm", "", 1))
	if err == nil {
		profile.Height = height
	}
	if len(purple) == 9 {
		weight, err := strconv.Atoi(strings.Replace(purple[4], "kg", "", 1))
		if err == nil {
			profile.Weight = weight
		}
		profile.Income = purple[6]
		profile.Marriage = purple[0]
		profile.Education = purple[8]
		profile.Occupation = purple[7]
		profile.Hukou = purple[5]
		profile.Xingzuo = purple[2]
	}else if len(purple) == 8 {
		profile.Income = purple[5]
		profile.Marriage = purple[0]
		profile.Education = purple[7]
		profile.Occupation = purple[6]
		profile.Hukou = purple[4]
		profile.Xingzuo = purple[2]
	}else if len(purple) == 7 {
		profile.Income = purple[5]
		profile.Marriage = purple[0]
		profile.Education = purple[6]
		//profile.Occupation = purple[6]
		profile.Hukou = purple[4]
		profile.Xingzuo = purple[2]
	}
	profile.House = findOrEmpty(pink,5)
	profile.Car = findOrEmpty(pink,6)

	return engine.ParserResult{
		Items: []interface{}{profile},
	}
}

func getGender(dom *goquery.Document) string {
	gender := "未知"
	str := dom.Find(".m-userInfo .right").Text()
	if strings.Contains(str, "她") {
		gender = "女"
	} else if strings.Contains(str, "他") {
		gender = "男"
	}
	return gender
}

func findOrEmpty(sl []string, index int) string {
	if len(sl) > index + 1 {
		return sl[index]
	}else {
		return ""
	}
}

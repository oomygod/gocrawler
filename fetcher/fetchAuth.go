package fetcher

const dpUrl = "https://album.zhenai.com/u/1596207608"
var cookieStr string

func init() {
	//SetCookie("https://album.zhenai.com/u/1677398857")
	SetCookie(dpUrl)
}

func SetCookie(url string) {
	GetAllCookies(url)
	cookie := ReadCookie()
	var str string
	for _, ck := range cookie {
		str += ck.Name + "=" + ck.Value + ";"
	}
	cookieStr = str
}

func GetCookie() string {
	//cookieStr = "FSSBBIl1UgzbN7NO=5KZuvqBoshkeW0__6rJOnEBSh6.qQ39ZLUGLESlIj_53RoVN5PR4yJwS6O3YBJFTF0IcdibhkIiz_3vA5VsYZCG; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1658317539,1658391662,1658515266; _exid=DuAa7kYyuTuOTQlC+mxnW4p5IpfO9vbeUv61wq9p+tgzKJPtDIQKRqZRRFE/0YqRBmHwI1xu/kUeGt9edEllzw==; ec=cc44wIZ0-1658317514747-8abcd230f70161481524163; _efmdata=P3Mq1FIIEH/CPc7qDAwNXDmDpe9l6IsVVq9bqCbR6jY/BUr3OwelouHKoWNAPlvvCv0Oyrx+b1qcQABktqCLrU1cz7Dndp6deVqQR9rrlmI=; sid=X2j1OuFSfFN0yqRwogQP; FSSBBIl1UgzbN7NP=53z0cSbhwq.gqqqDchTh3_qLxzMPzTU7.yuEWAwz.eGWxA_Npx.WHRy7.vGKYtEzSonpxNgLcYjbTcgBV.Zu_INuLhz33FPYOblHDiLXpa0fq7p3ViDRdDH38gFC4eNRM9Z1b_zIgKJjxMDDEtHKeMYwljh0RZ6i36PSnm9qJKmoPXccb58xDA_tHWpt_5GyhKX15Hw_n06ZIeVowqfaBwPq1bHVpywGgKkXqBR9FFznbvpidFY6ZbirNK4TqXJgzL; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1658651329"
	return cookieStr
}

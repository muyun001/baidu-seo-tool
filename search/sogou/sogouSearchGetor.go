// 这个类是获取搜狗内容用的
package sogou

import (
	"fmt"
	"github.com/kevin-zx/http-util"
	"net/url"
	"strings"
)

var basicSogouUrl = "https://www.sogou.com/web?query=%s&ie=utf8&s_from=index&cid=&page=%d"

// 获取搜狗端html
func GetSogouSearchHtml(keyword string, pageNum int) (string, string, error) {
	searchUrl := combineSogouPcSearchUrl(keyword, pageNum)
	webCon, err := http_util.GetWebConFromUrlWithHeader(searchUrl,
		map[string]string{
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Accept-Encoding":           "gzip, deflate, br",
			"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8",
			"Connection":                "keep-alive",
			"Cookie":                    "IPLOC=CN3205; SUID=5AAAE0DD3118960A000000005D84492C; SUV=1568950571589337; pgv_pvi=3035993088; browerV=3; osV=1; CXID=CB3B16EB010022837BAF0A4A314D24E3; sw_uuid=3308802789; ssuid=3633323418; ABTEST=0|1571618290|v17; ad=nZllllllll2N$QIPlllllVLHlR9lllllTWj$Qkllll9lllllVylll5@@@@@@@@@@; sg_uuid=2848462209; SNUID=48B8F2CE121685B7E5040A2513FE170B; sst0=824; taspeed=taspeedexist; pgv_si=s2613157888; sct=160; ld=lZllllllll2NAxd3lllllVTj5b6lllllTWj$Qklllx9lllllxklll5@@@@@@@@@@; LSTMV=742%2C323; LCLKINT=283713",
			"Host":                      "www.sogou.com",
			"Referer":                   "https://www.sogou.com/web?query=%E5%BA%9F%E5%93%81%E5%9B%9E%E6%94%B6&_asf=www.sogou.com&w=01029901&cid=&s_from=result_up",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-Site":            "same-origin",
			"Sec-Fetch-User":            "?1",
			"Upgrade-Insecure-Requests": "1",
			"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
		},
	)
	if err != nil {
		return "", "", err
	}
	return webCon, searchUrl, nil
}

func combineSogouPcSearchUrl(wd string, pageNum int) string {
	wd = escapeSogouKeyword(wd)
	sogouSearchUrl := fmt.Sprintf(basicSogouUrl, wd, pageNum)
	sogouSearchUrl = strings.Replace(sogouSearchUrl, "&page=1", "", 1) // 第一页不需要写
	return sogouSearchUrl
}

// 对关键词转码
func escapeSogouKeyword(keyword string) string {
	keyword = url.QueryEscape(keyword)
	//keyword = strings.Replace(keyword, "+", "%20", -1)
	return keyword
}

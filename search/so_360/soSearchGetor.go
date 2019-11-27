// 这个类是获取360内容用的
package so_360

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/kevin-zx/http-util"
	"net/url"
	"strings"
)

var basicSoUrl = "https://www.so.com/index.php?q=%s&rc=srp&fr=none&pn=%d"

// 获取360端html
func GetSoSearchHtml(keyword string, pageNum int) (string, string, error) {
	searchUrl := combineSoPcSearchUrl(keyword, pageNum)
	webCon, err := http_util.GetWebConFromUrl(searchUrl)
	if err != nil {
		return "", "", err
	}

	return webCon, searchUrl, nil
}

func combineSoPcSearchUrl(wd string, pageNum int) string {
	wd = escapeSoKeyword(wd)
	soSearchUrl := fmt.Sprintf(basicSoUrl, wd, pageNum)
	soSearchUrl = strings.Replace(soSearchUrl, "&pn=1", "", 1) // 第一页不需要写
	return soSearchUrl
}

// 对关键词转码
func escapeSoKeyword(keyword string) string {
	keyword = url.QueryEscape(keyword)
	return keyword
}

func ConvertToByte(src string, srcCode string, targetCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(targetCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

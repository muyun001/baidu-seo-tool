package so_360

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type SoSearchInfo struct {
	Port              string
	SoMatchCount      int
	MainPageCount     int
	IsEscape          bool
	EscapeWord        string
	SoSearchResults   *[]SoSearchResult
	SoSearchAdResults *[]SoSearchResult
}

// 结果解析
func ParseSoSearchInfoFromHtml(html string, pageNum int) (ssi *SoSearchInfo, err error) {
	ssi = &SoSearchInfo{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}

	numStr := dom.Find("span.nums").Text()
	if numStr == "" {
		ssi.SoMatchCount = -1
	}

	numStr = strings.Replace(numStr, "找到相关结果约", "", -1)
	numStr = strings.Replace(numStr, "个", "", -1)
	numStr = strings.Replace(numStr, ",", "", -1)
	if numStr != "" {
		ssi.SoMatchCount, err = strconv.Atoi(numStr)
		if err != nil {
			return
		}
	}

	searchResults, err := ParseSoSearchResultHtml(html, pageNum)
	if err != nil {
		return
	}
	ssi.SoSearchResults = searchResults

	searchAdResults, err := ParseSoSearchAdResultHtml(html)
	if err != nil {
		return
	}
	ssi.SoSearchAdResults = searchAdResults

	return ssi, nil
}

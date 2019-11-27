package sogou_test

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search/sogou"
	"testing"
)

func TestGetSogouSearchHtml(t *testing.T) {
	keyword := "服装加工"
	pageNum := 1
	searchHTML, searchUrl, err := sogou.GetSogouSearchHtml(keyword, pageNum)
	ssi, err := sogou.ParseSogouSearchInfoFromHtml(searchHTML, pageNum)
	if err != nil {
		panic(err)
	}

	fmt.Println(ssi, searchUrl)
}

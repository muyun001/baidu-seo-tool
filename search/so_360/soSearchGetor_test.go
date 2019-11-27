package so_360_test

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search/so_360"
	"testing"
)

func TestGetSoSearchHtml(t *testing.T) {
	keyword := "废品回收"
	searchHTML,searchUrl, err := so_360.GetSoSearchHtml(keyword, 2)
	ssi, err := so_360.ParseSoSearchInfoFromHtml(searchHTML, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(ssi, searchUrl)
}

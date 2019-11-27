package so_360_test

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search/so_360"
	"io/ioutil"
	"strings"
	"testing"
)

const dataFileSoPc = "./test_html/so_pc.html"

func TestParseSoSearchResultHtml(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSoPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	searhHTML := strings.Replace(string(contents), "\n", "", 1)
	sogouPaserResult, err := so_360.ParseSoSearchResultHtml(searhHTML, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(sogouPaserResult)
}

func TestParseSoSearchAdResultHtml(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSoPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	searhHTML := strings.Replace(string(contents), "\n", "", 1)
	sogouPaserAdResult, err := so_360.ParseSoSearchAdResultHtml(searhHTML)
	if err != nil {
		panic(err)
	}

	fmt.Println(sogouPaserAdResult)
}

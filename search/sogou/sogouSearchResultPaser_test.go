package sogou_test

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/search/sogou"
	"io/ioutil"
	"strings"
	"testing"
)

const dataFileSogouPc = "./test_html/sogou_pc.html"

func TestParseSogouSearchResultHtml(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSogouPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	searhHTML := strings.Replace(string(contents), "\n", "", 1)
	sogouPaserResult, err := sogou.ParseSogouSearchResultHtml(searhHTML)
	if err != nil {
		panic(err)
	}

	fmt.Println(sogouPaserResult)
}

func TestParseSogouSearchAdResultHtml(t *testing.T) {
	contents, err := ioutil.ReadFile(dataFileSogouPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}

	searhHTML := strings.Replace(string(contents), "\n", "", 1)
	sogouPaserAdResult, err := sogou.ParseSogouSearchAdResultHtml(searhHTML)
	if err != nil {
		panic(err)
	}

	fmt.Println(sogouPaserAdResult)
}

package baidu

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

const dataFileBaiduPc = "./test_html/baidu_pc.html"

func TestParseBaiduPcSearchInfoFromHtml(t *testing.T) {
	//searhHTML, err := GetBaiduPCSearchHtml("苏州废气处理", 1)
	//if err != nil {
	//	panic(err)
	//}

	contents, err := ioutil.ReadFile(dataFileBaiduPc)
	if err != nil {
		t.Fatal("读取文件错误")
	}
	searhHTML := strings.Replace(string(contents), "\n", "", 1)

	bi, err := ParseBaiduPcSearchInfoFromHtml(searhHTML)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d, %d\n",len(*bi.SearchResults),len(*bi.SearchAdResults))
	for _, sr := range *bi.SearchAdResults {
		fmt.Printf("%v\n",sr)
	}
	for _, sr := range *bi.SearchResults {
		fmt.Println(sr.IsEnterpriseCertificate)
	}
}

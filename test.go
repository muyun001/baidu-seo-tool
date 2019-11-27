package main

import (
	"fmt"
	"github.com/kevin-zx/baidu-seo-tool/utils"
	"time"
)

var (
	nowCount  int
	sumCount  int
	startTime = time.Now()
)

func main() {
	initDomain := "www.hfbennuo.cn/articl..."
	domain, err := utils.Domain(initDomain)
	fmt.Println(domain, err)
}

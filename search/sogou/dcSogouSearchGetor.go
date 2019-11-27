package sogou

import (
	"encoding/json"
	"github.com/panwenbin/ghttpclient"
)

type BuilderRequest struct {
	SearchWord  string `json:"search_word" bson:"search_word"`
	Page        int    `json:"page" bson:"page"`
	Capture     bool   `json:"capture" bson:"capture"`
	SearchCycle int    `json:"search_cycle"`
	Priority    string `json:"priority"`
}

type BuilderRequestResponse struct {
	UniqueKey string `json:"unique_key"`
	Request   struct {
		Url       string `json:"url"`
		UserAgent string `json:"user_agent"`
		Cookie    string `json:"cookie"`
		Body      string `json:"body"`
	} `json:"request"`
	Config struct {
		District       string   `json:"district"`
		ResponseTypes  []string `json:"response_types" bson:"response_types"`
		FollowRedirect bool     `json:"follow_redirect"`
		Priority       string   `json:"priority"`
	} `json:"config"`
	Status int `json:"status"`
}

type DcSetTaskRequest struct {
	UserID  string `json:"user_id"`
	Headers string `json:"headers"`
	Config  string `json:"config"`
	Urls    string `json:"urls"`
}

var buildRequestApi = "127.0.0.1:8081/request-builder/360-pc"

// 通过下载中心发送请求
func DcSogouSearchHtml(keyword string, pageNum int) (string, string, error) {
	// 构建下载中心请求
	requestBuilderRequest := BuilderRequest{
		SearchWord:  keyword,
		Page:        pageNum,
		Capture:     false,
		SearchCycle: 1,
		Priority:    "normal",
	}
	jsonBytes, err := json.Marshal(requestBuilderRequest)
	var dcRequest BuilderRequestResponse
	err = ghttpclient.PostJson(buildRequestApi, jsonBytes, nil).ReadJsonClose(&dcRequest)
	if err != nil {
		return "", "", err
	}

	//// 发送请求到下载中心
	//dcSetTaskRequest := &DcSetTaskRequest{
	//	UserID:  "25",
	//	Headers: fmt.Sprintf(`{"User-Agent": "%s", "Cookie": "%s"}`, dcRequest.Request.UserAgent, dcRequest.Request.Cookie),
	//	Config:  fmt.Sprintf(`{"redirect": 0, "priority": 1}`),
	//	Urls:    fmt.Sprintf(`[{"url": "%s", "type": 1, "unique_key": "%s"}]`, dcRequest.UniqueKey),
	//}



	return "", "", nil
}

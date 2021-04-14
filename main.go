package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Grafanawebhook struct {
	dashboardId string `json:"dashboardId"`
	evalMatches string `json:"evalMatches"`
	ImageUrl    string `json:"imageUrl"`
	Message     string `json:"message"`
	orgId       string `json:"orgId"`
	panelId     string `json:"panelId"`
	ruleId      string `json:"ruleId"`
	ruleName    string `json:"ruleName"`
	RuleUrl     string `json:"ruleUrl"`
	state       string `json:"state"`
	tags        string `json:"tags"`
	Title       string `json:"title"`
}

func Wechat(k, a, b, c, d string) {
	//转发key到 企业微信接口
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + k
	msgStr := fmt.Sprintf(`
		{
			"msgtype": "news",
			"news": {
			  "articles": [
				{
				  "title": "%s",
				  "description": "%s",
				  "url": "%s",
				  "picurl": "%s"
				}
			  ]
			}
		  }
		`, a, b, c, d)
	jsonStr := []byte(msgStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending to WeChat Work API")
		return
	}
	defer resp.Body.Close()
}

func main() {
	r := gin.New()
	// use recovery
	r.Use(gin.Recovery())
	// default logs 日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s %s %s %s %d %s %s %s \"\n",
			param.TimeStamp.Format(time.RFC1123),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	r.POST("/:wechatkey", func(c *gin.Context) {
		json := new(Grafanawebhook)
		if err := c.ShouldBindJSON(json); err != nil {
			fmt.Printf("message: %s\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		key := c.Param("wechatkey")
		// debug
		//fmt.Printf("KEY: %s\t Tile: %s\t msg: %s\t RuleUrl: %s\t ImageUrl: %s\n", key,json.Title,json.Message,json.RuleUrl,json.ImageUrl)
		Wechat(key, json.Title, json.Message, json.RuleUrl, json.ImageUrl)
	})
	r.Run(":8081")
}

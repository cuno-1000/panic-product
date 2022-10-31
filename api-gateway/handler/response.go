package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code      int         `json:"code"`
	Message   interface{} `json:"message"`
	Time      string      `json:"time"`
	TimeStamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		nil,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Unix(),
		data,
	})
}

func Fail(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, Response{
		-1,
		msg,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Unix(),
		nil,
	})
}

func BznFail(c *gin.Context, errorCode int) {
	c.JSON(http.StatusOK, Response{
		errorCode,
		CustomDecode(errorCode),
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Unix(),
		nil,
	})
}

var m []string

func init() {
	m = append(m,
		"抢购成功",
		"用户逾期还款次数超过活动规则要求",
		"用户被列入当前严重违法失信被执行人名单，未执行完毕",
		"用户小于18岁",
		"用户工作状态为”无业/失业“",
		"限购",
		"活动未开始",
		"余额不足",
		"抢光了",
	)
}

func CustomDecode(num int) (messages []string) {
	for i := 0; num > 0; num >>= 1 {
		if num&1 == 1 {
			messages = append(messages, m[i])
		}
		i++
	}
	return messages
}

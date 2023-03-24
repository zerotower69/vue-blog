package res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// ListResponse the common return value
type ListResponse[T any] struct {
	List  []T `json:"list"`
	Count any `json:"count""`
}

const (
	ERROR   = 400
	SUCCESS = 200
)

// Result function get the response data with the code or message
func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OK(c *gin.Context) {
	Result(SUCCESS, map[string]any{}, "操作成功", c)
}

func OkWihMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]any{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data any, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func OkWithList[T any](List []T, count any, c *gin.Context) {
	if len(List) == 0 {
		List = []T{}
	}
	Result(SUCCESS, ListResponse[T]{
		List:  List,
		Count: count,
	}, "成功", c)
}

func Fail(data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]any{}, message, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
		return
	}
	//if not match any code defined, get unknown error
	Result(ERROR, map[string]any{}, "未知错误", c)
}

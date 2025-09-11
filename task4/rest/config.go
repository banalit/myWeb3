package task4

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func SuccessResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{Code: 0, Data: data, Message: msg})
}

func ErrorResponse(c *gin.Context, code int, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{Code: code, Message: msg, Data: data})
}

func GetUserId(c *gin.Context) (uint, error) {
	var userID, hasValue = c.Get("UserID")
	if !hasValue {
		return 0, errors.New("未授权")
	}

	// 确保UserID是string类型
	userIDInt, ok := userID.(uint)
	if !ok {
		return 0, errors.New("用户未授权")
	}
	return userIDInt, nil
}

func init() {

}

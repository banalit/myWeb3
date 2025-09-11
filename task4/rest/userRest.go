package task4

import (
	"log"

	task4Db "task4/db"
	task4Middle "task4/middleware"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

var router *gin.Engine

func init() {

}

func Register(c *gin.Context) {
	var user task4Db.User
	// body, _ := io.ReadAll(c.Request.Body)
	// log.Printf("原始表单数据: %s", body)
	// 重新将数据写回请求体，避免后续绑定失败
	if err := c.ShouldBindWith(&user, binding.Form); err != nil {
		ErrorResponse(c, 500, nil, err.Error())
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ErrorResponse(c, 500, nil, err.Error())
		return
	}
	user.Password = string(hashedPassword)
	if err := task4Db.CreateUser(&user); err != nil {
		log.Println("发生错误：", err)
		ErrorResponse(c, 500, nil, "Failed to create user")
		return
	}

	SuccessResponse(c, user, "user created successfully")
}

func Login(c *gin.Context) {
	var userInput task4Db.User
	if err := c.ShouldBindWith(&userInput, binding.Form); err != nil {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	if userDb, err := task4Db.GetUser(userInput.UserName); err != nil {
		ErrorResponse(c, 400, nil, "user login failed")
		return
	} else {
		if !validatePwd(userDb.Password, userInput.Password) {
			ErrorResponse(c, 400, nil, "user login failed")
			return
		}
		userInput = userDb
	}
	if token, err := task4Middle.GenerateJwtToken(userInput.ID, userInput.UserName); err != nil {
		ErrorResponse(c, 400, nil, "生成令牌失败")
		return
	} else {
		SuccessResponse(c, gin.H{
			"id":       userInput.ID,
			"username": userInput.UserName,
			"email":    userInput.Email,
			"token":    token,
		}, "登录成功")
	}

}

func validatePwd(dbPwd string, inputPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(inputPwd)); err != nil {
		return false
	}
	return true
}

// 退出登录处理函数
func Logout(c *gin.Context) {
	// 从上下文获取当前token
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 8 {
		c.JSON(401, gin.H{"error": "无法获取token"})
		c.Abort()
		return
	}
	tokenString := authHeader[7:]
	if err := task4Middle.AddBlacklistToken(tokenString); err != nil {
		ErrorResponse(c, 400, nil, "fail to delete token")
		return
	}
	SuccessResponse(c, nil, "退出登录成功")
}

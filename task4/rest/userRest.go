package task4

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	task4Db "github.com/luke/web3Learn/task4/db"
	task4Middle "github.com/luke/web3Learn/task4/middleware"
	"golang.org/x/crypto/bcrypt"
)

var router *gin.Engine

func init() {

}

func AddRestApi(r *gin.Engine) {
	apiPublic := r.Group("/api")
	{
		apiPublic.POST("/register", register)
		apiPublic.POST("/login", login)
	}
	apiJwt := r.Group("/api")
	apiJwt.Use(task4Middle.JwtAuthMiddleware)
	{
		apiJwt.POST("/logout", logout)

		apiJwt.POST("/post", createPost)
		apiJwt.DELETE("/post", deletePost)
		apiJwt.PATCH("/post", patchPost)
		apiJwt.GET("/post", listPost)
		apiJwt.POST("/comment", createComment)
		apiJwt.DELETE("/comment", deleteComment)
		apiJwt.PATCH("/comment", patchComment)
		apiJwt.GET("/comment", listComment)
	}

}

func register(c *gin.Context) {
	var user task4Db.User
	if err := c.ShouldBindWith(&user, binding.Form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)
	if err := task4Db.CreateUser(&user); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

func login(c *gin.Context) {
	var userInput task4Db.User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if userDb, err := task4Db.GetUser(userInput.UserName); err != nil {
		c.JSON(200, gin.H{"error": "user login failed"})
		return
	} else {
		if !validatePwd(userDb.Password, userInput.Password) {
			c.JSON(200, gin.H{"error": "user login failed"})
			return
		}
		userInput = userDb
	}
	if token, err := task4Middle.generateJwtToken(userInput.UserID, userInput.UserName); err != nil {
		c.JSON(500, gin.H{"error": "生成令牌失败"})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "登录成功",
			"token":   token,
			"user": gin.H{
				"id":       userInput.UserID,
				"username": userInput.Username,
				"email":    userInput.Email,
			},
		})
	}

}

func validatePwd(dbPwd string, inputPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(inputPwd)); err != nil {
		return false
	}
	return true
}

// 退出登录处理函数
func logout(c *gin.Context) {
	// 从上下文获取当前token
	tokenString, exists := c.Get("tokenString")
	if !exists {
		c.JSON(400, gin.H{"error": "无法获取token"})
		return
	}
	if err := task4Middle.addBlacklistToken(tokenString); err != nil {
		c.JSON(200, gin.H{
			"error": "fail to delete token",
		})
		return
	}

	c.JSON(200, gin.H{"message": "退出登录成功"})
}

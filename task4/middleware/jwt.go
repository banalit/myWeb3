package task4

import (
	"errors"
	"os"
	"time"

	task4Db "../db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

type JwtClaims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func init() {
	jwtSecretStr := os.Getenv("JWT_SECRET")
	if jwtSecretStr == "" {
		jwtSecret = []byte("luke/web3Learning")
	} else {
		jwtSecret = []byte(jwtSecretStr)
	}
}

func generateJwtToken(userId uint64, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	nowNumericDate := jwt.NewNumericDate(time.Now())
	claim := &JwtClaims{
		UserID:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  nowNumericDate,
			NotBefore: nowNumericDate,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "git-jwt-auth",
			Subject:   "user-token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	tokenString, err := token.SignedString(token)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func JwtAuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || len(authHeader) < 8 {
		c.JSON(401, gin.H{"error": "authorization is invalide"})
		c.Abort()
		return
	}
	tokenString := authHeader[7:]

	// 检查token是否在黑名单中
	if task4Db.isTokenBlacklisted(tokenString) {
		c.JSON(401, gin.H{"error": "token已失效，请重新登录"})
		c.Abort()
		return
	}

	claim := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "token invalide"})
		c.Abort()
		return
	}

	c.Set("UserID", claim.UserID)
	c.Set("Username", claim.Username)
	c.Next()

}

func addBlacklistToken(tokenString interface{}) error {
	// 解析token以获取过期时间
	claims := JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString.(string), claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return errors.New("无效的token")
	}

	// 将token加入黑名单
	blacklist := task4Db.TokenBlacklist{
		Token:     tokenString.(string),
		ExpiresAt: claims.ExpiresAt.Time,
	}

	if err := task4Db.CreateTokenBlackList(&blacklist); err != nil {
		return errors.New("退出登录失败")
	}
	return nil
}

package task4

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"time"

	task4Db "task4/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey *ecdsa.PrivateKey
var publicKey *ecdsa.PublicKey

type JwtClaims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
	jwt.RegisteredClaims
}

func init() {
	// 定义密钥文件路径
	privateKeyPath := "ec_private.pem"
	publicKeyPath := "ec_public.pem"

	// 检查密钥文件是否存在，不存在则生成
	if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
		if err := generateAndSaveKeys(privateKeyPath, publicKeyPath); err != nil {
			log.Fatalf("生成密钥失败: %v", err)
		}
		log.Println("已生成新的密钥对")
	}

	// 加载密钥对
	pvk, err := loadPrivateKey(privateKeyPath)
	if err != nil {
		log.Fatalf("加载私钥失败: %v", err)
	}
	privateKey = pvk
	pubKey, err := loadPublicKey(publicKeyPath)
	if err != nil {
		log.Fatalf("加载公钥失败: %v", err)
	}
	publicKey = pubKey
}

// 从文件加载私钥
func loadPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	// 读取文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 解析PEM块
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, errors.New("invalid private key PEM file")
	}

	// 解析EC私钥
	return x509.ParseECPrivateKey(block.Bytes)
}

// 从文件加载公钥
func loadPublicKey(path string) (*ecdsa.PublicKey, error) {
	// 读取文件内容
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 解析PEM块
	block, _ := pem.Decode(data)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid public key PEM file")
	}

	// 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言为ECDSA公钥
	ecdsaPub, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("not an ECDSA public key")
	}

	return ecdsaPub, nil
}

// 生成ECDSA密钥对并保存到文件
func generateAndSaveKeys(privateKeyPath, publicKeyPath string) error {
	// 生成P256曲线的ECDSA私钥
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}

	// 保存私钥
	if err := savePrivateKey(privateKeyPath, privateKey); err != nil {
		return err
	}

	// 保存公钥
	return savePublicKey(publicKeyPath, &privateKey.PublicKey)
}

// 保存私钥到PEM文件
func savePrivateKey(path string, key *ecdsa.PrivateKey) error {
	// 将私钥序列化为DER格式
	derBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	// 创建PEM块
	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derBytes,
	}

	// 写入文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, pemBlock)
}

// 保存公钥到PEM文件
func savePublicKey(path string, key *ecdsa.PublicKey) error {
	// 将公钥序列化为DER格式
	derBytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return err
	}

	// 创建PEM块
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}

	// 写入文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return pem.Encode(file, pemBlock)
}

func GenerateJwtToken(userId uint, username *string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	nowNumericDate := jwt.NewNumericDate(time.Now())
	claim := &JwtClaims{
		UserID:   userId,
		Username: *username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  nowNumericDate,
			NotBefore: nowNumericDate,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "git-jwt-auth",
			Subject:   "user-token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	tokenString, err := token.SignedString(privateKey)
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
	if task4Db.IsTokenBlacklisted(tokenString) {
		c.JSON(401, gin.H{"error": "token已失效，请重新登录"})
		c.Abort()
		return
	}

	claim := &JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
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

func AddBlacklistToken(tokenString interface{}) error {
	// 解析token以获取过期时间
	claims := JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString.(string), &claims, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
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

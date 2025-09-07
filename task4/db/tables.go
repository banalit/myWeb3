package task4

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         string    `gorm:"not null"`
	Content       string    `gorm:"not null"`
	UserID        uint64    `gorm:"not null"`
	User          User      `gorm:"references:ID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentCount  uint64    `gorm:"default:0"`
	CommentStatus string    `gorm:"default:no comment"`
}

type Comment struct {
	gorm.Model
	UserID  uint64
	User    User   `gorm:"references:ID"`
	Content string `gorm:"not null"`
	PostID  uint64
	Post    Post `gorm:"references:ID"`
}

type User struct {
	gorm.Model
	UserName  string    `gorm:"unique;not null" json:"username" binding:"required"`
	Password  string    `gorm:"not null" json:"-"`
	PostCount uint16    `gorm:"default:0"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`

	Email string `gorm:"unique;not null"`
}

// TokenBlacklist 存储已失效的token
type TokenBlacklist struct {
	gorm.Model
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("after create post:", p, p.UserID, p.User.ID)
	err = tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_count", gorm.Expr("post_count + 1")).Error
	return err
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("after create comment:", c)
	err = tx.Model(&Post{}).Where("id=?", c.PostID).Updates(map[string]interface{}{
		"comment_count":  gorm.Expr("comment_count + 1"),
		"comment_status": "have comment",
	}).Error
	return err
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete comment:", c)
	if c.ID == 0 {
		return
	}
	err = tx.Model(&Post{}).Where("id=?", c.PostID).Updates(map[string]interface{}{
		"comment_count":  gorm.Expr("comment_count-1"),
		"comment_status": gorm.Expr("case when comment_count=0 then 'no comment' else 'have comment' end"),
	}).Error
	return err
}

func init() {
	db := getGormSqlliteDb()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{}, &TokenBlacklist{})
}

func GetUser(userName string) (User, error) {
	db := getGormSqlliteDb()
	var user User
	err := db.Where("user_name=? ", userName).First(&user).Error
	return user, err
}

func CreateUser(user *User) error {
	db := getGormSqlliteDb()
	return db.Create(user).Error
}

// 检查token是否在黑名单中
func IsTokenBlacklisted(tokenString string) bool {
	db := getGormSqlliteDb()
	var blacklist TokenBlacklist
	result := db.Where("token = ? AND expires_at > ?", tokenString, time.Now()).First(&blacklist)
	return result.Error == nil
}

func CreateTokenBlackList(tokenBlackList *TokenBlacklist) error {
	db := getGormSqlliteDb()
	return db.Create(tokenBlackList).Error
}

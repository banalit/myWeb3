package task3

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         string
	Content       string
	UserID        uint64
	User          User      `gorm:"references:ID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentCount  uint64    `gorm:"default:0"`
	CommentStatus string    `gorm:"default:no comment"`
}

type Comment struct {
	gorm.Model
	UserID  uint64
	User    User `gorm:"references:ID"`
	Content string
	PostID  uint64
	Post    Post `gorm:"references:ID"`
}

type User struct {
	gorm.Model
	Name      string
	PostCount uint16    `gorm:"default:0"`
	Posts     []Post    `gorm:"foreignKey:UserID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
	UserName  string
	Password  string
	Email     string
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

package task4

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title         *string   `gorm:"not null" validate:"required"`
	Content       *string   `gorm:"not null" validate:"required"`
	UserID        uint      `gorm:"not null" validate:"required"`
	User          *User     `gorm:"references:ID" json:"user,omitempty" validate:"-"`
	Comments      []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty" validate:"-"`
	CommentCount  uint64    `gorm:"default:0"`
	CommentStatus *string   `gorm:"default:no comment"`
}

type Comment struct {
	gorm.Model
	UserID  uint
	User    *User   `gorm:"references:ID" json:"user,omitempty" validate:"-"`
	Content *string `gorm:"not null"`
	PostID  uint
	Post    *Post `gorm:"references:ID" json:"post,omitempty" validate:"-"`
}

type User struct {
	gorm.Model
	UserName  *string   `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	PostCount uint16    `gorm:"default:0"`
	Posts     []Post    `gorm:"foreignKey:UserID" validate:"-" json:"posts,omitempty"`
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty" validate:"-"`
	Email     *string   `gorm:"unique;not null" form:"email"`
}

// TokenBlacklist 存储已失效的token
type TokenBlacklist struct {
	gorm.Model
	Token     string    `gorm:"uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("after create post:", p, p.UserID, p.User)
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

func GetUser(userName *string) (User, error) {
	db := getGormSqlliteDb()
	var user User
	err := db.Where("user_name=? ", userName).First(&user).Error
	return user, err
}

func Get(obj interface{}) (interface{}, error) {
	db := getGormSqlliteDb()
	result := db.Model(&obj).Where(obj).Take(&obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
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

func UpdatePost(post *Post) error {
	db := getGormSqlliteDb()
	return db.Model(&Post{}).Where("id=?", post.ID).Updates(post).Error
}

func UpdateComment(comment *Comment) error {
	db := getGormSqlliteDb()
	return db.Debug().Model(&Comment{}).Where("id=?", comment.ID).Updates(comment).Error
}

func CreateComment(comment *Comment) error {
	db := getGormSqlliteDb()
	return db.Create(comment).Error
}

func CreatePost(post *Post) error {
	db := getGormSqlliteDb()
	return db.Create(post).Error
}

func Create(obj interface{}) error {
	db := getGormSqlliteDb()
	return db.Create(obj).Error
}

func Delete(obj interface{}) error {
	db := getGormSqlliteDb()
	return db.Unscoped().Debug().Where(obj).Delete(&obj).Error
}

func ListComment(filter Comment, c *gin.Context) ([]Comment, int64, int, int, error) {
	db := getGormSqlliteDb()
	var comments []Comment
	var query = db.Model(&Comment{}).Where(filter)
	var total int64
	page, pageSize := QueryPageHelper(c)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, page, pageSize, err
	}
	if err := query.Scopes(Paginate(page, pageSize)).Find(&comments).Error; err != nil {
		return nil, total, page, pageSize, err
	}
	return comments, total, page, pageSize, nil
}

func ListPost(filter Post, c *gin.Context) ([]Post, int64, int, int, error) {
	db := getGormSqlliteDb()
	var posts []Post
	var query = db.Model(&Post{})
	if filter.Title != nil {
		query.Where("title like ?", "%"+*filter.Title+"%")
	}
	if filter.Content != nil {
		query.Where("content like ?", "%"+*filter.Content+"%")
	}
	query.Where("User_ID = ?", filter.UserID)
	var total int64
	page, pageSize := QueryPageHelper(c)
	if err := query.Debug().Count(&total).Error; err != nil {
		return nil, 0, page, pageSize, err
	}
	if err := query.Debug().Scopes(Paginate(page, pageSize)).Find(&posts).Error; err != nil {
		return nil, total, page, pageSize, err
	}
	return posts, total, page, pageSize, nil
}

// 分页
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func QueryPageHelper(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return page, pageSize
}

//   db.Scopes(Paginate(r)).Find(&users)

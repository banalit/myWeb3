package task3

import (
	"fmt"
)

func PrepareBlogData() {
	db := getGormSqlliteDb()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})

	var users []User = []User{
		{Name: "peter", PostCount: 0},
		{Name: "lucy", PostCount: 0},
		{Name: "bana", PostCount: 0},
		{Name: "apple", PostCount: 0},
	}
	fmt.Println("to create user")
	db.Create(&users)

	fmt.Println("to create post")
	var posts []Post = []Post{
		{Title: "title1", Content: "abc", UserID: uint64(users[0].ID)},
		{Title: "title2", Content: "123", UserID: uint64(users[0].ID)},
		{Title: "title3", Content: "nnnn", UserID: uint64(users[1].ID)},
		{Title: "title4", Content: "ew wff", UserID: uint64(users[2].ID)},
		{Title: "title5", Content: "gafe fds", UserID: uint64(users[0].ID)},
	}
	db.Create(&posts)

	fmt.Println("to create comment")

	var comments []Comment = []Comment{
		{Content: "haha", UserID: uint64(posts[0].UserID), PostID: uint64(posts[0].ID)},
		{Content: "abc", UserID: uint64(posts[0].UserID), PostID: uint64(posts[0].ID)},
		{Content: "efe", UserID: uint64(posts[0].UserID), PostID: uint64(posts[0].ID)},
		{Content: "123", UserID: uint64(posts[1].UserID), PostID: uint64(posts[1].ID)},
		{Content: "fa fea 2", UserID: uint64(posts[2].UserID), PostID: uint64(posts[2].ID)},
		{Content: "1357", UserID: uint64(posts[3].UserID), PostID: uint64(posts[3].ID)},
		{Content: "jkk", UserID: uint64(posts[2].UserID), PostID: uint64(posts[2].ID)},
		{Content: "iou", UserID: uint64(posts[0].UserID), PostID: uint64(posts[0].ID)},
	}
	db.Create(&comments)
}

func BlogTest() {
	db := getGormSqlliteDb()
	c := &Comment{}
	db.First(c, 1)
	err := db.Unscoped().Debug().Delete(c).Error
	fmt.Println("err:", err)

	fmt.Println("save post")
	var p = Post{Title: "title1", Content: "abc", UserID: 3}
	db.Create(&p)

	var user = User{Name: "peter"}
	result := db.Preload("Posts").Preload("Comments").Preload("Comments.Post").First(&user)
	if result.Error != nil {
		fmt.Println("blog test error:", result.Error)
		return
	}
	fmt.Println("blog user:", user.ID, user.Name, user.PostCount)
	for _, post := range user.Posts {
		fmt.Println("post:", post.Title, post.Content)
	}
	for _, comment := range user.Comments {
		fmt.Println("comment:", comment.Content, comment.Post.Title)
	}

}

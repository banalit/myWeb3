package task4

import (
	task4Db "task4/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}

	var post task4Db.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	post.UserID = userId
	if err := task4Db.Create(&post); err != nil {
		ErrorResponse(c, 400, nil, "create err:"+err.Error())
		return
	}
	SuccessResponse(c, post, "成功添加")
}

func DeletePost(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	var req struct {
		ID uint `json:"id" binding:"required"` // 假设需要ID来指定删除哪个帖子
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, 400, nil, "请求参数错误: "+err.Error())
		return
	}
	post := task4Db.Post{
		Model:  gorm.Model{ID: req.ID},
		UserID: userId,
	}
	if err := task4Db.Delete(&post); err != nil {
		ErrorResponse(c, 500, nil, "删除帖子失败: "+err.Error())
		return
	}
	SuccessResponse(c, nil, "成功删除")
}

func PatchPost(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	var post = task4Db.Post{}
	if err := c.ShouldBindJSON(&post); err != nil {
		ErrorResponse(c, 400, nil, "data err:"+err.Error())
		return
	}
	post.UserID = userId
	task4Db.UpdatePost(&post)
	SuccessResponse(c, post, "成功更新")
}

func ListPost(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	var filter task4Db.Post
	if err := c.ShouldBindJSON(&filter); err != nil {
		ErrorResponse(c, 400, nil, "data err:"+err.Error())
		return
	}
	filter.UserID = userId
	if posts, total, page, pageSize, err := task4Db.ListPost(filter, c); err != nil {
		ErrorResponse(c, 400, nil, err.Error())
	} else {
		SuccessResponse(c, gin.H{
			"data":     posts,
			"page":     page,
			"pageSize": pageSize,
			"total":    total,
		}, "success")
	}

}

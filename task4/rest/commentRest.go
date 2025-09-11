package task4

import (
	task4Db "task4/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteComment(c *gin.Context) {
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
	comment := task4Db.Comment{
		Model:  gorm.Model{ID: req.ID},
		UserID: userId,
	}
	if err := task4Db.Delete(&comment); err != nil {
		ErrorResponse(c, 500, nil, "删除评论失败: "+err.Error())
		return
	}
	SuccessResponse(c, nil, "成功删除")
}

func PatchComment(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	comment := task4Db.Comment{}
	if err := c.ShouldBindJSON(&comment); err != nil {
		ErrorResponse(c, 400, nil, "data err:"+err.Error())
		return
	}
	comment.UserID = userId
	if err := task4Db.UpdateComment(&comment); err != nil {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	task4Db.Get(&comment)
	SuccessResponse(c, comment, "成功更新")
}

func ListComment(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	var filter task4Db.Comment
	if err := c.ShouldBindJSON(&filter); err != nil {
		ErrorResponse(c, 400, nil, "data err:"+err.Error())
		return
	}
	filter.UserID = userId
	if comments, total, page, pageSize, err := task4Db.ListComment(filter, c); err != nil {
		ErrorResponse(c, 400, nil, err.Error())
	} else {
		SuccessResponse(c, gin.H{
			"data":     comments,
			"page":     page,
			"pageSize": pageSize,
			"total":    total,
		}, "success")
	}

}

func CreateComment(c *gin.Context) {
	var userId uint
	if uid, err := GetUserId(c); err == nil {
		userId = uid
	} else {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}

	var comment task4Db.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		ErrorResponse(c, 400, nil, err.Error())
		return
	}
	comment.UserID = userId
	if err := task4Db.Create(&comment); err != nil {
		ErrorResponse(c, 400, nil, "create err:"+err.Error())
		return
	}
	SuccessResponse(c, comment, "成功添加")
}

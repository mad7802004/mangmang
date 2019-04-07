package app

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/setting"
)

// 获取页码和页数
func GetPageSize(c *gin.Context) (int, int) {
	var page, size int

	page = com.StrTo(c.Query("page")).MustInt()
	size = com.StrTo(c.Query("size")).MustInt()

	if page <= 0 {
		page = setting.AppSetting.Page
	}
	if size <= 0 {
		size = setting.AppSetting.PageSize
	}

	return page, size
}

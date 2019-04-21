package website

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Home(c *gin.Context) {
	userId := com.ToStr(c.MustGet("user_id"))
	userInfo, _ := models.FindUserIdInfo(userId)
	c.HTML(http.StatusOK, "home.html", userInfo)
}

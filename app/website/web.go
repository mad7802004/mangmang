package website

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func Project(c *gin.Context) {
	userId := com.ToStr(c.MustGet("user_id"))
	data, _, _ := models.FindUserProject(userId, 1, 99)
	c.HTML(http.StatusOK, "project.html",gin.H{"data":data} )
}

func ProjectInfo(c *gin.Context) {
	key := c.Param("key")
	query, err := models.FindProject(key)
	if err != nil {
		c.HTML(http.StatusForbidden, "403.html", gin.H{})
	}

	c.HTML(http.StatusOK, "project_detail.html", query)
}

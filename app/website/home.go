package website

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "MangMang",
	})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "MangMang",
	})
}

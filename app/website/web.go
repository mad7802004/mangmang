package website

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{})
}

func Project(c *gin.Context)  {
	c.HTML(http.StatusOK, "project.html", gin.H{})
}
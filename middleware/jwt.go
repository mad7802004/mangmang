package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/utils"

	"net/http"
)

func JWT() gin.HandlerFunc {

	return func(c *gin.Context) {
		lang := c.GetHeader("lang")
		appG := app.Gin{C: c, M: e.NewMsg(lang)}
		token := c.GetHeader("Authorization")

		if token == "" {
			appG.Response(http.StatusOK, e.InvalidAuthorization, nil)
			c.Abort()
			return
		}

		userId, ok := utils.Identify(token)
		if !ok {
			appG.Response(http.StatusOK, e.InvalidAuthorization, nil)
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Next()
	}
}

func CheckWebLogin() gin.HandlerFunc {

	return func(c *gin.Context) {
		userId, err := c.Cookie("user_id")
		if err != nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}
		c.Set("user_id", userId)
		c.Next()

	}

}

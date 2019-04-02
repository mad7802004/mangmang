package service_user

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"net/http"
)

// 测试
func Test(c *gin.Context) {
	appG := app.New(c)
	sCode := com.ToStr(c.Param("k"))
	k := map[string]interface{}{
		"qin":  "1",
		"qin1": sCode,
	}
	d, _ := k["nickname"].(string)
	fmt.Print(d)
	appG.Response(http.StatusOK, e.SUCCESS, d)
	return
}

func UserLogin(c *gin.Context) {

}

package service_user

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"github.com/mangmang/pkg/gredis"
	"github.com/mangmang/pkg/utils"
	"net/http"
	"time"
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

// 获取手机验证码
func GetVerificationCode(c *gin.Context) {
	appG := app.New(c)
	phone := c.Query("phone")
	if !utils.CheckPhone(phone) {
		appG.Response(http.StatusOK, e.MobileNumberError, nil)
		return
	}

	expireTime, err := gredis.Hget(phone, "expire_time")
	if err == nil {
		nowTime := time.Now()
		expireTime, _ := time.Parse("2006-01-02 15:04:05", string(expireTime))
		if nowTime.Unix()-expireTime.Unix() < 60 {
			appG.Response(http.StatusOK, e.FrequentOperation, nil)
			return
		}
	}
	//code := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	code := "123456"
	err = gredis.Hset(phone, "code", code)
	if err != nil {
		appG.Response(http.StatusOK, e.FAIL, nil)
		return
	}
	err = gredis.Hset(phone, "expire_time", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		appG.Response(http.StatusOK, e.FAIL, nil)
		return
	}
	err = gredis.Expire(phone, 60*3)
	if err != nil {
		appG.Response(http.StatusOK, e.FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 用户手机号注册
func PhoneRegister(c *gin.Context) {
	var obj struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		PassWord1 string `json:"pass_word_1"`
		PassWord2 string `json:"pass_word_2"`
		Code      string `json:"code"`
	}
	appG := app.New(c)

	//参数解析失败
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 两次密码不一致
	if obj.PassWord1 != obj.PassWord2 {
		appG.Response(http.StatusOK, e.InconsistentPassword, nil)
		return
	}

	// 验证码错误
	if !utils.CheckPhoneCode(obj.Phone, obj.Code, false) {
		appG.Response(http.StatusOK, e.VerificationCodeError, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 用户密码登陆
func UserLoginAPW(c *gin.Context) {
	var obj struct {
		Phone    string `json:"phone"`
		PassWord string `json:"pass_word"`
	}
	appG := app.New(c)

	//参数解析失败
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}

// 用户验证码登陆
func UserLoginAV(c *gin.Context) {
	var obj struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	appG := app.New(c)

	//参数解析失败
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}

// 用户修改密码
func ChangePW(c *gin.Context) {
	var obj struct {
		Phone       string `json:"phone"`
		Code        string `json:"code"`
		OldPassWord string `json:"old_pass_word"`
		PassWord1   string `json:"pass_word_1"`
		PassWord2   string `json:"pass_word_2"`
	}
	appG := app.New(c)

	//参数解析失败
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}

// 用户修改个人信息
func ChangeInfo(c *gin.Context) {
	var obj struct {
	}
	appG := app.New(c)

	//参数解析失败
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return

}

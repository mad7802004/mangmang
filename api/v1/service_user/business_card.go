package service_user

import (
	"github.com/gin-gonic/gin"
	"github.com/mangmang/models"
	"github.com/mangmang/pkg/app"
	"github.com/mangmang/pkg/e"
	"net/http"
)

// 获取个人名片
func GetBusinessCard(c *gin.Context) {
	appG := app.New(c)
	key := c.Param("key")
	userId := c.Query("user_id")
	// 判断是否查询一个名片
	if key != "" {
		query, err := models.FindBusinessCard(key)
		if err != nil {
			appG.Response(http.StatusOK, e.NoResourcesFound, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, query)
		return
	}
	// 用户未填
	if userId == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}
	// 获取分页
	page, size := app.GetPageSize(c)
	data, total, err := models.FindUserBusinessCard(userId, page, size)
	// 未找到数据
	if err != nil || len(data) == 0 {
		appG.AddField("total", total)
		appG.Response(http.StatusOK, e.NoResourcesFound, nil)
		return
	}

	appG.AddField("total", total)
	appG.Response(http.StatusOK, e.SUCCESS, data)
	return
}

// 新建个人名片
func CreateBusinessCard(c *gin.Context) {
	var obj struct {
		UserId   string `json:"user_id"binding:"uuid4"`
		Name     string `json:"name"binding:"max=10"`
		Company  string `json:"company"binding:"max=255"`
		Position string `json:"position"binding:"max=255"`
		Phone    string `json:"phone"binding:"len=11"`
		Qq       string `json:"qq"`
		Wx       string `json:"wx"`
	}

	appG := app.New(c)
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}
	// 判断用户是否存在
	if !models.IsExistUser(obj.UserId) {
		appG.Response(http.StatusOK, e.AccountDoesNotExist, nil)
		return
	}
	newBusinessCard := &models.BusinessCard{
		UserId:   obj.UserId,
		Name:     obj.Name,
		Company:  obj.Company,
		Position: obj.Position,
		Phone:    obj.Phone,
		Qq:       obj.Qq,
		Wx:       obj.Wx,
	}
	// 新建名片
	if !models.Create(newBusinessCard) {
		appG.Response(http.StatusOK, e.NewFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 更新个人名片
func UpdateBusinessCard(c *gin.Context) {
	var obj struct {
		BusinessId string `json:"business_id"binding:"uuid4"`
		Name       string `json:"name"binding:"max=10"`
		Company    string `json:"company"binding:"max=255"`
		Position   string `json:"position"binding:"max=255"`
		Phone      string `json:"phone"binding:"len=11"`
		Qq         string `json:"qq"`
		Wx         string `json:"wx"`
	}
	appG := app.New(c)
	// 解析参数
	if c.ShouldBindJSON(&obj) != nil {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}
	// 判断名片是否存在
	businessCard, err := models.FindBusinessCard(obj.BusinessId)
	if err != nil {
		appG.Response(http.StatusOK, e.BusinessCardDoesNotExist, nil)
		return
	}
	// 更新名片
	if !models.UpdateBusinessCard(businessCard, obj) {
		appG.Response(http.StatusOK, e.UpdateFailed, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

// 删除个人名片
func DeleteBusinessCard(c *gin.Context) {

	appG := app.New(c)
	key := c.Param("key")

	if key == "" {
		appG.Response(http.StatusOK, e.InvalidParameter, nil)
		return
	}

	// 判断名片是否存在
	businessCard, err := models.FindBusinessCard(key)
	if err != nil {
		appG.Response(http.StatusOK, e.BusinessCardDoesNotExist, nil)
		return
	}
	// 删除名片
	if !models.DeleteBusinessCard(businessCard) {
		appG.Response(http.StatusOK, e.FailedToDelete, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

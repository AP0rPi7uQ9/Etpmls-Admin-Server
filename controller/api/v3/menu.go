package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get all menu
// 获取全部菜单
func MenuGetAll(c *gin.Context)  {
	var m model.Menu
	j, err := m.MenuGetAll()
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Get, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Get, j)
	return
}


// Create Menu
// 创建菜单
func MenuCreate(c *gin.Context)  {
	var j model.ApiMenuCreate

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_BindData, nil, err)
		return
	}

	var m model.Menu
	err := m.MenuCreate(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.ERROR_MESSAGE_Create, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.SUCCESS_MESSAGE_Create, nil)
	return
}
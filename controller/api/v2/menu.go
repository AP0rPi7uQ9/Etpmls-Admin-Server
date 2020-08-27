package v2

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

func MenuGetAll(c *gin.Context)  {
	var (
		ctx []byte
		err error
	)

	if library.Config.App.Cache {
		ctx, err = library.Redis.Get(library.RedisCtx, core.Cache_MenuGetAll).Bytes()
		if err != nil {
			ctx, err = ioutil.ReadFile("./storage/menu/menu.json")
			if err != nil {
				core.JsonError(c, http.StatusBadRequest, core.ERROR_MenuGetCurrent_GET_MENU_FAILED, core.ERROR_MESSAGE_MenuGetCurrent_GET_MENU_FAILED, nil, err)
				return
			}
			// Save menu
			// 储存菜单
			err2 := library.Redis.Set(library.RedisCtx, core.Cache_MenuGetAll, ctx, 0).Err()
			if err2 != nil {
				core.LogError.Output(err)
			}
		}
	} else {
		ctx, err = ioutil.ReadFile("./storage/menu/menu.json")
		if err != nil {
			core.JsonError(c, http.StatusBadRequest, core.ERROR_MenuGetCurrent_GET_MENU_FAILED, core.ERROR_MESSAGE_MenuGetCurrent_GET_MENU_FAILED, nil, err)
			return
		}
	}

	var j interface{}
	err = json.Unmarshal(ctx, &j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_MenuGetCurrent_JSON_UNMARSHAL_FAILED, core.ERROR_MESSAGE_MenuGetCurrent_JSON_UNMARSHAL_FAILED, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_MenuGetCurrent, core.SUCCESS_MESSAGE_MenuGetCurrent, j)
	return
}

func MenuCreate(c *gin.Context)  {
	var j model.ApiMenuCreateV2

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_MenuCreate_Bind, core.ERROR_MESSAGE_MenuCreate_Bind, nil, err)
		return
	}

	// Move files
	// 移动文件
	err := os.Rename("storage/menu/menu.json", "storage/menu/menu.json.bak")
	if err != nil {
		core.LogError.AutoOutputDebug("备份菜单文件失败！", err)
		core.JsonError(c, http.StatusBadRequest, core.ERROR_MenuCreate_Move_failed, core.ERROR_MESSAGE_MenuCreate_Move_failed, nil, err)
		return
	}

	// Write file
	// 写入文件
	var s = []byte(j.Menu)
	err = ioutil.WriteFile("storage/menu/menu.json", s, 0666)
	if err != nil {
		core.LogError.AutoOutputDebug("写入菜单文件失败！", err)
		core.JsonError(c, http.StatusBadRequest, core.ERROR_MenuCreate_Write, core.ERROR_MESSAGE_MenuCreate_Write, nil, err)

		// 还原历史菜单
		err = os.Rename("storage/menu/menu.json.bak", "storage/menu/menu.json")
		if err != nil {
			core.LogError.AutoOutputDebug("备份菜单文件还原失败！", err)
		}

		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_MenuCreate, core.SUCCESS_MESSAGE_MenuCreate, j)
	return
}
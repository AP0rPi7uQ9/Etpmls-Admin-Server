package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/model"
	"Etpmls-Admin-Server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Clear all cache
// 清除全部缓存
func SettingClearCache(c *gin.Context)  {
	// If the cache is not turned on, return to the prompt
	// 如果没开启缓存，返回提示
	if !library.Config.App.Cache {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_CacheIsNotEnabled"), nil, nil)
		return
	}

	library.Cache.ClearAllCache()
	core.LogDebug.Output(utils.MessageWithLineNum("Cleared all cache!"))
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Delete"), nil)
}


// Disk Cleanup
// 清理磁盘
func SettingDiskCleanup(c *gin.Context)  {
	var attachment model.Attachment
	err := attachment.DeleteUnused()
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Delete"), nil, err)
		return
	}

	core.LogDebug.Output(utils.MessageWithLineNum("Disk cleanup complete!"))
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Delete"), nil)
}
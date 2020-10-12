package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Clear all cache
// 清除全部缓存
func SettingClearCache(c *gin.Context)  {
	library.Cache.ClearAllCache()
	core.LogInfo.Output(core.MessageWithLineNum("Cleared all cache!"))
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Delete"), nil)
}
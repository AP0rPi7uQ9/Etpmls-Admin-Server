package route

import (
	"Etpmls-Admin-Server/controller/api/v3"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteApi(r *gin.Engine)  {
	if library.Config.App.ServiceDiscovery {
		r.GET(library.Config.ServiceDiscovery.Service.CheckUrl, func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Status": "Running",
			})
		})
	}
	r.GET("/api/v3/captcha/getOne", v3.CaptchaGetOne)
	r.GET("/api/v3/captcha/getPicture/:captchaId", v3.CaptchaGetPicture)
	r.POST("/api/v3/user/register", v3.UserRegister)
	r.POST("/api/v3/user/login", v3.UserLogin)

	api:= r.Group("/api")
	{
		version2 := api.Group("/v3")
		{
			version2.POST("/user/logout", v3.UserLogout, middleware.BasicCheck())

			user := version2.Group("/user")
			{
				user.GET("/getCurrent", v3.UserGetCurrent, middleware.BasicCheck())
				user.GET("/getAll", v3.UserGetAll, middleware.RoleCheck())
				user.POST("/create", v3.UserCreate, middleware.RoleCheck())
				user.PUT("/edit", v3.UserEdit, middleware.RoleCheck())
				user.DELETE("/delete", v3.UserDelete, middleware.RoleCheck())
				user.PUT("/updateInformation", v3.UserUpdateInformation, middleware.BasicCheck())
			}
			role := version2.Group("/role", middleware.RoleCheck())
			{
				role.GET("/getAll", v3.RoleGetAll)
				role.POST("/create", v3.RoleCreate)
				role.PUT("/edit", v3.RoleEdit)
				role.DELETE("/delete", v3.RoleDelete)
			}
			permission := version2.Group("/permission", middleware.RoleCheck())
			{
				permission.GET("/getAll", v3.PermissionGetAll)
				permission.POST("/create", v3.PermissionCreate)
				permission.PUT("/edit", v3.PermissionEdit)
				permission.DELETE("/delete", v3.PermissionDelete)
			}
			menu := version2.Group("/menu")
			{
				menu.GET("/getAll", v3.MenuGetAll, middleware.BasicCheck())
				menu.POST("/create", v3.MenuCreate, middleware.RoleCheck())
			}
			attachment := version2.Group("/attachment", middleware.RoleCheck())
			{
				attachment.POST("/uploadImage", v3.AttachmentUploadImage)
				attachment.DELETE("/deleteImage", v3.AttachmentDeleteImage)
			}
			setting := version2.Group("/setting", middleware.RoleCheck())
			{
				setting.GET("/clearCache", v3.SettingClearCache)
				setting.GET("/diskCleanup", v3.SettingDiskCleanup)
			}
		}
	}
}
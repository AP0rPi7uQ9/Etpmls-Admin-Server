package route

import (
	v2 "Etpmls-Admin-Server/controller/api/v2"
	"Etpmls-Admin-Server/module"
	"Etpmls-Admin-Server/middleware"
	"github.com/gin-gonic/gin"
)

func RouteApi(r *gin.Engine)  {
	r.GET("/api/v2/captcha/getOne", v2.CaptchaGetOne)
	r.GET("/api/v2/captcha/getPicture/:captchaId", v2.CaptchaGetPicture)
	r.POST("/api/v2/user/register", v2.UserRegister)
	r.POST("/api/v2/user/login", v2.UserLogin)

	api:= r.Group("/api")
	{
		version2 := api.Group("/v2")
		{
			version2.POST("/user/logout", v2.UserLogout, middleware.BasicCheck())

			user := version2.Group("/user", middleware.RoleCheck())
			{
				user.GET("/getCurrent", v2.UserGetCurrent)
				user.GET("/getAll", v2.UserGetAll)
				user.POST("/create", v2.UserCreate)
				user.PUT("/edit", v2.UserEdit)
				user.DELETE("/delete", v2.UserDelete)
			}
			role := version2.Group("/role", middleware.RoleCheck())
			{
				role.GET("/getAll", v2.RoleGetAll)
				role.POST("/create", v2.RoleCreate)
				role.PUT("/edit", v2.RoleEdit)
				role.DELETE("/delete", v2.RoleDelete)
			}
			permission := version2.Group("/permission", middleware.RoleCheck())
			{
				permission.GET("/getAll", v2.PermissionGetAll)
				permission.POST("/create", v2.PermissionCreate)
				permission.PUT("/edit", v2.PermissionEdit)
				permission.DELETE("/delete", v2.PermissionDelete)
			}
			menu := version2.Group("/menu", middleware.RoleCheck())
			{
				menu.GET("/getAll", v2.MenuGetAll)
				menu.POST("/create", v2.MenuCreate)
			}
			attachment := version2.Group("/attachment", middleware.RoleCheck())
			{
				attachment.POST("/uploadImage", v2.AttachmentUploadImage)
				attachment.DELETE("/deleteImage", v2.AttachmentDeleteImage)
			}
		}
	}

	/*r.POST("/api/v1/user/register", v1.UserRegister)
	r.POST("/api/v1/user/login", v1.UserLogin)
	r.GET("/api/v1/library/captcha/get/one", v1.LibraryCaptchaGetOne)


	api:= r.Group("/api", middleware.AuthCheck())
	{
		// v1
		version1 := api.Group("/v1")
		{
			user := version1.Group("/user")
			{
				user.POST("/logout", v1.UserLogout)
				user.GET("/get/current", v1.UserGetCurrent)
				user.GET("/get/all", v1.UserGetAll)
				user.POST("/create", v1.UserCreate)
				user.PUT("/edit/:id", v1.UserEdit)
				user.DELETE("/delete/:id", v1.UserDelete)
			}
			role := version1.Group("/role")
			{
				role.GET("/get/all", v1.RoleGetAll)
				role.POST("/create", v1.RoleCreate)
				role.PUT("/edit/:id", v1.RoleEdit)
				role.DELETE("/delete/:id", v1.RoleDelete)
			}
			menu := version1.Group("/menu")
			{
				menu.GET("/get/all", v1.MenuGetAll)
				menu.POST("/create", v1.MenuCreate)
				menu.DELETE("/delete/:id", v1.MenuDelete)
				menu.PUT("/edit/:id", v1.MenuEdit)
			}
			attachment := version1.Group("/attachment")
			{
				attachment.POST("/image/upload", v1.AttachmentImageUpload)
				attachment.DELETE("/image/delete", v1.AttachmentImageDelete)

			}
		}
		// common
		versonCommon := api.Group("/common")
		{
			role := versonCommon.Group("/role")
			{
				role.GET("/get/all", common.RoleGetAll)
			}
			user := versonCommon.Group("/user")
			{
				user.GET("/get/all", common.UserGetAll)
			}
			app := versonCommon.Group("/app")
			{
				app.GET("/get/detail", common.GetAppDetail)
			}
		}
	}*/

	module.Module_RouteApi(r)
}
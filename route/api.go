package route

import (
	v2 "Etpmls-Admin-Server/controller/api/v2"
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

			user := version2.Group("/user")
			{
				user.GET("/getCurrent", v2.UserGetCurrent, middleware.BasicCheck())
				user.GET("/getAll", v2.UserGetAll, middleware.RoleCheck())
				user.POST("/create", v2.UserCreate, middleware.RoleCheck())
				user.PUT("/edit", v2.UserEdit, middleware.RoleCheck())
				user.DELETE("/delete", v2.UserDelete, middleware.RoleCheck())
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
			menu := version2.Group("/menu")
			{
				menu.GET("/getAll", v2.MenuGetAll, middleware.BasicCheck())
				menu.POST("/create", v2.MenuCreate, middleware.RoleCheck())
			}
			attachment := version2.Group("/attachment", middleware.RoleCheck())
			{
				attachment.POST("/uploadImage", v2.AttachmentUploadImage)
				attachment.DELETE("/deleteImage", v2.AttachmentDeleteImage)
			}
		}
	}
}
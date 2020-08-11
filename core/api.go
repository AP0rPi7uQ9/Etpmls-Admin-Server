package core

import (
	"Etpmls-Admin-Server/library"
	"github.com/gin-gonic/gin"
)

// Error Code
// 错误码
const (
	USER_LOGIN_BIND_ERROR = "000001"
	USER_LOGIN_BIND_ERROR_MESSAGE = "用户登录信息的请求参数存在错误！"
	USER_LOGIN_VALIDATE_ERROR = "000002"
	USER_LOGIN_VALIDATE_ERROR_MESSAGE = "用户登录信息的请求参数验证失败！"
	USER_LOGIN_BIND_WITHOUT_CAPTCHA_ERROR = "000003"
	USER_LOGIN_BIND_WITHOUT_CAPTCHA_ERROR_MESSAGE = "用户登录信息的请求参数存在错误！"
	USER_LOGIN_VALIDATE_WITHOUT_CAPTCHA_ERROR = "000004"
	USER_LOGIN_VALIDATE_WITHOUT_CAPTCHA_ERROR_MESSAGE = "用户登录信息的请求参数验证失败！"
	USER_LOGIN_ERROR = "000005"
	USER_LOGIN_ERROR_MESSAGE = "用户登录失败！"
	USER_LOGIN_WITHOUT_CAPTCHA_ERROR = "000006"
	USER_LOGIN_WITHOUT_CAPTCHA_ERROR_MESSAGE = "用户登录失败！"
	USER_GET_TOKEN_ERROR = "000007"
	USER_GET_TOKEN_ERROR_MESSAGE = "获取令牌失败！"
	USER_GET_TOKEN_WITHOUT_CAPTCHA_ERROR = "000008"
	USER_GET_TOKEN_WITHOUT_CAPTCHA_ERROR_MESSAGE = "获取令牌失败！"
	USER_CLOSE_REGISTER_ERROR = "000009"
	USER_CLOSE_REGISTER_ERROR_MESSAGE = "管理员关闭了注册功能！"
	USER_REGISTER_BIND_ERROR = "000010"
	USER_REGISTER_BIND_ERROR_MESSAGE = "用户注册信息的请求参数存在错误！"
	USER_REGISTER_VALIDATE_ERROR = "000011"
	USER_REGISTER_VALIDATE_ERROR_MESSAGE = "用户注册信息的请求参数验证失败！"
	USER_EXISTS_ERROR = "000012"
	USER_EXISTS_ERROR_MESSAGE = "用户名已存在！"
	USER_REGISTER_ERROR = "000013"
	USER_REGISTER_ERROR_MESSAGE = "用户注册失败！"
	MIDDLEWARE_GET_TOKEN_ERROR = "000014"
	MIDDLEWARE_GET_TOKEN_ERROR_MESSAGE = "获取令牌失败！"
	MIDDLEWARE_PARSE_TOKEN_ERROR = "000015"
	MIDDLEWARE_PARSE_TOKEN_ERROR_MESSAGE = "令牌校验失败！"
	UserGetCurrent_GET_TOKEN_ERROR = "000016"
	UserGetCurrent_GET_TOKEN_ERROR_MESSAGE = "获取令牌失败！"
	ERROR_UserGetCurrent_GET_USER_FAILED = "000017"
	ERROR_MESSAGE_UserGetCurrent_GET_USER_FAILED = "获取用户信息失败！"
	ERROR_MenuGetCurrent_GET_MENU_FAILED = "000018"
	ERROR_MESSAGE_MenuGetCurrent_GET_MENU_FAILED = "获取菜单失败！"
	ERROR_MenuGetCurrent_JSON_UNMARSHAL_FAILED = "000019"
	ERROR_MESSAGE_MenuGetCurrent_JSON_UNMARSHAL_FAILED = "JSON反序列化失败！"
	ERROR_RoleCreate_BIND = "000020"
	ERROR_MESSAGE_RoleCreate_BIND = "提交的参数存在错误！"
	ERROR_RoleCreate_VALIDATE = "000021"
	ERROR_MESSAGE_RoleCreate_VALIDATE = "参数验证失败！"
	ERROR_RoleCreate_ROLE_NAME_EXISTS = "000022"
	ERROR_MESSAGE_RoleCreate_ROLE_NAME_EXISTS = "当前角色名已存在！"
	ERROR_RoleCreate_ROLE_CREATE = "000023"
	ERROR_MESSAGE_RoleCreate_ROLE_CREATE = "角色创建失败！"
	ERROR_RoleEdit_BIND = "000024"
	ERROR_MESSAGE_RoleEdit_BIND = "提交的参数存在错误！"
	ERROR_RoleEdit_VALIDATE = "000025"
	ERROR_MESSAGE_RoleEdit_VALIDATE = "参数验证失败！"
	ERROR_RoleEdit_Duplicate_role = "000026"
	ERROR_MESSAGE_RoleEdit_Duplicate_role = "该角色名已存在！"
	ERROR_RoleEdit = "000027"
	ERROR_MESSAGE_RoleEdit = "更新角色失败！"
	ERROR_RoleDelete_Bind = "000028"
	ERROR_MESSAGE_RoleDelete_Bind = "提交的参数存在错误！"
	ERROR_RoleDelete = "000029"
	ERROR_MESSAGE_RoleDelete = "角色删除失败！"
	ERROR_RoleDelete_Remove_admin_role = "000030"
	ERROR_MESSAGE_RoleDelete_Remove_admin_role = "不能删除管理员角色！"
	ERROR_UserCreate_Bind = "000031"
	ERROR_MESSAGE_UserCreate_Bind = "提交的参数存在错误！"
	ERROR_UserCreate_VALIDATE = "000032"
	ERROR_MESSAGE_UserCreate_VALIDATE = "参数验证失败！"
	ERROR_UserCreate_Count = "000033"
	ERROR_MESSAGE_UserCreate_Count = "查询用户名失败！"
	ERROR_UserCreate_Duplicate_user = "000034"
	ERROR_MESSAGE_UserCreate_Duplicate_user = "当前用户名已存在！"
	ERROR_UserCreate = "000035"
	ERROR_MESSAGE_UserCreate = "创建用户失败！"
	ERROR_UserEdit_Bind = "000036"
	ERROR_MESSAGE_UserEdit_Bind = "提交的参数存在错误！"
	ERROR_UserEdit_Validate = "000037"
	ERROR_MESSAGE_UserEdit_Validate = "参数验证失败！"
	ERROR_UserEdit_Duplicate_user = "000038"
	ERROR_MESSAGE_UserEdit_Duplicate_user = "当前用户名已存在！"


	ERROR_UserEdit = "000040"
	ERROR_MESSAGE_UserEdit = "更新用户失败！"
	ERROR_UserDelete_Bind = "000041"
	ERROR_MESSAGE_UserDelete_Bind = "提交的参数存在错误！"
	ERROR_UserDelete_Remove_admin = "000042"
	ERROR_MESSAGE_UserDelete_Remove_admin = "不能删除管理员！"
	ERROR_UserDelete = "000043"
	ERROR_MESSAGE_UserDelete = "删除用户失败！"
	ERROR_MenuCreate_Move_failed = "000044"
	ERROR_MESSAGE_MenuCreate_Move_failed = "菜单备份失败！"
	ERROR_MenuCreate_Bind = "000045"
	ERROR_MESSAGE_MenuCreate_Bind = "提交的参数存在错误！"
	ERROR_MenuCreate_Write = "000046"
	ERROR_MESSAGE_MenuCreate_Write = "写入菜单文件失败！"
	ERROR_PermissionCreate_Bind = "000047"
	ERROR_MESSAGE_PermissionCreate_Bind = "提交的参数存在错误！"
	ERROR_PermissionCreate_Validate = "000048"
	ERROR_MESSAGE_PermissionCreate_Validate = "参数验证失败！"
	ERROR_PermissionCreate = "000049"
	ERROR_MESSAGE_PermissionCreate = "权限创建失败！"
	ERROR_PermissionEdit_Bind = "000050"
	ERROR_MESSAGE_PermissionEdit_Bind = "提交的参数存在错误！"
	ERROR_PermissionEdit_Validate = "000051"
	ERROR_MESSAGE_PermissionEdit_Validate = "参数验证失败！"
	ERROR_PermissionEdit = "000052"
	ERROR_MESSAGE_PermissionEdit = "权限更新失败！"
	ERROR_PermissionDelete_Bind = "000053"
	ERROR_MESSAGE_PermissionDelete_Bind = "提交的参数存在错误！"
	ERROR_PermissionDelete = "000054"
	ERROR_MESSAGE_PermissionDelete = "权限删除失败！"
	ERROR_Permission_Check = "000055"
	ERROR_MESSAGE_Permission_Check = "您没有权限！"
	ERROR_AttachmentUploadImage_Get_file = "000056"
	ERROR_MESSAGE_AttachmentUploadImage_Get_file = "获取文件失败！"
	ERROR_AttachmentUploadImage_Validate = "000057"
	ERROR_MESSAGE_AttachmentUploadImage_Validate = "图片验证失败！"
	ERROR_AttachmentUploadImage = "000058"
	ERROR_MESSAGE_AttachmentUploadImage = "图片上传失败！"
	ERROR_AttachmentDeleteImage_Bind = "000059"
	ERROR_MESSAGE_AttachmentDeleteImage_Bind = "提交参数不完整！"
	ERROR_AttachmentDeleteImage_Validate = "000060"
	ERROR_MESSAGE_AttachmentDeleteImage_Validate = "数据有误，验证失败！"
	ERROR_AttachmentDeleteImage_Validate_path = "000061"
	ERROR_MESSAGE_AttachmentDeleteImage_Validate_path = "删除的路径有误，禁止删除！"
	ERROR_AttachmentDeleteImage = "000062"
	ERROR_MESSAGE_AttachmentDeleteImage = "删除失败！"
)

// Success Code
// 成功码
const (
	USER_LOGIN_SUCCESS = "100000"
	USER_LOGIN_SUCCESS_MESSAGE = "用户登录成功！"
	USER_LOGIN_WITHOUT_CAPTCHA_SUCCESS = "100001"
	USER_LOGIN_WITHOUT_CAPTCHA_SUCCESS_MESSAGE = "用户登录成功！"
	USER_REGISTER_SUCCESS = "100002"
	USER_REGISTER_SUCCESS_MESSAGE = "用户注册成功！"
	USER_GET_ALL_SUCCESS = "100003"
	USER_GET_ALL_SUCCESS_MESSAGE = "获取用户列表成功！"
	SUCCESS_UserGetCurrent = "100004"
	SUCCESS_MESSAGE_UserGetCurrent = "获取当前用户信息成功！"
	SUCCESS_MenuGetCurrent = "100005"
	SUCCESS_MESSAGE_MenuGetCurrent = "获取菜单成功！"
	SUCCESS_RoleCreate = "100006"
	SUCCESS_MESSAGE_RoleCreate = "角色创建成功！"
	SUCCESS_RoleGetAll = "100007"
	SUCCESS_MESSAGE_RoleGetAll = "获取角色列表成功！"
	SUCCESS_RoleEdit = "100008"
	SUCCESS_MESSAGE_RoleEdit = "角色创建成功！"
	SUCCESS_RoleDelete = "100009"
	SUCCESS_MESSAGE_RoleDelete = "角色删除成功！"
	SUCCESS_UserCreate = "100010"
	SUCCESS_MESSAGE_UserCreate = "用户创建成功！"
	SUCCESS_UserEdit = "100011"
	SUCCESS_MESSAGE_UserEdit = "更新用户成功！"
	SUCCESS_UserDelete = "100012"
	SUCCESS_MESSAGE_UserDelete = "更新用户成功！"
	SUCCESS_MenuCreate = "100013"
	SUCCESS_MESSAGE_MenuCreate = "创建菜单成功！"
	SUCCESS_PermissionCreate = "100014"
	SUCCESS_MESSAGE_PermissionCreate = "创建权限成功！"
	SUCCESS_PermissionGetAll = "100015"
	SUCCESS_MESSAGE_PermissionGetAll = "权限获取成功！"
	SUCCESS_PermissionEdit = "100016"
	SUCCESS_MESSAGE_PermissionEdit = "权限修改成功！"
	SUCCESS_PermissionDelete = "100017"
	SUCCESS_MESSAGE_PermissionDelete = "权限删除成功！"
	SUCCESS_UserLogout = "100018"
	SUCCESS_MESSAGE_UserLogout = "用户登出成功！"
	SUCCESS_CaptchaGetOne = "100019"
	SUCCESS_MESSAGE_CaptchaGetOne = "验证码获取成功！"
	SUCCESS_UploadImage = "100020"
	SUCCESS_MESSAGE_UploadImage = "图片上传成功！"
	SUCCESS_AttachmentDeleteImage = "100021"
	SUCCESS_MESSAGE_AttachmentDeleteImage = "图片删除成功！"
)

var (
	api_code = library.Config.App.Api.Field.Code
	api_status = library.Config.App.Api.Field.Status
	api_message = library.Config.App.Api.Field.Message
	api_data = library.Config.App.Api.Field.Data
)


func JsonError(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{}, err error)  {
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if library.Config.App.Api.UseHttpCode == true {
		code = httpStatusCode
	}

	// If it is a Debug environment, return information with Error
	// 如果是Debug环境，返回带有Error的信息
	if IsDebug() && err != nil {
		c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"error", api_message: message + "Error: " + err.Error(), api_data: data})
		return
	}

	c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"error", api_message: message, api_data: data})
	return
}

func JsonSuccess(c *gin.Context, httpStatusCode int, code interface{}, message string, data interface{})  {
	// If enabled, use HTTP CODE instead of system default CODE
	// 如果开启使用HTTP CODE 代替系统的默认CODE
	if library.Config.App.Api.UseHttpCode == true {
		code = httpStatusCode
	}

	c.JSON(httpStatusCode, gin.H{api_code: code, api_status:"success", api_message: message, api_data: data})
	return
}
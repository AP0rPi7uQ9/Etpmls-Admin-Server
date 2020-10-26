package v3

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// User registration
// 用户注册
func UserRegister(c *gin.Context)  {
	// Registration is not enabled
	// 注册功能未开启
	if library.Config.App.Register == false {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_RegistrationClosed"), nil, nil)
		return
	}

	var j model.Api_UserRegister

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	//Validate
	err := library.ValidateRegister(&j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	//Validate User If exists
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", j.Username).Count(&count)
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_DuplicateUserName"), nil, err)
		return
	}

	//Create User
	var u model.User
	id, err := u.Register(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Register"), nil, err)
		return
	}

	//Return Token
	data := map[string]uint{"id": id}
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Register"), data)
	return
}

// User login
// 用户登录
func UserLogin(c *gin.Context)  {

	// Turn off verification code
	// 关闭验证码
	if library.Config.App.Captcha == false {
		UserLoginWithoutCaptcha(c)
		return
	}

	var j model.Api_UserLogin

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	//Validate
	err := library.ValidateZh(&j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	var u model.User
	usr, err := u.Login(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Login"), nil, err)
		return
	}

	//JWT
	var us model.User
	token, err := us.UserGetToken(usr.ID, usr.Username)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_GetToken"), nil, err)
		return
	}

	//Return Token
	resData := make(map[string]string)
	resData["token"] = token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Login"), resData)
}

// User login Without Captcha
// 用户无验证码登录
func UserLoginWithoutCaptcha(c *gin.Context)  {
	var j model.ApiUserLoginWithoutCaptcha

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	//Validate
	err := library.ValidateZh(&j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	var u model.User
	usr, err := u.UserLoginWithoutCaptcha(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Login"), nil, err)
		return
	}

	//JWT
	var us model.User
	token, err := us.UserGetToken(usr.ID, usr.Username)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_GetToken"), nil, err)
		return
	}

	//Return Token
	resData := make(map[string]string)
	resData["token"] = token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Login"), resData)
}

// 用户登出
func UserLogout(c *gin.Context)  {
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Logout"), nil)
	return
}

// Get all users
// 获取所有用户
func UserGetAll(c *gin.Context)  {
	var u model.User
	data := u.UserGetAll(c)
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Get"), data)
	return
}

// Get current user
// 获取当前用户
func UserGetCurrent(c *gin.Context)  {
	var us model.User
	u, err := us.UserGetCurrent(c)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Get"), nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Get"), u)
	return
}

// Create User
// 创建用户
func UserCreate(c *gin.Context)  {
	var j model.ApiUserCreate

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	// Validate Json
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	// Validate Username unique
	var count int64
	if err := database.DB.Model(&model.User{}).Where("username = ?", j.Username).Count(&count).Error; err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Get"), nil, err)
		return
	}
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_DuplicateUserName"), nil, err)
		return
	}

	//Create User
	var u model.User
	err = u.UserCreate(c, j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Create"), nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Create"), nil)
	return
}

// Edit user
// 编辑用户
func UserEdit(c *gin.Context)  {
	var j model.ApiUserEdit

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	// Validate Json
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	// Validate Username Unique
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", j.Username).Where("id != ?", j.ID).Count(&count)
	if (count > 0) {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_DuplicateUserName"), nil, err)
		return
	}

	var u model.User
	err = u.UserEdit(c, j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Edit"), nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Edit"), nil)
	return
}

// Delete User
// 删除用户
func UserDelete(c *gin.Context)  {
	var j model.ApiUserDelete
	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	var ids []uint
	for _, v := range j.Users {
		ids = append(ids, v.ID)
	}

	// Find if admin is included in ids
	// 查找ids中是否包含admin
	b := model.Common_CheckIfOneIsIncludeInIds(ids)
	if b {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_PermissionDenied"), nil, nil)
		return
	}

	var u model.User
	err := u.UserDelete(c, ids)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Delete"), nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Delete"), nil)
	return
}

// Update user information
// 更新用户信息
func UserUpdateInformation(c *gin.Context)  {
	var j model.ApiUserUpdateInformation

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_BindData"), nil, err)
		return
	}

	// Validate Json
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Validate"), nil, err)
		return
	}

	// Get current user id
	var us model.User
	j.ID, err = us.User_GetUserIdByRequest(c)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Get"), nil, err)
		return
	}

	var u model.User
	err = u.UserUpdateInformation(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_Code, core.Translate(c, "ERROR_MESSAGE_Edit"), nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Code, core.Translate(c, "SUCCESS_MESSAGE_Edit"), nil)
	return
}
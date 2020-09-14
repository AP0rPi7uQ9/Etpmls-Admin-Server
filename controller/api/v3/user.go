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
		core.JsonError(c, http.StatusBadRequest, core.USER_CLOSE_REGISTER_ERROR, core.USER_CLOSE_REGISTER_ERROR_MESSAGE, nil, nil)
		return
	}

	var j model.ApiUserRegister

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_REGISTER_BIND_ERROR, core.USER_REGISTER_BIND_ERROR_MESSAGE, nil, err)
		return
	}

	//Validate
	err := library.ValidateRegister(&j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_REGISTER_VALIDATE_ERROR, core.USER_REGISTER_VALIDATE_ERROR_MESSAGE, nil, err)
		return
	}

	//Validate User If exists
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", j.Username).Count(&count)
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.USER_EXISTS_ERROR, core.USER_EXISTS_ERROR_MESSAGE, nil, err)
		return
	}

	//Create User
	var u model.User
	id, err := u.UserRegister(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_REGISTER_ERROR, core.USER_REGISTER_ERROR_MESSAGE, nil, err)
		return
	}

	//Return Token
	data := map[string]uint{"id": id}
	core.JsonSuccess(c, http.StatusOK, core.USER_REGISTER_SUCCESS, core.USER_REGISTER_SUCCESS_MESSAGE, data)
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

	var j model.ApiUserLogin

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_BIND_ERROR, core.USER_LOGIN_BIND_ERROR_MESSAGE, nil, err)
		return
	}

	//Validate
	err := library.ValidateZh(&j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_VALIDATE_ERROR, core.USER_LOGIN_VALIDATE_ERROR_MESSAGE, nil, err)
		return
	}

	var u model.User
	id, username, err := u.UserLogin(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_ERROR, core.USER_LOGIN_ERROR_MESSAGE, nil, err)
		return
	}

	//JWT
	var us model.User
	token, err := us.UserGetToken(id, username)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_GET_TOKEN_ERROR, core.USER_GET_TOKEN_ERROR_MESSAGE, nil, err)
		return
	}

	//Return Token
	resData := make(map[string]string)
	resData["token"] = token
	core.JsonSuccess(c, http.StatusOK, core.USER_LOGIN_SUCCESS, core.USER_LOGIN_SUCCESS_MESSAGE, resData)
}

// User login Without Captcha
// 用户无验证码登录
func UserLoginWithoutCaptcha(c *gin.Context)  {
	var j model.ApiUserLoginWithoutCaptcha

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_BIND_WITHOUT_CAPTCHA_ERROR, core.USER_LOGIN_BIND_WITHOUT_CAPTCHA_ERROR_MESSAGE, nil, err)
		return
	}

	//Validate
	err := library.ValidateZh(&j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_VALIDATE_WITHOUT_CAPTCHA_ERROR, core.USER_LOGIN_VALIDATE_WITHOUT_CAPTCHA_ERROR_MESSAGE, nil, err)
		return
	}

	var u model.User
	id, username, err := u.UserLoginWithoutCaptcha(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_WITHOUT_CAPTCHA_ERROR, core.USER_LOGIN_WITHOUT_CAPTCHA_ERROR_MESSAGE, nil, err)
		return
	}

	//JWT
	var us model.User
	token, err := us.UserGetToken(id, username)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_GET_TOKEN_WITHOUT_CAPTCHA_ERROR, core.USER_GET_TOKEN_WITHOUT_CAPTCHA_ERROR_MESSAGE, nil, err)
		return
	}

	//Return Token
	resData := make(map[string]string)
	resData["token"] = token
	core.JsonSuccess(c, http.StatusOK, core.USER_LOGIN_WITHOUT_CAPTCHA_SUCCESS, core.USER_LOGIN_WITHOUT_CAPTCHA_SUCCESS_MESSAGE, resData)
}

// 用户登出
func UserLogout(c *gin.Context)  {
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UserLogout, core.SUCCESS_MESSAGE_UserLogout, nil)
	return
}

// Get all users
// 获取所有用户
func UserGetAll(c *gin.Context)  {
	var u model.User
	data := u.UserGetAll(c)
	core.JsonSuccess(c, http.StatusOK, core.USER_GET_ALL_SUCCESS, core.USER_GET_ALL_SUCCESS_MESSAGE, data)
	return
}

// Get current user
// 获取当前用户
func UserGetCurrent(c *gin.Context)  {
	var us model.User
	u, err := us.UserGetCurrent(c)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserGetCurrent_GET_USER_FAILED, core.ERROR_MESSAGE_UserGetCurrent_GET_USER_FAILED, nil, err)
		return
	}

	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UserGetCurrent, core.SUCCESS_MESSAGE_UserGetCurrent, u)
	return
}

// Create User
// 创建用户
func UserCreate(c *gin.Context)  {
	var j model.ApiUserCreate

	//Bind Data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserCreate_Bind, core.ERROR_MESSAGE_UserCreate_Bind, nil, err)
		return
	}

	// Validate Json
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserCreate_VALIDATE, core.ERROR_MESSAGE_UserCreate_VALIDATE, nil, err)
		return
	}

	// Validate Username unique
	var count int64
	if err := database.DB.Model(&model.User{}).Where("username = ?", j.Username).Count(&count).Error; err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserCreate_Count, core.ERROR_MESSAGE_UserCreate_Count, nil, err)
		return
	}
	if count > 0 {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserCreate_Duplicate_user, core.ERROR_MESSAGE_UserCreate_Duplicate_user, nil, err)
		return
	}

	//Create User
	var u model.User
	err = u.UserCreate(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserCreate, core.ERROR_MESSAGE_UserCreate, nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UserCreate, core.SUCCESS_MESSAGE_UserCreate, nil)
	return
}

// Edit user
// 编辑用户
func UserEdit(c *gin.Context)  {
	var j model.ApiUserEdit

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserEdit_Bind, core.ERROR_MESSAGE_UserEdit_Bind, nil, err)
		return
	}

	// Validate Json
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserEdit_Validate, core.ERROR_MESSAGE_UserEdit_Validate, nil, err)
		return
	}

	// Validate Username Unique
	var count int64
	database.DB.Model(&model.User{}).Where("username = ?", j.Username).Where("id != ?", j.ID).Count(&count)
	if (count > 0) {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserEdit_Duplicate_user, core.ERROR_MESSAGE_UserEdit_Duplicate_user, nil, err)
		return
	}

	var u model.User
	err = u.UserEdit(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserEdit, core.ERROR_MESSAGE_UserEdit, nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UserEdit, core.SUCCESS_MESSAGE_UserEdit, nil)
	return
}

// Delete User
// 删除用户
func UserDelete(c *gin.Context)  {
	var j model.ApiUserDelete
	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserDelete_Bind, core.ERROR_MESSAGE_UserDelete_Bind, nil, err)
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
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserDelete_Remove_admin, core.ERROR_MESSAGE_UserDelete_Remove_admin, nil, nil)
		return
	}

	var u model.User
	err := u.UserDelete(ids)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserDelete, core.ERROR_MESSAGE_UserDelete, nil, err)
		return
	}

	// Return Message
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_UserDelete, core.SUCCESS_MESSAGE_UserDelete, nil)
	return
}

// Update user information
// 更新用户信息
func UserUpdateInformation(c *gin.Context)  {
	var j model.ApiUserUpdateInformation

	// Bind data
	if err := c.ShouldBindJSON(&j); err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation_Bind, core.ERROR_MESSAGE_Bind_data, nil, err)
		return
	}

	// Validate Json
	err := library.ValidateZh(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation_Validate, core.ERROR_MESSAGE_Validate, nil, err)
		return
	}

	// Get current user id
	var us model.User
	j.ID, err = us.User_GetUserIdByRequest(c)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation_Get_current_user_id, core.ERROR_MESSAGE_Get_current_user_information, nil, err)
		return
	}

	// Modification of other people’s information is not allowed
	/*if id != j.ID {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation_Edit_others, core.ERROR_MESSAGE_Permission_denied, nil, err)
		return
	}*/

	var u model.User
	err = u.UserUpdateInformation(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation, core.ERROR_MESSAGE_Internal, nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Internal, core.SUCCESS_MESSAGE_Internal, nil)
	return
}
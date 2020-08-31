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

	var j model.ApiUserRegisterV2

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
	id, err := model.UserRegisterV2(j)
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

	var j model.ApiUserLoginV2

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

	id, username, err := model.UserLoginV2(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_ERROR, core.USER_LOGIN_ERROR_MESSAGE, nil, err)
		return
	}

	//JWT
	token, err := model.UserGetTokenV2(id, username)
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
	var j model.ApiUserLoginWithoutCaptchaV2

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

	id, username, err := model.UserLoginWithoutCaptchaV2(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.USER_LOGIN_WITHOUT_CAPTCHA_ERROR, core.USER_LOGIN_WITHOUT_CAPTCHA_ERROR_MESSAGE, nil, err)
		return
	}

	//JWT
	token, err := model.UserGetTokenV2(id, username)
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
	data := model.UserGetAllV2(c)
	core.JsonSuccess(c, http.StatusOK, core.USER_GET_ALL_SUCCESS, core.USER_GET_ALL_SUCCESS_MESSAGE, data)
	return
}

// Get current user
// 获取当前用户
func UserGetCurrent(c *gin.Context)  {
	/*var a =  []string{"admin"}
	c.JSON(http.StatusOK, gin.H{"code": 200, "status":"success", "msg": "success", "data": gin.H{
		"avatar": "https://i.gtimg.cn/club/item/face/img/8/15918_100.gif",
		"username" : "admin",
		"roles": a,
	}})
	return*/


	//Get Token
	// 获取token
	/*token, err := core.GetToken(c)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.UserGetCurrent_GET_TOKEN_ERROR, core.UserGetCurrent_GET_TOKEN_ERROR_MESSAGE, nil, err)
		return
	}*/

	u, err := model.UserGetCurrentV3(c)
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
	var j model.ApiUserCreateV2

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
	err = model.UserCreateV2(j)
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
	var j model.ApiUserEditV3

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
	//Validate Password is not empty, Hash Password
	/*if len(j.Password) > 0 {
		j.Password, err = model.User_BcryptPasswordV2(j.Password)
		if err != nil {
			core.JsonError(c, http.StatusBadRequest, core.ERROR_UserEdit_Password_resolution, core.ERROR_MESSAGE_UserEdit_Password_resolution, nil, err)
			return
		}
	}*/

	err = model.UserEditV3(j)
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
	var j model.ApiUserDeleteV2
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

	err := model.UserDeleteV2(ids)
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
	var j model.ApiUserUpdateInformationV1

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
	j.ID, err = model.User_GetUserIdByRequest(c)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation_Get_current_user_id, core.ERROR_MESSAGE_Get_current_user_information, nil, err)
		return
	}

	// Modification of other people’s information is not allowed
	/*if id != j.ID {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation_Edit_others, core.ERROR_MESSAGE_Permission_denied, nil, err)
		return
	}*/

	err = model.UserUpdateInformationV1(j)
	if err != nil {
		core.JsonError(c, http.StatusBadRequest, core.ERROR_UserUpdateInformation, core.ERROR_MESSAGE_Internal, nil, err)
		return
	}

	//Return Token
	core.JsonSuccess(c, http.StatusOK, core.SUCCESS_Internal, core.SUCCESS_MESSAGE_Internal, nil)
	return
}
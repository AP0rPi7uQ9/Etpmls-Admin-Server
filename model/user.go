package model

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/library"
	"errors"
	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	ID        uint `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username string `json:"username"`
	Password string `json:"password"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}

type ApiUserRegisterV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"max=255"`
	Password string `binding:"required" json:"password" validate:"max=255"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}

type ApiUserLoginV2 struct{
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"max=255"`
	Password string `binding:"required" json:"password" validate:"max=255"`
	CaptchaId string `binding:"required" json:"captcha_id"`
	Captcha string `binding:"required" json:"captcha"`
}

type ApiUserLoginWithoutCaptchaV2 struct{
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"max=255"`
	Password string `binding:"required" json:"password" validate:"max=255"`
}

type ApiUserGetAllV2 struct {
	ID        uint `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username string `json:"username"`
	Password string `json:"-"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}

type ApiUserGetCurrentV2 struct {
	ID        uint `json:"id"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}

type ApiUserCreateV2 struct {
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"required,max=50"`
	Password string `binding:"required" json:"password" validate:"required,max=50"`
	Roles []Role `gorm:"many2many:role_users" binding:"required" json:"roles"`
}

type ApiUserEditV2 struct{
	ID uint `json:"id" binding:"required"`
	CreatedAt time.Time	`gorm:"-" json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"-" json:"-"`
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" validate:"max=50"`
	Roles []Role `gorm:"many2many:role_users" binding:"required" json:"roles"`
}

type ApiUserDeleteV2 struct {
	ID uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Users []User `json:"users" binding:"required" validate:"min=1"`
}



func UserRegisterV2(j ApiUserRegisterV2) (id uint, err error) {
	type User ApiUserRegisterV2
	var form = User(j)

	// Password bcrypt
	form.Password, err = UserBcryptPasswordV2(form.Password)
	if err != nil {
		return id, err
	}

	result := database.DB.Create(&form);
	if result.Error != nil {
		return id, result.Error
	}

	return form.ID, err
}

func UserLoginV2(j ApiUserLoginV2) (id uint, username string, err error) {
	// Validate Captcha
	if !captcha.VerifyString(j.CaptchaId, j.Captcha){
		return id, username, errors.New("验证码错误！")
	}

	usrID, usrName, err := UserVerifyV2(j.Username, j.Password)

	return usrID, usrName, err
}

func UserLoginWithoutCaptchaV2(j ApiUserLoginWithoutCaptchaV2) (id uint, username string, err error) {
	usrID, usrName, err := UserVerifyV2(j.Username, j.Password)

	return usrID, usrName, err
}

func UserGetAllV2(c *gin.Context) interface{} {
	// 重写ApiUserGetAllV2的Roles字段，防止泄露隐私字段信息
	type Role ApiRoleGetAllV2
	type User struct {
		ApiUserGetAllV2
		Roles []Role `gorm:"many2many:role_users" json:"roles"`
	}
	var data []User

	// 获取分页和标题
	limit, offset := CommonGetPageByQuery(c)
	var count int64
	// Get the title of the search, if not get all the data
	// 获取搜索的标题，如果没有获取全部数据
	search := c.Query("search")

	database.DB.Model(&User{}).Preload("Roles").Where("username " + database.FUZZY_SEARCH + " ?", "%"+ search +"%").Count(&count).Limit(limit).Offset(offset).Find(&data)

	return gin.H{"data": data, library.Config.App.Api.Pagination.Field.Count: count}
}

func UserCreateV2(j ApiUserCreateV2) (err error) {
	type User ApiUserCreateV2
	form := User(j)

	// Bcrypt Password
	form.Password, err = User_BcryptPasswordV2(j.Password)
	if err != nil {
		return errors.New("密码加密失败！")
	}

	result := database.DB.Create(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UserEditV2(j ApiUserEditV2) (err error) {
	// If user not set new password
	if len(j.Password) == 0 {
		// Find User
		var u User
		database.DB.First(&u, j.ID)
		// Set old password
		j.Password = u.Password
	} else {
		// Bcrypt Password
		j.Password, err = User_BcryptPasswordV2(j.Password)
		if err != nil {
			return errors.New("密码加密失败！")
		}
	}

	type User ApiUserEditV2
	form := User(j)

	// 删除关联
	err = database.DB.Model(&User{ID:form.ID}).Association("Roles").Clear()
	if err != nil {
		return err
	}
	// 创建数据及关联
	result := database.DB.Save(&form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Delete users (allow multiple deletions at the same time)
// 删除用户（允许同时删除多个）
func UserDeleteV2(ids []uint) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var u []User
		database.DB.Where("id IN ?", ids).Find(&u)

		// 删除用户
		result := database.DB.Delete(&u)
		if result.Error != nil {
			return result.Error
		}

		// 删除关联
		err = database.DB.Model(&u).Association("Roles").Clear()
		if err != nil {
			return err
		}

		return nil
	})

	return err
}


// Encrypt user password
// 加密用户密码
func UserBcryptPasswordV2(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verify user password
// 验证用户密码
func UserVerifyPasswordV2(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, err
}

// Verify user logic
// 验证用户逻辑
func UserVerifyV2(username string, password string) (id uint, unm string, err error) {
	//Search User
	var user User
	database.DB.Where("username = ?", username).First(&user)
	if !(user.ID > 0) {
		return id, unm, errors.New("该用户名不存在！")
	}

	//Password is wrong
	b, err := UserVerifyPasswordV2(password, user.Password)
	if err != nil || !b {
		return id, unm, errors.New("校验失败或密码错误！")
	}

	return user.ID, user.Username, err
}

// Get token by ID
// 通过ID获取Token
func UserGetTokenV2(userId uint, username string) (string, error) {
	var k = library.JwtGo {
		MySigningKey: []byte(library.Config.App.Key),
	}
	return k.JwtGoCreateToken(&jwt.StandardClaims{
		Id: strconv.Itoa(int(userId)),	// 用户ID
		ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),	// 过期时间 - 12个小时
		Issuer:    username,	// 发行者
	})

}

// 通过Token获取当前用户
func UserGetCurrentV2(token string) (u interface{}, err error) {
	// Get Claims
	// 获取Claims
	var k = library.JwtGo {
		MySigningKey: []byte(library.Config.App.Key),
	}
	id, err := k.JwtGoGetToeknId(token)
	if err != nil {
		return u, err
	}

	username, err  := k.JwtGoGetTokenIssuer(token)
	if err != nil {
		return u, err
	}

	type User ApiUserGetCurrentV2
	var data User
	result := database.DB.Where("id = ? AND username = ?", id, username).First(&data)
	if !(result.RowsAffected > 0) {
		return u, errors.New("没有在数据库中找到当前用户！")
	}

	return data, nil
}

// 加密密码
func User_BcryptPasswordV2(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


// 根据token获取用户
func User_GetUserByToken(token string) (u User, err error) {
	// 获取Claims
	var k = library.JwtGo {
		MySigningKey: []byte(library.Config.App.Key),
	}
	// 从Token获取ID
	id, err := k.JwtGoGetToeknId(token)
	if err != nil {
		return u, err
	}
	// 获取用户
	result := database.DB.First(&u, id)
	if result.Error != nil {
		return u, result.Error
	}
	return u, nil
}

func User_GetUserByRequest(c *gin.Context) (u User, err error) {
	token, err := core.GetToken(c)
	if err != nil {
		return u, err
	}
	u, err = User_GetUserByToken(token)
	if err != nil {
		return u, err
	}
	return u, err
}
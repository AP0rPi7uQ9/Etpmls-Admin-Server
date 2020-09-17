package model

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/library"
	"encoding/json"
	"errors"
	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID        uint `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username string `json:"username"`
	Password string `json:"password"`
	Avatar Attachment	`gorm:"polymorphic:Owner;polymorphicValue:user-avatar" json:"avatar"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}


// User register
// 用户注册
type ApiUserRegister struct {
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"max=255"`
	Password string `binding:"required" json:"password" validate:"max=255"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}
func (this *User) UserRegister(j ApiUserRegister) (id uint, err error) {
	type User ApiUserRegister
	var form = User(j)

	// Password bcrypt
	form.Password, err = this.UserBcryptPassword(form.Password)
	if err != nil {
		return id, err
	}

	result := database.DB.Create(&form);
	if result.Error != nil {
		return id, result.Error
	}

	return form.ID, err
}


// User login
// 用户登录
type ApiUserLogin struct{
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"max=255"`
	Password string `binding:"required" json:"password" validate:"max=255"`
	CaptchaId string `binding:"required" json:"captcha_id"`
	Captcha string `binding:"required" json:"captcha"`
}
func (this *User) UserLogin(j ApiUserLogin) (id uint, username string, err error) {
	// Validate Captcha
	if !captcha.VerifyString(j.CaptchaId, j.Captcha){
		return id, username, errors.New("验证码错误！")
	}

	usrID, usrName, err := this.UserVerify(j.Username, j.Password)

	return usrID, usrName, err
}


// User login without captcha
// 用户免验证码登录
type ApiUserLoginWithoutCaptcha struct{
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"max=255"`
	Password string `binding:"required" json:"password" validate:"max=255"`
}
func (this *User) UserLoginWithoutCaptcha(j ApiUserLoginWithoutCaptcha) (id uint, username string, err error) {
	usrID, usrName, err := this.UserVerify(j.Username, j.Password)

	return usrID, usrName, err
}


// Get all user
// 获取全部用户
type ApiUserGetAll struct {
	ID        uint `json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Username string `json:"username"`
	Password string `json:"-"`
	Roles []Role `gorm:"many2many:role_users" json:"roles"`
}
func (this *User) UserGetAll(c *gin.Context) interface{} {
	// 重写ApiUserGetAllV2的Roles字段，防止泄露隐私字段信息
	type Role ApiRoleGetAll
	type User struct {
		ApiUserGetAll
		Roles []Role `gorm:"many2many:role_users" json:"roles"`
	}
	var data []User

	// 获取分页和标题
	limit, offset := Common_GetPageByQuery(c)
	var count int64
	// Get the title of the search, if not get all the data
	// 获取搜索的标题，如果没有获取全部数据
	search := c.Query("search")

	database.DB.Model(&User{}).Preload("Roles").Where("username " + database.FUZZY_SEARCH + " ?", "%"+ search +"%").Count(&count).Limit(limit).Offset(offset).Find(&data)

	return gin.H{"data": data, library.Config.Field.Pagination.Count: count}
}


// Create User
// 创建用户
type ApiUserCreate struct {
	ID        uint `json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `binding:"required" json:"username" validate:"required,max=50"`
	Password string `binding:"required" json:"password" validate:"required,max=50"`
	Roles []Role `gorm:"many2many:role_users" binding:"required" json:"roles"`
}
func (this *User) UserCreate(c *gin.Context, j ApiUserCreate) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		type User ApiUserCreate
		form := User(j)

		// Bcrypt Password
		form.Password, err = this.User_BcryptPasswordV2(j.Password)
		if err != nil {
			return errors.New("密码加密失败！")
		}

		result := tx.Create(&form)
		if result.Error != nil {
			return result.Error
		}

		// User Create Event
		u, err := this.User_InterfaceToUser(form)
		if err != nil {
			return err
		}
		select {
		case core.Event.Event_UserCreate <- u:
		case <- time.After(time.Second * 3):
		}

		return nil
	})


	return err
}


// Edit User
// 编辑用户
type ApiUserEdit struct{
	ID uint `json:"id" binding:"required"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"-" json:"-"`
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" validate:"max=50"`
	Roles []Role `gorm:"many2many:role_users" binding:"required" json:"roles"`
}
func (this *User) UserEdit(c *gin.Context, j ApiUserEdit) (err error) {
	// Find User
	var form User
	database.DB.First(&form, j.ID)

	// If user set new password
	if len(j.Password) > 0 {
		form.Password, err = this.User_BcryptPasswordV2(j.Password)
		if err != nil {
			return errors.New("密码加密失败！")
		}
	}

	form.Username = j.Username	// Username

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// 删除关联
		err = tx.Model(&User{ID:form.ID}).Association("Roles").Replace(j.Roles)
		if err != nil {
			return err
		}
		// 创建数据及关联
		result := tx.Save(&form)
		if result.Error != nil {
			return result.Error
		}

		// User Edit Event for module
		u, err := this.User_InterfaceToUser(form)
		if err != nil {
			return err
		}
		select {
		case core.Event.Event_UserEdit <- u:
		case <- time.After(time.Second * 3):
		}


		return nil
	})

	if library.Config.App.Cache {
		library.Cache.DeleteHash(core.Cache_UserGetCurrent, strconv.Itoa(int(j.ID)))
	}

	return err
}


// Delete users (allow multiple deletions at the same time)
// 删除用户（允许同时删除多个）
type ApiUserDelete struct {
	ID uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Users []User `json:"users" binding:"required" validate:"min=1"`
}
func (this *User) UserDelete(c *gin.Context, ids []uint) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var u []User
		tx.Where("id IN ?", ids).Find(&u)

		// 删除用户
		result := tx.Delete(&u)
		if result.Error != nil {
			return result.Error
		}

		// 删除关联
		err = tx.Model(&u).Association("Roles").Clear()
		if err != nil {
			return err
		}

		// User Delete Event for module
		select {
		case core.Event.Event_UserDelete <- u:
		case <- time.After(time.Second * 3):
		}

		return nil
	})

	if library.Config.App.Cache {
		var tmp []string
		for _, v := range ids {
			tmp = append(tmp, strconv.Itoa(int(v)))
		}
		library.Cache.DeleteHash(core.Cache_UserGetCurrent, strings.Join(tmp, " "))
	}

	return err
}


// Update user information
// 更新用户信息
type ApiUserUpdateInformation struct{
	ID uint `json:"-"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"-" json:"-"`
	Username string `json:"-"`
	Password string `json:"password" validate:"omitempty,min=6,max=50"`
	Avatar Attachment	`gorm:"polymorphic:Owner;polymorphicValue:user-avatar" json:"avatar"`
}
func (this *User) UserUpdateInformation(j ApiUserUpdateInformation) error {

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 如果表单包含缩略图，
		if len(j.Avatar.Path) > 0 {
			// 1.删除同名缓存
			result := tx.Unscoped().Where("path = ?", j.Avatar.Path).Delete(Attachment{})
			if result.Error != nil {
				return result.Error
			}
		}

		// 2.删除历史avatar
		var old Attachment
		result2 := tx.Where("owner_id = ?", j.ID).Where("owner_type = ?", "user-avatar").First(&old)
		// 如果找到记录则删除
		if result2.RowsAffected > 0 {
			// 根据Path删除附件
			err := old.AttachmentBatchDelete([]string{old.Path})
			if err != nil {
				return err
			}
		}

		// 3.新增avatar
		err := tx.Model(&User{ID: j.ID}).Association("Avatar").Replace(&Attachment{Path:j.Avatar.Path})
		if err != nil {
			return err
		}

		// 4.更新

		// Update password if exists
		if len(j.Password) > 0 {
			j.Password, err = this.User_BcryptPasswordV2(j.Password)
		}

		result := tx.Model(&User{ID: j.ID}).Updates(&j)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if library.Config.App.Cache {
		library.Cache.DeleteHash(core.Cache_UserGetCurrent, strconv.Itoa(int(j.ID)))
	}

	return err
}


// Encrypt user password
// 加密用户密码
func (this *User) UserBcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


// Verify user password
// 验证用户密码
func (this *User) UserVerifyPassword(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}
	return true, err
}


// Verify user logic
// 验证用户逻辑
func (this *User) UserVerify(username string, password string) (id uint, unm string, err error) {
	//Search User
	var user User
	database.DB.Where("username = ?", username).First(&user)
	if !(user.ID > 0) {
		return id, unm, errors.New("该用户名不存在！")
	}

	//Password is wrong
	b, err := this.UserVerifyPassword(password, user.Password)
	if err != nil || !b {
		return id, unm, errors.New("校验失败或密码错误！")
	}

	return user.ID, user.Username, err
}


// Get token by ID
// 通过ID获取Token
func (this *User) UserGetToken(userId uint, username string) (string, error) {
	return library.Jwt_Token.CreateToken(&jwt.StandardClaims{
		Id: strconv.Itoa(int(userId)),	// 用户ID
		ExpiresAt: time.Now().Add(time.Second * library.Config.App.TokenExpirationTime).Unix(),	// 过期时间 - 12个小时
		Issuer:    username,	// 发行者
	})

}


// 通过Token获取当前用户
type ApiUserGetCurrent struct {
	ID        uint `json:"id"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	Avatar string	`json:"avatar"`
	Roles []string `json:"roles"`
}
func (this *User) UserGetCurrent(c *gin.Context) (interface{}, error) {
	if library.Config.App.Cache {
		return this.user_GetCurrent_Cache(c)
	} else {
		return this.user_GetCurrent_NotCache(c)
	}
}
func (this *User) user_GetCurrent_NotCache(c *gin.Context) (interface{}, error) {
	// Get User By request
	u, err := this.User_GetUserByRequest(c)
	if err != nil {
		return nil, err
	}

	// Ignore the avatar tag in the User structure
	type tmp struct {
		User
		Avatar string	`json:"avatar"`
	}
	var tmpUser = tmp{User: u}

	var userApi ApiUserGetCurrent
	b, err := json.Marshal(tmpUser)
	if err != nil {
		return nil, err
	}
	// Get the filtered structure - ApiUserGetCurrentV3
	err = json.Unmarshal(b, &userApi)
	if err != nil {
		return nil, err
	}

	// Avatar
	var a Attachment
	err = database.DB.Model(&u).Association("Avatar").Find(&a)
	if err != nil {
		return nil, err
	}
	userApi.Avatar = a.Path

	// Roles
	var r []Role
	_ = database.DB.Model(&u).Association("Roles").Find(&r)
	for _, v := range r {
		userApi.Roles = append(userApi.Roles, v.Name)
	}

	if library.Config.App.Cache {
		b, err := json.Marshal(userApi)
		if err != nil {
			core.LogError.Output(err)
		} else {
			var m = make(map[string]string)
			m[strconv.Itoa(int(u.ID))] = string(b)
			library.Cache.SetHash(core.Cache_UserGetCurrent, m)
		}
	}

	return userApi, nil
}
func (this *User) user_GetCurrent_Cache(c *gin.Context) (interface{}, error) {
	id, err := this.User_GetUserIdByRequest(c)
	if err != nil {
		return nil, err
	}

	str, err := library.Cache.GetHash(core.Cache_UserGetCurrent, strconv.Itoa(int(id)))
	if err != nil {
		if err == redis.Nil {
			return this.user_GetCurrent_NotCache(c)
		}
		return nil, err
	}

	var userApi ApiUserGetCurrent
	err = json.Unmarshal([]byte(str), &userApi)
	if err != nil {
		library.Cache.DeleteHash(core.Cache_UserGetCurrent, strconv.Itoa(int(id)))
	}

	return userApi, nil
}


// 加密密码
func (this *User) User_BcryptPasswordV2(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}


// 根据token获取用户
func (this *User) User_GetUserByToken(token string) (u User, err error) {
	// 从Token获取ID
	id, err := library.Jwt_Token.GetIdByToken(token)
	if err != nil {
		return u, err
	}
	// 从Token获取username
	username, err  := library.Jwt_Token.GetIssuerByToken(token)
	if err != nil {
		return u, err
	}
	// 获取用户
	var data User
	result := database.DB.Where("id = ? AND username = ?", id, username).First(&data)
	if !(result.RowsAffected > 0) {
		return u, errors.New("没有在数据库中找到当前用户！")
	}

	return data, nil
}


// 根据请求信息获取用户
func (this *User) User_GetUserByRequest(c *gin.Context) (u User, err error) {
	token, err := core.GetToken(c)
	if err != nil {
		return u, err
	}
	u, err = this.User_GetUserByToken(token)
	if err != nil {
		return u, err
	}
	return u, err
}


// 根据请求信息获取用户id
func (this *User) User_GetUserIdByRequest(c *gin.Context) (id uint, err error) {
	token, err := core.GetToken(c)
	if err != nil {
		return 0, err
	}
	id, err = this.User_GetUserIdByToken(token)
	if err != nil {
		return 0, err
	}
	return id, err
}


// 根据token获取用户id
func (this *User) User_GetUserIdByToken(token string) (id uint, err error) {
	// 从Token获取ID
	id, err = library.Jwt_Token.GetIdByToken(token)
	if err != nil {
		return 0, err
	}

	return id, nil
}


// interface conversion User
// interface转换User
func (this *User) User_InterfaceToUser(i interface{}) (User, error) {
	var u User
	us, err := json.Marshal(i)
	if err != nil {
		core.LogError.Output("User_InterfaceToUser:对象转JSON失败! err:" + err.Error())
		return User{}, err
	}
	err = json.Unmarshal(us, &u)
	if err != nil {
		core.LogError.Output("User_InterfaceToUser:JSON转换对象失败! err:" + err.Error())
		return User{}, err
	}
	return u, nil
}
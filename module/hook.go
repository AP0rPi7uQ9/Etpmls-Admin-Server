package module

import (
	"github.com/gin-gonic/gin"
)

type Hook struct {

}


// User Create Hook
// 用户创建钩子
func (this *Hook) UserCreate(c *gin.Context, u interface{}) (err error) {
	// What you need to do when a user is created...
	// 用户创建时你需要做的事情
	// ...
	return err
}


// User Edit Hook
// 用户修改钩子
func (this *Hook) UserEdit(c *gin.Context, u interface{}) (err error) {
	// Things you need to do when users modify
	// 用户修改时你需要做的事情
	// ...
	return err
}


// User Delete Hook
// 用户删除钩子
func (this *Hook) UserDelete(c *gin.Context, u interface{}) (err error) {
	// Things you need to do when users delete
	// 用户删除时你需要做的事情
	// ...
	return err
}


// Role Create Hook
// 角色创建钩子
func (this *Hook) RoleCreate(c *gin.Context, r interface{}) (err error) {
	// What you need to do when a role is created...
	// 角色创建时你需要做的事情
	// ...
	return err
}


// User Edit Hook
// 角色修改钩子
func (this *Hook) RoleEdit(c *gin.Context, r interface{}) (err error) {
	// Things you need to do when role modify
	// 角色修改时你需要做的事情
	// ...
	return err
}


// User Delete Hook
// 角色删除钩子
func (this *Hook) RoleDelete(c *gin.Context, r interface{}) (err error) {
	// Things you need to do when role delete
	// 角色删除时你需要做的事情
	// ...
	return err
}


// Permission Create Hook
// 权限创建钩子
func (this *Hook) PermissionCreate(c *gin.Context, p interface{}) (err error) {
	// ...
	return err
}


// Permission Edit Hook
// 权限修改钩子
func (this *Hook) PermissionEdit(c *gin.Context, p interface{}) (err error) {
	// ...
	return err
}


// Permission Delete Hook
// 权限删除钩子
func (this *Hook) PermissionDelete(c *gin.Context, p interface{}) (err error) {
	// ...
	return err
}
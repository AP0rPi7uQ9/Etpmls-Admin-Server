package model

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"time"
)

type Menu struct {}


// Create Menu
// 创建菜单
type ApiMenuCreate struct {
	ID        uint `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Menu string `json:"menu" binding:"required"`
}
func (this *Menu) MenuCreate(j ApiMenuCreate) (error) {
	// Move files
	// 移动文件
	err := os.Rename("storage/menu/menu.json", "storage/menu/menu.json.bak")
	if err != nil {
		core.LogError.AutoOutputDebug("备份菜单文件失败！", err)
		return err
	}

	// Write file
	// 写入文件
	var s = []byte(j.Menu)
	err = ioutil.WriteFile("storage/menu/menu.json", s, 0666)
	if err != nil {
		core.LogError.AutoOutputDebug("写入菜单文件失败！", err)

		// 还原历史菜单
		err = os.Rename("storage/menu/menu.json.bak", "storage/menu/menu.json")
		if err != nil {
			core.LogError.AutoOutputDebug("备份菜单文件还原失败！", err)
		}

		return err
	}

	return nil
}


// Get all menu
// 获取全部菜单
func (this *Menu) MenuGetAll() (interface{}, error) {
	if library.Config.App.Cache {
		return this.menu_GetAll_Cache()
	} else {
		return this.menu_GetAll_NoCache()
	}
}
func (this *Menu) menu_GetAll_Cache() (interface{}, error) {
	// Get the menu from cache
	// 从缓存中获取menu
	ctx, err := library.Redis.Get(library.RedisCtx, core.Cache_MenuGetAll).Bytes()
	if err != nil {
		if err == redis.Nil {
			return this.menu_GetAll_NoCache()
		}
		return nil, err
	}

	var j interface{}
	err = json.Unmarshal(ctx, &j)
	if err != nil {
		core.LogError.Output(err)
		err2 := library.Redis.Del(library.RedisCtx, core.Cache_MenuGetAll).Err()
		if err2 != nil {
			core.LogError.Output(err2)
			return nil, err
		}
		return nil, err
	}
	return j, nil
}
func (this *Menu) menu_GetAll_NoCache() (interface{}, error) {
	ctx, err := ioutil.ReadFile("./storage/menu/menu.json")
	if err != nil {
		return nil, err
	}
	// Save menu
	// 储存菜单
	err = library.Redis.Set(library.RedisCtx, core.Cache_MenuGetAll, ctx, 0).Err()
	if err != nil {
		core.LogError.Output(err)
		err = library.Redis.Del(library.RedisCtx, core.Cache_MenuGetAll).Err()
		if err != nil {
			core.LogError.Output(err)
		}
	}

	var j interface{}
	err = json.Unmarshal(ctx, &j)
	if err != nil {
		return nil, err
	}

	return j, nil
}
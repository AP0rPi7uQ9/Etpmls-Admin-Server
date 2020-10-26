package model

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/utils"
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
		core.LogError.Output(utils.MessageWithLineNum("Failed to backup menu file!" + err.Error()))
		return err
	}

	// Write file
	// 写入文件
	var s = []byte(j.Menu)
	err = ioutil.WriteFile("storage/menu/menu.json", s, 0666)
	if err != nil {
		core.LogError.Output(utils.MessageWithLineNum("Failed to write menu file!" + err.Error()))

		// 还原历史菜单
		err2 := os.Rename("storage/menu/menu.json.bak", "storage/menu/menu.json")
		if err2 != nil {
			core.LogError.Output(utils.MessageWithLineNum("Failed to restore the backup menu file!" + err2.Error()))
		}

		return err
	}

	// Delete Cache
	// 删除缓存
	if library.Config.App.Cache {
		library.Cache.DeleteString(core.Cache_MenuGetAll)
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
	ctx, err := library.Cache.GetString(core.Cache_MenuGetAll)
	if err != nil {
		if err == redis.Nil {
			return this.menu_GetAll_NoCache()
		}
		core.LogError.Output(utils.MessageWithLineNum(err.Error()))
		return nil, err
	}

	var j interface{}
	err = json.Unmarshal([]byte(ctx), &j)
	if err != nil {
		core.LogError.Output(utils.MessageWithLineNum(err.Error()))
		library.Cache.DeleteString(core.Cache_MenuGetAll)
		return nil, err
	}
	return j, nil
}
func (this *Menu) menu_GetAll_NoCache() (interface{}, error) {
	ctx, err := ioutil.ReadFile("./storage/menu/menu.json")
	if err != nil {
		core.LogError.Output(utils.MessageWithLineNum(err.Error()))
		return nil, err
	}
	// Save menu
	// 储存菜单
	if library.Config.App.Cache {
		library.Cache.SetString(core.Cache_MenuGetAll, string(ctx), 0)
	}

	var j interface{}
	err = json.Unmarshal(ctx, &j)
	if err != nil {
		core.LogError.Output(utils.MessageWithLineNum(err.Error()))
		return nil, err
	}

	return j, nil
}
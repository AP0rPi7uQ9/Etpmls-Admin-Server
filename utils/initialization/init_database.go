package initialization

import (
	"Etpmls-Admin-Server/database"
	"Etpmls-Admin-Server/library"
	"Etpmls-Admin-Server/module"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"strings"
)

/*
	Initialize the database
	初始化数据库
	(Init Database)
*/
func InitDatabase()  {
	env, err := godotenv.Read("./.env")
	if err != nil {
		library.Log.Error(err)
		return
	}

	if _, ok := env["INIT_DATABASE"]; ok {
		if strings.ToUpper(env["INIT_DATABASE"]) == "TRUE" {
			InsertBasicDataToDatabase()
			module.InsertModuleDataToDatabase()
			env["INIT_DATABASE"] = "FALSE"
		}
	}


	err = godotenv.Write(env, "./.env")
	if err != nil {
		library.Log.Error(err)
		return
	}
}

func InsertBasicDataToDatabase()  {
	// Create Role
	role := database.Role{
		Name:        "Administrator",
		Remark: "系统管理员",
	}
	if err := database.DB.Debug().Create(&role).Error; err != nil {
		library.Log.Error("utils/initialization/init_database.go:", err)
	}


	// Create User
	user := database.User{
		Username: "admin",
		Password: "$2a$10$yNoJrsN7mrtHzUyvm6s8KOwHrnkkGmqcRJvcieQKItIfQNwyzqfMy",
		Roles: []database.Role{
			{
				Model:       gorm.Model{ID:1},
			},
		},
	}
	if err := database.DB.Debug().Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Create(&user).Error; err != nil {
		library.Log.Error("utils/initialization/init_database.go:", err)
	}

	// Create Permission
	permission := []database.Permission{
		{
			Name: "查看用户",
			Method: "GET",
			Path: "/api/*/user/getAll",
			Remark: "查看用户列表",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "创建用户",
			Method: "POST",
			Path: "/api/*/user/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "编辑用户",
			Method: "PUT",
			Path: "/api/*/user/edit",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "删除用户",
			Method: "DELETE",
			Path: "/api/*/user/delete",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "查看角色",
			Method: "GET",
			Path: "/api/*/role/getAll",
			Remark: "查看角色列表",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "创建角色",
			Method: "POST",
			Path: "/api/*/role/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "编辑角色",
			Method: "PUT",
			Path: "/api/*/role/edit",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "删除角色",
			Method: "DELETE",
			Path: "/api/*/role/delete",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "查看权限",
			Method: "GET",
			Path: "/api/*/permission/getAll",
			Remark: "查看权限列表",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "创建权限",
			Method: "POST",
			Path: "/api/*/permission/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "编辑权限",
			Method: "PUT",
			Path: "/api/*/permission/edit",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "删除权限",
			Method: "DELETE",
			Path: "/api/*/permission/delete",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "创建/编辑菜单",
			Method: "POST",
			Path: "/api/*/menu/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},

	}
	if err := database.DB.Debug().Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false).Create(&permission).Error; err != nil {
		library.Log.Error("utils/initialization/init_database.go:", err)
	}
}

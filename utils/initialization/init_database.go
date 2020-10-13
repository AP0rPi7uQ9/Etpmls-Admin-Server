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
		Remark: "System Administrator",
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
			Name: "View User",
			Method: "GET",
			Path: "/api/*/user/getAll",
			Remark: "View user list",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Create User",
			Method: "POST",
			Path: "/api/*/user/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Edit User",
			Method: "PUT",
			Path: "/api/*/user/edit",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Delete User",
			Method: "DELETE",
			Path: "/api/*/user/delete",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "View Role",
			Method: "GET",
			Path: "/api/*/role/getAll",
			Remark: "View role list",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Create Role",
			Method: "POST",
			Path: "/api/*/role/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Edit Role",
			Method: "PUT",
			Path: "/api/*/role/edit",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Delete Role",
			Method: "DELETE",
			Path: "/api/*/role/delete",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "View Permission",
			Method: "GET",
			Path: "/api/*/permission/getAll",
			Remark: "View permission list",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Create Permission",
			Method: "POST",
			Path: "/api/*/permission/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Edit Permission",
			Method: "PUT",
			Path: "/api/*/permission/edit",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Delete Permission",
			Method: "DELETE",
			Path: "/api/*/permission/delete",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Create/Edit Menu",
			Method: "POST",
			Path: "/api/*/menu/create",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Clear Cache",
			Method: "GET",
			Path: "/api/*/setting/clearCache",
			Roles: []database.Role{
				{
					Model:       gorm.Model{ID:1},
				},
			},
		},
		{
			Name: "Disk Cleanup",
			Method: "GET",
			Path: "/api/*/setting/diskCleanup",
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

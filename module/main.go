package module

import (
	"Etpmls-Admin-Server/core"
	"Etpmls-Admin-Server/library"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


func InitModule()  {
	initYaml()
	initDatabase()
	initHook()
}


// Initialization Yaml
// 初始化Yaml
func initYaml()  {
	var yamlPath string

	for _, v := range library.Config.Module.Name {
		if os.Getenv("DEBUG") == "TRUE" {
			yamlPath = "storage/config/" + v + "_debug.yaml"
		} else{
			yamlPath = "storage/config/" + v + ".yaml"
		}

		b, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			library.Instance_Logrus.Error("Failed to read the configuration file! Error:", err)
			continue
		}

		library.Config_Module[v] = make(map[interface{}]interface{})
		err = yaml.Unmarshal(b, library.Config_Module[v])
		if err != nil {
			library.Instance_Logrus.Error("Failed to unmarshal the configuration file! Error:", err)
			continue
		}
	}

	return

}


// Initialization Event
// 初始化事件
func initHook()  {
	// Register Event
	// 注册事件
	go core.Event.UserCreate(UserCreate)
	go core.Event.UserEdit(UserEdit)
	go core.Event.UserDelete(UserDelete)
	go core.Event.RoleCreate(RoleCreate)
	go core.Event.RoleEdit(RoleEdit)
	go core.Event.RoleDelete(RoleDelete)
	go core.Event.PermissionCreate(PermissionCreate)
	go core.Event.PermissionEdit(PermissionEdit)
	go core.Event.PermissionDelete(PermissionDelete)
}

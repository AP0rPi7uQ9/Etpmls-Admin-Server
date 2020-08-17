package module

import (
	"Etpmls-Admin-Server/library"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


var Module_Config = make(map[string]map[interface{}]interface{})

func InitModule()  {
	var yamlPath string

	for _, v := range library.Config.Module.Name {
		if os.Getenv("DEBUG") == "TRUE" {
			yamlPath = "storage/config/" + v + "_debug.yaml"
		} else{
			yamlPath = "storage/config/" + v + ".yaml"
		}

		b, err := ioutil.ReadFile(yamlPath)
		if err != nil {
			library.Log.Error("Failed to read the configuration file! Error:", err)
			continue
		}

		Module_Config[v] = make(map[interface{}]interface{})
		err = yaml.Unmarshal(b, Module_Config[v])
		if err != nil {
			library.Log.Error("Failed to unmarshal the configuration file! Error:", err)
			continue
		}
	}

	return

}
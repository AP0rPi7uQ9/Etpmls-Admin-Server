package module

import (
	"Etpmls-Admin-Server/library"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func InitModule()  {
	initYaml()
	initDatabase()
}

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
			library.Log.Error("Failed to read the configuration file! Error:", err)
			continue
		}

		library.Config_Module[v] = make(map[interface{}]interface{})
		err = yaml.Unmarshal(b, library.Config_Module[v])
		if err != nil {
			library.Log.Error("Failed to unmarshal the configuration file! Error:", err)
			continue
		}
	}

	return

}

// https://github.com/go-yaml/yaml

package library

import (
	"Etpmls-Admin-Server/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type configuration struct {
	App struct{
		Port string
		Captcha bool
		Register bool
		Key string
		Log struct {
			Level string
			Panic	int
			Fatal	int
			Error	int
			Warning	int
			Info	int
			Debug	int
			Trace	int
		}
		Api struct{
			UseHttpCode bool	`yaml:"use_http_code"`
			Field struct{
				Code string
				Message string
				Status string
				Data string
			}
			Pagination struct {
				Field struct {
					PageNo string `yaml:"page_no"`
					PageSize string `yaml:"page_size"`
					Count string
				}
			}
		}
	}
	Database struct{
		Host string
		Port string
		Name string
		User string
		Password string
		Prefix string
	}
	Redis struct{
		Address string
		Password string
		DB int
	}
	Module struct{
		Name []string
	}
}

var Config = configuration{}

var Config_Module = make(map[string]map[interface{}]interface{})

func init() {
	var yamlPath string

	if os.Getenv("DEBUG") == "TRUE" {
		yamlPath = "storage/config/app_debug.yaml"
	} else{
		yamlPath = "storage/config/app.yaml"
	}

	b, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		Log.Fatal("Failed to read the configuration file! Error:", err)
		return
	}

	err = yaml.Unmarshal(b, &Config)
	if err != nil {
		Log.Fatal("Failed to unmarshal the configuration file! Error:", err)
		return
	}

	if len(Config.App.Key) < 50 {
		Config.App.Key = utils.GenerateRandomString(50)

		out, err := yaml.Marshal(Config)
		if err != nil {
			Log.Fatal("配置文件解析成yaml格式失败！", err)
			return
		}

		err = ioutil.WriteFile(yamlPath, out, os.ModeAppend)
		if err != nil {
			Log.Fatal("写入yaml配置文件失败！", err)
			return
		}
	}


	return

}











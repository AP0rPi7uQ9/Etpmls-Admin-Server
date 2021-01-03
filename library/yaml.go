// https://github.com/go-yaml/yaml

package library

import (
	"Etpmls-Admin-Server/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type configuration struct {
	App struct{
		Port string
		Captcha bool
		Register bool
		Key string
		Cache bool
		ServiceDiscovery	bool	`yaml:"service-discovery"`
		TokenExpirationTime time.Duration	`yaml:"token-expiration-time"`
		UseHttpCode bool	`yaml:"use-http-code"`
		TimeZone string		`yaml:"time-zone"`
	}
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
	Database struct{
		Host string
		Port string
		Name string
		User string
		Password string
		Prefix string
	}
	Cache struct{
		Address string
		Password string
		DB int
	}
	ServiceDiscovery struct{
		Address string
		Service struct{
			Name string
			Address string
			Port int
			Tag []string
			CheckInterval string	`yaml:"check-interval"`
			CheckUrl string	`yaml:"check-url"`
		}
	}	`yaml:"service-discovery"`
	Field struct{
		Api struct{
			Code string
			Message string
			Status string
			Data string
		}
		Pagination struct {
			PageNo string `yaml:"page_no"`
			PageSize string `yaml:"page_size"`
			Count string
		}
	}
	Module struct{
		Name []string
	}
}

var Config = configuration{}

var Config_Module = make(map[string]map[interface{}]interface{})

func init_Yaml() {
	var yamlPath string

	if os.Getenv("DEBUG") == "TRUE" {
		yamlPath = "storage/config/app_debug.yaml"
	} else{
		yamlPath = "storage/config/app.yaml"
	}

	b, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		Instance_Logrus.Fatal("Failed to read the configuration file! Error:", err)
		return
	}

	err = yaml.Unmarshal(b, &Config)
	if err != nil {
		Instance_Logrus.Fatal("Failed to unmarshal the configuration file! Error:", err)
		return
	}

	if len(Config.App.Key) < 50 {
		Config.App.Key = utils.GenerateRandomString(50)

		out, err := yaml.Marshal(Config)
		if err != nil {
			Instance_Logrus.Fatal("配置文件解析成yaml格式失败！", err)
			return
		}

		err = ioutil.WriteFile(yamlPath, out, os.ModeAppend)
		if err != nil {
			Instance_Logrus.Fatal("写入yaml配置文件失败！", err)
			return
		}
	}


	return

}











package library

import (
	"github.com/google/uuid"
	Package_Consul "github.com/hashicorp/consul/api"
)

var (
	Instance_Consul *Package_Consul.Client
	service_Id string
)

func init_Consul()  {
	config := Package_Consul.DefaultConfig()
	config.Address = Config.ServiceDiscovery.Address

	// Establish connection
	// 建立连接
	var err error
	Instance_Consul, err = Package_Consul.NewClient(config)
	if err != nil {
		Instance_Logrus.Warning("Consul initialization failed.")
	} else {
		Instance_Logrus.Info("Consul initialized successfully.")
	}

	// Registration Service
	// 注册服务
	service_Id = uuid.New().String()
	c := NewConsul()
	_ = c.RegistrationService()
}

type consul struct {

}

func NewConsul() *consul {
	return &consul{}
}

// Registration Service
// 注册服务
func (this *consul) RegistrationService() error {
	r := Package_Consul.AgentServiceRegistration{
		ID:				   service_Id,
		Name:              Config.ServiceDiscovery.Service.Name,
		Tags:              Config.ServiceDiscovery.Service.Tag,
		Port:              Config.ServiceDiscovery.Service.Port,
		Address:           Config.ServiceDiscovery.Service.Address,
	}

	c := Package_Consul.AgentServiceCheck{
		Interval:                       Config.ServiceDiscovery.Service.CheckInterval,
		HTTP:                           Config.ServiceDiscovery.Service.CheckUrl,
	}

	r.Check = &c
	err := Instance_Consul.Agent().ServiceRegister(&r)
	if err != nil {
		Instance_Logrus.Error("Consul Service registration failed! Error:", err.Error())
		return err
	}

	return nil
}

// Cancel Service
// 取消服务
func (this *consul) CancelService() error {
	err := Instance_Consul.Agent().ServiceDeregister(service_Id)
	if err != nil {
		Instance_Logrus.Error("Cancel Consul service failed! Error:", err.Error())
		return err
	}
	return nil
}
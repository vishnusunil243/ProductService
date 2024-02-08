package servicediscovery

import (
	"fmt"
	"log"

	consulapi "github.com/hashicorp/consul/api"
)

const (
	port      = 8080
	serviceId = "product-service"
)

func RegisterService() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err.Error())
		return
	}
	addr := "localhost"
	fmt.Printf("address is %v", addr)

	registration := &consulapi.AgentServiceRegistration{
		ID:      serviceId,
		Name:    "product-server",
		Port:    port,
		Address: addr,
		Check: &consulapi.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d/%s", addr, port, serviceId),
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "10m",
		},
	}
	log.Printf(fmt.Sprintf("%s:%d/%s", addr, port, serviceId))
	regiErr := consul.Agent().ServiceRegister(registration)
	fmt.Println(regiErr)
	if regiErr != nil {
		log.Printf("failed to register service %s:%v", addr, port)
	} else {
		log.Printf("successfully registered services %s:%v", addr, port)
	}
}

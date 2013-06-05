package models

import (
	"github.com/skynetservices/skynet"
	"github.com/skynetservices/zkmanager"
	"time"
)

func GetRegions() []string {
	sm := zkmanager.NewZookeeperServiceManager("localhost:2181", 1*time.Second)
	return sm.ListRegions(skynet.ServiceQuery{})
}
func GetServices() []string {
	sm := zkmanager.NewZookeeperServiceManager("localhost:2181", 1*time.Second)
	return sm.ListServices(skynet.ServiceQuery{})
}
func GetInstances() []skynet.ServiceInfo {
	sm := zkmanager.NewZookeeperServiceManager("localhost:2181", 1*time.Second)
	return sm.ListInstances(skynet.ServiceQuery{})
}
func GetHosts() []string {
	sm := zkmanager.NewZookeeperServiceManager("localhost:2181", 1*time.Second)
	return sm.ListHosts(skynet.ServiceQuery{})
}

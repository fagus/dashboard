package controllers

import (
	"github.com/robfig/revel"
	"github.com/skynetservices/dashboard/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	regions := models.GetRegions()
	regionCount := len(regions)
	services := models.GetServices()
	serviceCount := len(services)
	instances := models.GetInstances()
	revel.ERROR.Println(instances)
	instanceCount := len(instances)
	hosts := models.GetHosts()
	hostCount := len(hosts)
	return c.Render(regions, regionCount, services, serviceCount, instances, instanceCount, hosts, hostCount)
}

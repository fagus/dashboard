package controllers

import (
	"github.com/robfig/revel"
	"github.com/skynetservices/dashboard/app/models"
)

type Regions struct {
	*revel.Controller
}

func (c Regions) Index() revel.Result {
	regions := models.GetRegions()
	return c.Render(regions)
}

func (c Regions) Nodes(id string) revel.Result {
	revel.ERROR.Println(id)
	hosts := models.GetHosts()
	return c.Render(id, hosts)
}
func (c Regions) Instances(id string) revel.Result {
	revel.ERROR.Println(id)
	instances := models.GetInstances()
	return c.Render(id, instances)
}
func (c Regions) Services(id string) revel.Result {
	revel.ERROR.Println(id)
	services := models.GetServices()
	return c.Render(id, services)
}

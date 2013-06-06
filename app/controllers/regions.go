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

	return c.Render()
}
func (c Regions) Instances(id string) revel.Result {

	return c.Render()
}
func (c Regions) Services(id string) revel.Result {

	return c.Render()
}

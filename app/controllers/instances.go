package controllers

import (
	"github.com/robfig/revel"
	"github.com/skynetservices/dashboard/app/models"
)

type Instances struct {
	*revel.Controller
}

func (c Instances) Index() revel.Result {

	instances := models.GetInstances()
	return c.Render(instances)
}

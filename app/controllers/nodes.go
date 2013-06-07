package controllers

import (
	"github.com/robfig/revel"
	"github.com/skynetservices/dashboard/app/models"
)

type Nodes struct {
	*revel.Controller
}

func (c Nodes) Index() revel.Result {

	hosts := models.GetHosts()
	return c.Render(hosts)
}

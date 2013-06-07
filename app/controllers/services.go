package controllers

import (
	"github.com/robfig/revel"
	"github.com/skynetservices/dashboard/app/models"
)

type Services struct {
	*revel.Controller
}

func (c Services) Index() revel.Result {

	services := models.GetServices()
	return c.Render(services)
}

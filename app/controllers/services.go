package controllers

import (
	"github.com/robfig/revel"

)

type Services struct {
	*revel.Controller
}

func (c Services) Index() revel.Result {

	return c.Render()
}

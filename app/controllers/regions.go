package controllers

import (
	"github.com/robfig/revel"

)

type Regions struct {
	*revel.Controller
}

func (c Regions) Index() revel.Result {

	return c.Render()
}

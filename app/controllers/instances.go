package controllers

import (
	"github.com/robfig/revel"

)

type Instances struct {
	*revel.Controller
}

func (c Instances) Index() revel.Result {

	return c.Render()
}

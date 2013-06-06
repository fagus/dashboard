package controllers

import (
	"github.com/robfig/revel"

)

type Nodes struct {
	*revel.Controller
}

func (c Nodes) Index() revel.Result {

	return c.Render()
}

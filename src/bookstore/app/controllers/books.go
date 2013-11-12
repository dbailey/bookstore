package controllers

import "github.com/robfig/revel"

type Books struct {
	*revel.Controller
}

func (c Books) Index() revel.Result {
	return c.RenderJson("")
}

func (c Books) Create() revel.Result {
	return c.RenderJson("")
}

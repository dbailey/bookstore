package controllers

import (
	"encoding/json"
	"github.com/robfig/revel"
)

type Books struct {
	*revel.Controller
}

func (c Books) Index() revel.Result {
	return c.RenderJson("")
}

func (c Books) Create() revel.Result {
	decoder := json.NewDecoder(c.Request.Body)
	var jsonBody map[string]interface{}
	_ = decoder.Decode(&jsonBody)
	return c.RenderJson(jsonBody)
}

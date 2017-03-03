package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetFile() {
	path := c.Ctx.Input.Param(":path")
	ext := c.Ctx.Input.Param(":ext")
	data, err := ioutil.ReadFile("files/" + path + "." + ext)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(404)
		return
	}
	c.Ctx.Output.ContentType(ext)
	c.Ctx.Output.Body(data)
	return
}

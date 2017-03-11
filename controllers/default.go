package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
)

type MainController struct {
	beego.Controller
}

func init() {
	if err := os.Mkdir("files", 0755); err != nil {
		beego.Info("file folder not maked!")
	}
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

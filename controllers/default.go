package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetVideoFile() {
	beego.Info("GetVideoFile")
	directory := string([]byte(c.Ctx.Input.Param(":directory")))
	key := string([]byte(c.Ctx.Input.Param(":key")))
	file := string([]byte(c.Ctx.Input.Param(":file")))
	data, err := ioutil.ReadFile("files/" + directory + "/" + key + "/" + file)
	if err != nil {
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/mp4")
	c.Ctx.Output.Body(data)
	return
}

func (c *MainController) GetVideoThumbnail() {
	beego.Info("GetVideoThumbnail")
	directory := string([]byte(c.Ctx.Input.Param(":directory")))
	key := string([]byte(c.Ctx.Input.Param(":key")))
	file := string([]byte(c.Ctx.Input.Param(":file")))
	data, err := ioutil.ReadFile("files/" + directory + "/" + key + "/img/" + file)
	if err != nil {
		return
	}
	c.Ctx.Output.Header("Content-Type", "application/img")
	c.Ctx.Output.Body(data)
	return
}

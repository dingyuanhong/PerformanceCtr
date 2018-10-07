package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	fmt.Println("MainController");
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}


type RootController struct {
	beego.Controller
}

func (c *RootController) Get() {
	c.TplName = "index.html"
	fmt.Println("root");
}
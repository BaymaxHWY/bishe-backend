package controllers

import (
	"bishe/backend/models"
	"fmt"
	"github.com/astaxie/beego"
)
// position API
type PositionController struct {
	beego.Controller
}

func (p *PositionController) Prepare() {
	p.Data["language"] = p.Ctx.Input.Param(":language")
}
// @Title GetCity
// @Description 获得所有城市数量信息
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key [get]
func (p *PositionController) GetCity() {
	positionModel := models.PositionModel{}
	res := positionModel.GetCity(fmt.Sprintf("%v", p.Data["language"]))
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PositionController) GetSalary() {
	positionModel := models.PositionModel{}
	res := positionModel.GetSalary(fmt.Sprintf("%v", p.Data["language"]))
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PositionController) GetCompany() {
	positionModel := models.PositionModel{}
	res := positionModel.GetCompany(fmt.Sprintf("%v", p.Data["language"]))
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PositionController) GetWorkYear() {
	positionModel := models.PositionModel{}
	res := positionModel.GetWorkYear(fmt.Sprintf("%v", p.Data["language"]))
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PositionController) GetEducation() {
	positionModel := models.PositionModel{}
	res := positionModel.GetEducation(fmt.Sprintf("%v", p.Data["language"]))
	p.Data["json"] = &res
	p.ServeJSON()
}

func (p *PositionController) GetLanguage() {
	positionModel := models.PositionModel{}
	res := positionModel.GetLanguage()
	p.Data["json"] = &res
	p.ServeJSON()
}
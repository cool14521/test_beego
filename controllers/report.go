package controllers

import (
	"encoding/json"
	"finance_report/models"

	"github.com/astaxie/beego"
)

//ReportController 报表类型
type ReportController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Summary create object
// @Param	body		body 	models.Report	true		"The object content"
// @Success 200 {string} models.Report.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ReportController) Post() {
	var ob models.Report
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

//Get 获取报表
// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Report
// @Failure 403 :objectId is empty
// @router /:objectID  [get]
func (o *ReportController) Get() {
	objectID := o.Ctx.Input.Param(":objectID")
	if objectID != "" {
		ob, err := models.GetOne(objectID)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Report
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ReportController) GetAll() {
	obs := models.GetAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Report	true		"The body"
// @Success 200 {object} models.Report
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *ReportController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Report
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ReportController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

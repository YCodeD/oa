package controllers

import (
	"encoding/json"
	"oa/models"
)

// 车辆申请
type VehicleController struct {
	BaseController
}

type ApplyRequest struct {
	// 申请人
	Name string `json:"applicant"`
	// 申请日期
	ApplyDate string `json:"applydate"`
	// 用车时间
	UseTime string `json:"usetime"`
	// 目的地
	Destination string `json:"destination"`
	// 申请记录
	Record string `json:"record"`
}

// @Title 添加用车申请
// @Description 添加用车申请
// @Param body body controllers.ApplyRequest "申请请求"
// @Success 200
// @Failure 403
// @router /application [post]
func (v *VehicleController) ApplyVehicle() {
	vr := &ApplyRequest{}
	json.Unmarshal(v.Ctx.Input.RequestBody, &vr)
	models.CreateVehicle(&models.Vehicle{
		Name:        vr.Name,
		ApplyDate:   vr.ApplyDate,
		UseTime:     vr.UseTime,
		Destination: vr.Destination,
		Record:      vr.Record,
	})
}

// @Title 获取用户申请记录
// @Description 获取用户申请记录
// @Param name path string "申请人"
// @Success 200
// @Failure 403
// @router /:name [post]
func (v *VehicleController) GetVehicleRecords() {
	name := v.Ctx.Input.Param(":name")
	res := models.FetchVehicleByName(name)
	v.Response(res)
}

// @Title 通过车辆申请
// @Description 通过车辆申请
// @Param id path int "车辆申请id"
// @Success 200
// @Failure 403
// @router /pass/:id [put]
func (v *VehicleController) PassVehicle() {
	id := v.GetPathInt("id")
	vehicle := models.FetchVehicleById(id)
	vehicle.Approval = 1
	res := models.UpdateVehicle(&vehicle)
	v.Response(res)
}

// @Title 拒绝车辆申请
// @Description 拒绝车辆申请
// @Param id path int "车辆申请id"
// @Success 200
// @Failure 403
// @router /reject/:id [put]
func (v *VehicleController) RejectVehicle() {
	id := v.GetPathInt("id")
	vehicle := models.FetchVehicleById(id)
	vehicle.Approval = 2
	res := models.UpdateVehicle(&vehicle)
	v.Response(res)
}

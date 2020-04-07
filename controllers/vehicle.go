package controllers

import (
	"encoding/json"

	"oa/models"
)

// VehicleController 车辆申请
type VehicleController struct {
	BaseController
}

// ApplyRequest 添加请求
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

// ApplyVehicle 添加方法
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

// GetVehicleRecords 获取方法
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

// PassVehicle 通过方法
// @Title 通过车辆申请
// @Description 通过车辆申请
// @Param id path int "车辆申请id"
// @Success 200
// @Failure 403
// @router /pass/:id [put]
func (v *VehicleController) PassVehicle() {
	id := v.GetPathInt("id")
	vehicle := models.FetchVehicleById(id)
	vehicle.Status = 1
	res := models.UpdateVehicle(&vehicle)
	v.Response(res)
}

// RejectVehicle 拒绝方法
// @Title 拒绝车辆申请
// @Description 拒绝车辆申请
// @Param id path int "车辆申请id"
// @Success 200
// @Failure 403
// @router /reject/:id [put]
func (v *VehicleController) RejectVehicle() {
	id := v.GetPathInt("id")
	vehicle := models.FetchVehicleById(id)
	vehicle.Status = 2
	res := models.UpdateVehicle(&vehicle)
	v.Response(res)
}

// GetVehicleByOption 获取方法
// @Title 通过选项获取车辆申请记录
// @Description 通过选项获取车辆申请记录
// @Param option path int false "0-未审批 1-通过 2-未通过"
// @Success 200
// @Failure 403
// @router /:option [get]
func (v *VehicleController) GetVehicleByOption() {
	option := v.GetPathInt("option")
	res := models.FetchVehicleByOption(option)
	v.Response(res)
}

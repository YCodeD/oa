package controllers

import (
	"encoding/json"

	"oa/models"
)

// TravelController 出游申请
type TravelController struct {
	BaseController
}

// TravelRequest 添加请求体
type TravelRequest struct {
	Name string `json:"name"`
	// 目的地
	Destination string `json:"destination"`
	//出游时间
	TravelTime string `json:"travel_time"`
	// 申请日期
	ApplyDate string `json:"apply_date"`
	// 人员名单
	StaffList string `json:"staff_list"`
}

// AddTravel 添加方法
// @Title 添加出游申请
// @Description 添加出游申请
// @Param body body controllers.TravelRequest "出游请求"
// @Success 200
// @Failure 403
// @router /addtravel [post]
func (t *TravelController) AddTravel() {
	req := &TravelRequest{}
	json.Unmarshal(t.Ctx.Input.RequestBody, &req)
	res := models.CreateTravel(&models.Travel{
		Name:        req.Name,
		Destination: req.Destination,
		TravelTime:  req.TravelTime,
		ApplyDate:   req.ApplyDate,
		StaffList:   req.StaffList,
	})

	t.Response(res)
}

// PassTravel 通过方法
// @Title 通过出游申请
// @Description 通过出游申请
// @Param id path int false "出游申请id"
// @Success 200
// @Failure 403
// @router /pass/:id [put]
func (t *TravelController) PassTravel() {
	id := t.GetPathInt("id")
	res := models.FetchTravelById(id)
	res.Status = 1
	err := models.UpdateTravel(&res)
	if err != nil {
		t.ResponseError(err)
	}
}

// RejectTravel 拒绝方法
// @Title 拒绝出游申请
// @Description 拒绝出游申请
// @Param id path int false "出游申请id"
// @Success 200
// @Failure 403
// @router /reject/:id [put]
func (t *TravelController) RejectTravel() {
	id := t.GetPathInt("id")
	res := models.FetchTravelById(id)
	res.Status = 2
	err := models.UpdateTravel(&res)
	if err != nil {
		t.ResponseError(err)
	}
}

// GetTravelByOption 获取方法
// @Title 通过选项获取出游申请
// @Description 通过选项获取出游申请
// @Param option path int false "0-未审批 1-通过 2-未通过"
// @Success 200
// @Failure 403
// @router /:option [get]
func (t *TravelController) GetTravelByOption() {
	option := t.GetPathInt("option")
	res := models.FetchTravelByOption(option)
	t.Response(res)
}

package controllers

import (
	"encoding/json"
	"oa/models"
)

// 出游申请
type TravelController struct {
	BaseController
}

// 出游请求
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

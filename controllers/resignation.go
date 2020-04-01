package controllers

import (
	"encoding/json"
	"oa/models"
)

// 离职
type ResignationController struct {
	BaseController
}

// Reignation 添加离职请求
type ResignationRequest struct {
	// 申请人
	Name string `json:"name"`
	// 申请日期
	Date string `json:"date"`
	// 离职时间
	ResignationDate string `json:"resignationtime"`
	// 交接人
	HandoverPerson string `json:"handover_person"`
	// 交接事项
	HandoverWork string `json:"handover_work"`
	// 离职原因
	Reason string `json:"reason"`
}

// @Title 添加离职申请
// @Description 离职申请
// @Param body body controllers.ResignationRequest "离职请求"
// @Success 200
// @Failure 403
// @router /application [post]
func (r *ResignationController) AddApplication() {
	req := &ResignationRequest{}
	json.Unmarshal(r.Ctx.Input.RequestBody, &req)
	err := models.CreateResignation(&models.Resignation{
		ApplicantName:   req.Name,
		ApplicationDate: req.Date,
		ResignationDate: req.ResignationDate,
		HandoverPerson:  req.HandoverPerson,
		HandoverWork:    req.HandoverWork,
		Reason:          req.Reason,
	})
	if err != nil {
		r.ResponseError(err)
	}

}

// @Title 通过离职申请
// @Description 通过离职申请
// @Param id path int true "离职申请id"
// @Success 200
// @Failure 403
// @router /pass/:id [put]
func (r *ResignationController) PassResignation () {
	id := r.GetPathInt("id")
	resign := models.FetchResignById(id)
	resign.Status = 1
	res := models.UpdateResignation(resign)
	r.Response(res)
}

// @Title 拒绝离职申请
// @Description 拒绝离职申请
// @Param id path int true "离职申请id"
// @Success 200
// @Failure 403
// @router /reject/:id [put]
func (r *ResignationController) RejectResignation () {
	id := r.GetPathInt("id")
	resign := models.FetchResignById(id)
	resign.Status = 2
	res := models.UpdateResignation(resign)
	r.Response(res)
}

// @Title 通过选项获取离职申请
// @Description 通过选项获取离职申请
// @Param option path int false "0-未审核 1-通过 2-未通过"
// @Success 200
// @Failure 403
// @router /:option [get]
func (r *ResignationController) GetResignationByOption() {
	option := r.GetPathInt("option")
	res := models.FetchResignationByOption(option)
	r.Response(res)
}
package controllers
//
//import (
//	"encoding/json"
//	"oa/models"
//)
//
//// 经费报销
//type ReimbursementController struct {
//	BaseController
//}
//
//type ReimRequest struct {
//	// 申请人
//	Applicant string `json:"applicant"`
//	// 金额
//	Sum int `json:"sum"`
//	// 用途
//	UseFor string `json:"use_for"`
//	// 申请日期
//	Date string `json:"date"`
//	// 经费记录
//	Record string `json:"record"`
//}
//
//// @Title 添加经费申请
//// @Description 添加经费申请
//// @Param body body controllers.ReimRequest "申请请求"
//// @Success 200
//// @Failure 403
//// @router /add [post]
//func (r *ReimbursementController) AddReim() {
//	req := &ReimRequest{}
//	json.Unmarshal(r.Ctx.Input.RequestBody, &req)
//	err := models.CreateReimbursement(&models.Reimbursement{
//		Applicant: req.Applicant,
//		Sum:       req.Sum,
//		UseFor:    req.UseFor,
//		Date:      req.Date,
//		Record:    req.Record,
//	})
//	if err != nil {
//		r.ResponseError(err)
//	}
//}

//// @Title 获取报销申请
//// @Description 获取审批流程的报销申请
//// @Param option path int true "审批选项 0-未审批  1-已通过  2-未通过"
//// @Success 200
//// @Failure 403
//// @router /:option [get]
//func (r *ReimbursementController) GetByOption() {
//	option := r.GetPathInt("option")
//	fmt.Printf("option: %#v \n", option)
//	res := models.FetchReimByOption(option)
//	r.Response(res)
//}

//// @Title 获取用户报销记录
//// @Description 获取报销记录
//// @Param name path string "申请人"
//// @Success 200
//// @Failure 403
//// @router /:name [get]
//func (r *ReimbursementController) GetReimRecord() {
//	name := r.Ctx.Input.Param(":name")
//	res := models.FetchReimByName(name)
//	r.Response(res)
//}

//// @Title 获取所有报销记录
//// @Description 获取报销记录
//// @Success 200
//// @Failure 403
//// @router /allReims [get]
//func (r *ReimbursementController) GetAllReims() {
//	res := models.FetchAllReimRecords()
//	r.Response(res)
//}

//// @Title 通过报销申请
//// @Description 通过报销申请
//// @Param id path int true "通过申请id"
//// @Success 200
//// @Failure 403
//// @router /pass/:id [put]
//func (r *ReimbursementController) PassReimbursement() {
//	id := r.GetPathInt("id")
//	reim := models.FetchReimById(id)
//	reim.Status = 1
//	res := models.UpdateReimbursement(reim)
//	r.Response(res)
//}
//
//// @Title 拒绝报销申请
//// @Description 拒绝报销申请
//// @Param id path int true "拒绝申请id"
//// @Success 200
//// @Failure 403
//// @router /reject/:id [put]
//func (r *ReimbursementController) RejectReimbursement() {
//	id := r.GetPathInt("id")
//	reim := models.FetchReimById(id)
//	reim.Status = 2
//	res := models.UpdateReimbursement(reim)
//	r.Response(res)
//}
//
//

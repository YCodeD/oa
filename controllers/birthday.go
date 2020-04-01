package controllers

import (
	"encoding/json"
	"oa/models"
)

// 庆生申请
type BirthdayController struct {
	BaseController
}

type BirthdayRequest struct {
	Name        string `json:"name"`
	Date        string `json:"date"`
	Destination string `json:"destination"`
}

// @Title 添加庆生申请
// @Description 添加庆生申请
// @Param body body controllers.BirthdayRequest "庆生请求"
// @Success 200
// @Failure 403
// @router /addbirthday [post]
func (b *BirthdayController) AddBirthday() {
	req := &BirthdayRequest{}
	json.Unmarshal(b.Ctx.Input.RequestBody, &req)
	res := models.CreateBirthday(&models.Birthday{
		Name:        req.Name,
		Date:        req.Date,
		Destination: req.Destination,
	})

	b.Response(res)
}

// @Title 通过庆生申请
// @Description 通过庆生申请
// @Param id path int true "庆生申请id"
// @Success 200
// @Failure 403
// @router /pass/:id [put]
func (b *BirthdayController) PassBirthday() {
	id := b.GetPathInt("id")
	birth := models.FetchBirthdayById(id)
	birth.Status = 1
	res := models.UpdateBirthday(birth)
	b.Response(res)
}

// @Title 拒绝庆生申请
// @Description 拒绝庆生申请
// @Param id path int true "庆生申请id"
// @Success 200
// @Failure 403
// @router /reject/:id [put]
func (b *BirthdayController) RejectBirthday() {
	id := b.GetPathInt("id")
	birth := models.FetchBirthdayById(id)
	birth.Status = 2
	res := models.UpdateBirthday(birth)
	b.Response(res)
}

// @Title 通过选项获取庆生申请
// @Description 通过选项获取庆生申请
// @Param option path int false "0-未审批 1-通过 2-未通过"
// @Seccess 200
// @Failure 403
// @router /:option [get]
func (b *BirthdayController) GetBirthByOption() {
	option := b.GetPathInt("option")
	res := models.FetchBirthdayByOption(option)
	b.Response(res)
}

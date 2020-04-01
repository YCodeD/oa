package models

import (
	"errors"
	"time"
)

type Resignation struct {
	ID int `json:"id"`
	// 申请人
	ApplicantName string `json:"applicant_name"`
	// 申请日期
	ApplicationDate string `json:"application_date"`
	// 离职时间
	ResignationDate string `json:"resignation_time"`
	// 离职原因
	Reason string `json:"reason"`
	// 交接人
	HandoverPerson string `json:"handover_person"`
	// 交接事项
	HandoverWork string `json:"handover_work"`
	// 审批 0-未审批  1-已审批  2-未通过  3-已通过
	Status int `json:"status" gorm:"default:0"`
}

// CreateResignation 添加离职申请
func CreateResignation(cr *Resignation) error {
	_, err1 := time.ParseInLocation("2006-01-02", cr.ApplicationDate, time.Local)
	_, err2 := time.ParseInLocation("2006-01-02", cr.ResignationDate, time.Local)
	if err2 != nil || err1 != nil {
		return errors.New("ResignationDate or ApplicationDate必须是2006-01-02的时间格式")
	}
	db.Create(&cr)
	return nil
}

// FetchAllResignation 获取所有离职申请
func FetchAllResignation() (res []Resignation) {
	db.Find(&res)
	return res
}

// FetchResignationById 通过id获取离职申请
func FetchResignById(id int) *Resignation {
	res := Resignation{}
	db.Where("id=?", id).First(&res)
	return &res
}

// UpdateResignation 更新离职申请
func UpdateResignation(ur *Resignation) Resignation {
	db.Save(&ur)
	return *ur
}

// FetchResignationByOption 通过选项获取离职申请
func FetchResignationByOption(option int) []Resignation {
	res := []Resignation{}
	db.Where("status=?", option).Find(&res)
	return res
}
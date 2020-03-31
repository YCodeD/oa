package models

import (
	"errors"
	"fmt"
	"time"
)

type Reimbursement struct {
	ID int `json:"id"`
	// 申请人
	Applicant string `json:"applicant"`
	// 金额
	Sum int `json:"sum"`
	// 用途
	UseFor string `json:"use_for"`
	// 申请日期
	Date string `json:"date"`
	// 经费记录
	Record string `json:"record"`
	// 审批 0-未审批  1-已通过  2-未通过
	Status int `json:"status" gorm:"default:0"`
}


// CreateReimbursement 添加报销申请
func CreateReimbursement(cr *Reimbursement) error {
	_, err := time.ParseInLocation("2006-01-02", cr.Date, time.Local)
	if err != nil {
		return errors.New("Date必须是2006-01-02的时间格式")
	}
	db.Create(&cr)
	return nil
}

// FetchAllReimRecords 获取所有报销记录
func FetchAllReimRecords() []Reimbursement {
	res := []Reimbursement{}
	db.Find(&res)
	return res
}

// FetchReimById 通过id获取用户报销记录
func FetchReimById(id int) *Reimbursement {
	res := Reimbursement{}
	db.Where("id=?", id).First(&res)
	//fmt.Printf(" by id %v \n", res)
	return &res
}

// UpdateReimbursement 更新报销记录
func UpdateReimbursement(ur *Reimbursement) Reimbursement {
	db.Save(&ur)
	//fmt.Printf("update %v", ur)
	return *ur
}

// FetchReimByOption 通过选项获取申请
func FetchReimByOption(option int) []Reimbursement {
	res := []Reimbursement{}
	fmt.Printf("FetchByOption  \n")
	db.Where("status=?", option).Find(&res)
	return res
}

package models

import "fmt"

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
	Approval int `json:"approval" gorm:"default:0"`
}

// CreateReimbursement 添加报销申请
func CreateReimbursement(cr *Reimbursement) Reimbursement {
	db.Create(&cr)
	return *cr
}

// UpdateReimbursement 更新报销记录
func UpdateReimbursement(ur *Reimbursement) Reimbursement {
	db.Save(&ur)
	fmt.Printf("update %v", ur)
	return *ur
}

// FetchReimByName 获取用户报销记录
func FetchReimByName(name string) []Reimbursement {
	res := []Reimbursement{}
	db.Where("applicant=?", name).Find(&res)
	return res
}

// FetchReimById 通过id获取用户报销记录
func FetchReimById(id int) *Reimbursement {
	res := Reimbursement{}
	db.Where("id=?", id).First(&res)
	fmt.Printf(" by id %v \n", res)
	return &res
}

// FetchAllReimRecords 获取所有报销记录
func FetchAllReimRecords() []Reimbursement {
	res := []Reimbursement{}
	db.Find(&res)
	return res
}

// FetchByOption 通过选项获取申请
func FetchByOption(option int) []Reimbursement {
	res := []Reimbursement{}
	db.Where("approval=?", option).Find(&res)
	return res
}

package models

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
func CreateResignation(res *Resignation) Resignation {
	//ad, _ := time.ParseInLocation("2006-01-02", res.ApplicationDate, time.Local)
	//et, _ := time.ParseInLocation("2006-01-02", res.EntryTime, time.Local)
	//rt, _ := time.ParseInLocation("2006-01-02", res.ResignationTime, time.Local)

	//tx := db.Begin()
	//err := tx.Create(&res).Error
	//if err != nil {
	//	fmt.Printf("%v \n", err)
	//	tx.Rollback()
	//}
	//tx.Commit()
	db.Create(&res)
	return *res
}

func FetchAllResignation() (res []Resignation) {
	db.Find(&res)
	return res
}

func FetchResignByName(name string) (res []Resignation) {
	db.Where("applicant_name=?", name).Find(&res)
	return res
}

func FetchResignById(id int) (res Resignation) {
	db.Where("id=?", id).First(&res)
	return res
}

func UpdateResignation(ur Resignation) error {
	err := db.Update(&ur).Error
	if err != nil {
		return err
	}
	return nil
}
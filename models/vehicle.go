package models

type Vehicle struct {
	ID int `json:"id"`
	// 申请人
	Name string `json:"applicant"`
	// 申请日期
	ApplyDate string `json:"apply_date"`
	// 用车时间
	UseTime string `json:"use_time"`
	// 目的地
	Destination string `json:"destination"`
	// 申请记录
	Record string `json:"record"`
	// 审批 0-未审批  1-已通过  2-未通过
	Approval int `json:"approval" gorm:"default:0"`
}

func CreateVehicle(vehicle *Vehicle) {
	db.Create(&vehicle)
}

func UpdateVehicle(uv *Vehicle) Vehicle {
	db.Save(&uv)
	return *uv
}

func FetchVehicleByName(name string) []Vehicle {
	res := []Vehicle{}
	db.Where("name=?", name).Find(&res)
	return res
}

func FetchVehicleById(id int) Vehicle {
	res := Vehicle{}
	db.Where("id=?", id).First(&res)
	return res
}

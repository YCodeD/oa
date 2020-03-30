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
	Status int `json:"status" gorm:"default:0"`
}

// CreateVehicle 添加用车申请
func CreateVehicle(vehicle *Vehicle) {
	db.Create(&vehicle)
}

// UpdateVehicle 更新用车申请
func UpdateVehicle(uv *Vehicle) Vehicle {
	db.Save(&uv)
	return *uv
}

// FetchVehicleByName 通过名字获取申请记录
func FetchVehicleByName(name string) []Vehicle {
	res := []Vehicle{}
	db.Where("name=?", name).Find(&res)
	return res
}

// FetchVehicleById 通过id获取用车申请
func FetchVehicleById(id int) Vehicle {
	res := Vehicle{}
	db.Where("id=?", id).First(&res)
	return res
}

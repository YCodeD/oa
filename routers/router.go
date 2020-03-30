// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"oa/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		// 公告专区
		beego.NSNamespace("/Announcement",
			beego.NSInclude(
				&controllers.AnnouncementController{},
			),
		),
		// 锻炼小组
		beego.NSNamespace("/exercise",
			beego.NSInclude(
				&controllers.ExerciseController{},
			),
		),
		// 离职申请
		beego.NSNamespace("/resignation",
			beego.NSInclude(
				&controllers.ResignationController{},
			),
		),
		// 经费报销
		beego.NSNamespace("/reimbursement",
			beego.NSInclude(
				&controllers.ReimbursementController{},
			),
		),
		// 车辆申请
		beego.NSNamespace("/vehicle",
			beego.NSInclude(
				&controllers.VehicleController{},
			),
		),
		// 出游申请
		beego.NSNamespace("/travel",
			beego.NSInclude(
				&controllers.TravelController{},
			),
		),
		// 庆生申请
		beego.NSNamespace("/birthday",
			beego.NSInclude(
				&controllers.BirthdayController{},
			),
		),
		// 教育培训申请
		beego.NSNamespace("/training",
			beego.NSInclude(
				&controllers.TrainingController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
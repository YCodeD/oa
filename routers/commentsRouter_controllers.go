package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "UpdateAnnouncement",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "GetAnnouncement",
            Router: `/?:aid([1-9][0-9]*)`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "AddAnnouncement",
            Router: `/addannouncement`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "DeleteAnnouncement",
            Router: `/delete/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "GetAllAnnouncement",
            Router: `/getall`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "GetReadAnnouncement",
            Router: `/read`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:AnnouncementController"] = append(beego.GlobalControllerRouter["oa/controllers:AnnouncementController"],
        beego.ControllerComments{
            Method: "GetUnreadAnnouncement",
            Router: `/unread`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:BirthdayController"] = append(beego.GlobalControllerRouter["oa/controllers:BirthdayController"],
        beego.ControllerComments{
            Method: "AddBirthday",
            Router: `/addbirthday`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:BirthdayController"] = append(beego.GlobalControllerRouter["oa/controllers:BirthdayController"],
        beego.ControllerComments{
            Method: "PassBirthday",
            Router: `/pass/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:BirthdayController"] = append(beego.GlobalControllerRouter["oa/controllers:BirthdayController"],
        beego.ControllerComments{
            Method: "RejectBirthday",
            Router: `/reject/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["oa/controllers:ExerciseController"],
        beego.ControllerComments{
            Method: "CreateTeam",
            Router: `/createteam`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["oa/controllers:ExerciseController"],
        beego.ControllerComments{
            Method: "ExerciseClock",
            Router: `/exerciseclock`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["oa/controllers:ExerciseController"],
        beego.ControllerComments{
            Method: "GetRecords",
            Router: `/getrecords/:name`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["oa/controllers:ExerciseController"],
        beego.ControllerComments{
            Method: "GetGroupList",
            Router: `/grouplist`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["oa/controllers:ExerciseController"],
        beego.ControllerComments{
            Method: "JionTeam",
            Router: `/join/:uid/:tid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ExerciseController"] = append(beego.GlobalControllerRouter["oa/controllers:ExerciseController"],
        beego.ControllerComments{
            Method: "OutTeam",
            Router: `/out/:uid/:tid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ReimbursementController"] = append(beego.GlobalControllerRouter["oa/controllers:ReimbursementController"],
        beego.ControllerComments{
            Method: "GetByOption",
            Router: `/:option`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ReimbursementController"] = append(beego.GlobalControllerRouter["oa/controllers:ReimbursementController"],
        beego.ControllerComments{
            Method: "AddReim",
            Router: `/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ReimbursementController"] = append(beego.GlobalControllerRouter["oa/controllers:ReimbursementController"],
        beego.ControllerComments{
            Method: "GetAllReims",
            Router: `/allReims`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ReimbursementController"] = append(beego.GlobalControllerRouter["oa/controllers:ReimbursementController"],
        beego.ControllerComments{
            Method: "PassReimbursement",
            Router: `/pass/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ReimbursementController"] = append(beego.GlobalControllerRouter["oa/controllers:ReimbursementController"],
        beego.ControllerComments{
            Method: "RejectReimbursement",
            Router: `/reject/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ResignationController"] = append(beego.GlobalControllerRouter["oa/controllers:ResignationController"],
        beego.ControllerComments{
            Method: "AddApplication",
            Router: `/application`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ResignationController"] = append(beego.GlobalControllerRouter["oa/controllers:ResignationController"],
        beego.ControllerComments{
            Method: "PassResignation",
            Router: `/pass/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:ResignationController"] = append(beego.GlobalControllerRouter["oa/controllers:ResignationController"],
        beego.ControllerComments{
            Method: "RejectResignation",
            Router: `/reject/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:TrainingController"] = append(beego.GlobalControllerRouter["oa/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "DeleteTraining",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:TrainingController"] = append(beego.GlobalControllerRouter["oa/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "AddTraining",
            Router: `/addtraining`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:TrainingController"] = append(beego.GlobalControllerRouter["oa/controllers:TrainingController"],
        beego.ControllerComments{
            Method: "GetAllTraining",
            Router: `/getalldata`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:TravelController"] = append(beego.GlobalControllerRouter["oa/controllers:TravelController"],
        beego.ControllerComments{
            Method: "AddTravel",
            Router: `/addtravel`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:VehicleController"] = append(beego.GlobalControllerRouter["oa/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "GetVehicleRecords",
            Router: `/:name`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:VehicleController"] = append(beego.GlobalControllerRouter["oa/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "ApplyVehicle",
            Router: `/application`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:VehicleController"] = append(beego.GlobalControllerRouter["oa/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "PassVehicle",
            Router: `/pass/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["oa/controllers:VehicleController"] = append(beego.GlobalControllerRouter["oa/controllers:VehicleController"],
        beego.ControllerComments{
            Method: "RejectVehicle",
            Router: `/reject/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

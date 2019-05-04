// @APIVersion 1.0.0
// @Title 数据展示 API
// @Description
// @Contact
package routers

import (
	"bishe/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/api/language", &controllers.PositionController{}, "get:GetLanguage")
    beego.Router("/api/city/:language", &controllers.PositionController{}, "get:GetCity")
	beego.Router("/api/salary/:language", &controllers.PositionController{}, "get:GetSalary")
	beego.Router("/api/company/:language", &controllers.PositionController{}, "get:GetCompany")
	beego.Router("/api/workyear/:language", &controllers.PositionController{}, "get:GetWorkYear")
	beego.Router("/api/education/:language", &controllers.PositionController{}, "get:GetEducation")
	// GetCityCompany
	beego.Router("/api/citycompany/:language", &controllers.PositionController{}, "get:GetCityCompany")
	// GetWorkYearSalary
	beego.Router("/api/workyearsalary/:language", &controllers.PositionController{}, "get:GetWorkYearSalary")
	// GetFinanceSalary
	beego.Router("/api/financesalary/:language", &controllers.PositionController{}, "get:GetFinanceSalary")
}

package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitValidators() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 自定义验证方法
		v.RegisterValidation("checkMobile", checkMobile)
	}
}


func checkMobile(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().Int()
	//mobileNum, _ := strconv.Atoi(mobileStr)
	if mobileNum > 10 {
		return true
	} else {
		return false
	}
}

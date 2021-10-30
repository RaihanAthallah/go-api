package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hayvee-website-development/go-api-hayvee/app/controller"
)

func DoctorRouter(r *gin.Engine, c controller.DoctorController) *gin.Engine {
	var gr = r.Group("doctor")
	{
		gr.GET("detail/:id", c.ClinicDetail)
		gr.GET("", c.ClinicList)
		gr.GET("city/:city", c.ClinicCity)
	}
	return r
}

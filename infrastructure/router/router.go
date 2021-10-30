package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hayvee-website-development/go-api-hayvee/app/controller"
)

func DoctorRouter(r *gin.Engine, c controller.DoctorController) *gin.Engine {
	var gr = r.Group("doctor")
	{
		gr.GET("", c.ClinicList)

	}
	return r
}

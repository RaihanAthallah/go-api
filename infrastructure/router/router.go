package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hayvee-website-development/go-api-hayvee/app/controller"
)

func DoctorRouter(r *gin.Engine, c controller.DoctorController) *gin.Engine {
	var gr = r.Group("clinics")
	{
		gr.GET("", c.ClinicList)
		gr.GET("detail/:id", c.ClinicByID)
		gr.GET("city/:city", c.ClinicByCity)
		gr.POST("/register", c.Register)

	}
	return r
}

func LoginRouter(r *gin.Engine, c controller.UserController) *gin.Engine {
	var gr = r.Group("api")
	{
		gr.POST("login", c.LoginEmail)
		gr.POST("register", c.Register)
	}
	return r
}

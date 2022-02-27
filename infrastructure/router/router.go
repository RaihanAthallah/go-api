package router

import (
	"github.com/gin-gonic/gin"
	inventorycontroller "github.com/inventory-management-tokobejo/go-api/app/controller/inventory"
)

func ProductRouter(r *gin.Engine, c inventorycontroller.ProductController) *gin.Engine {
	var gr = r.Group("inventory")
	{
		gr.GET("product/detail/id", c.FindByIDProduct)
		gr.GET("product/detail/sku", c.FindBySKU)
		gr.GET("product/list", c.List)
		gr.POST("product/create", c.Create)
		gr.GET("product/delete/id", c.DeleteByIDProduct)
		gr.GET("product/delete/sku", c.DeleteBySKU)
	}
	return r
}

func TrackingRouter(r *gin.Engine, c inventorycontroller.TrackingController) *gin.Engine {
	var gr = r.Group("inventory")
	{
		gr.GET("tracking/detail/id", c.FindByID)
		gr.GET("tracking/detail/id_product", c.FindByIDProduct)
		gr.GET("tracking/list", c.List)
		gr.POST("tracking/create", c.Create)
		gr.GET("tracking/delete/id", c.DeleteByID)
	}
	return r
}

func VariantRouter(r *gin.Engine, c inventorycontroller.VariantController) *gin.Engine {
	var gr = r.Group("inventory")
	{
		gr.GET("variant/detail/id", c.FindByID)
		gr.GET("variant/detail/id_product", c.FindByIDProduct)
		gr.GET("variant/list", c.List)
		gr.POST("variant/create", c.Create)
		gr.GET("variant/delete/id", c.DeleteByID)
	}
	return r
}

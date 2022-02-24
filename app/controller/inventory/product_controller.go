package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	response "github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	inventoryusecase "github.com/inventory-management-tokobejo/go-api/app/usecase/inventory"
)

type productController struct {
	ProductUsecase inventoryusecase.ProductUsecase
}

type ProductController interface {
	FindByIDProduct(c *gin.Context)
	FindBySKU(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
	DeleteByIDProduct(c *gin.Context)
	DeleteBySKU(c *gin.Context)
}

func NewProductController(
	pu inventoryusecase.ProductUsecase,
) ProductController {
	return &productController{
		pu,
	}
}

func (pc *productController) FindByIDProduct(c *gin.Context) {
	// var input inventoryrequest.RequestCreatedProduct
	paramid, _ := strconv.Atoi(c.Query("id"))
	requestid, _ := c.Get("RequestID")

	result, err := pc.ProductUsecase.FindByIDProduct(paramid)
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelErrUserSave,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
}

func (pc *productController) FindBySKU(c *gin.Context) {
	// var input inventoryrequest.RequestCreatedProduct
	paramsku := c.Query("sku")
	requestid, _ := c.Get("RequestID")

	result, err := pc.ProductUsecase.FindBySKU(paramsku)
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelErrUserSave,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
}

func (pc *productController) List(c *gin.Context) {
	// var input inventoryrequest.RequestCreatedProduct
	requestid, _ := c.Get("RequestID")

	result, err := pc.ProductUsecase.List()
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelErrUserSave,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
}

func (pc *productController) Create(c *gin.Context) {
	var input inventoryrequest.RequestCreatedProduct
	requestid, _ := c.Get("RequestID")

	if err := c.ShouldBind(&input); err != nil {
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, err := pc.ProductUsecase.Create(input)

	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}
	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelErrUserSave,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
}

func (pc *productController) DeleteByIDProduct(c *gin.Context) {
	paramid, _ := strconv.Atoi(c.Query("id"))
	requestid, _ := c.Get("RequestID")

	result, err := pc.ProductUsecase.DeleteByIDProduct(paramid)
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	if result == true {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserSave,
				RequestID: requestid.(string),
			},
			Data: result,
		}
		c.JSON(
			http.StatusOK,
			rsp,
		)
	}

	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelErrUserSave,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)

}

func (pc *productController) DeleteBySKU(c *gin.Context) {
	paramsku := c.Query("sku")
	requestid, _ := c.Get("RequestID")

	result, err := pc.ProductUsecase.DeleteBySKU(paramsku)
	if err != nil {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	if result == true {
		rsp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrUserSave,
				RequestID: requestid.(string),
			},
			Data: result,
		}
		c.JSON(
			http.StatusOK,
			rsp,
		)
	}

	rsp := response.Response{
		Meta: response.Meta{
			Message:   response.RespMeta.TelErrUserSave,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)

}

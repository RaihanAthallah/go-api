package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	response "github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	inventoryusecase "github.com/inventory-management-tokobejo/go-api/app/usecase/inventory"
)

type variantController struct {
	VariantUsecase inventoryusecase.VariantUsecase
}

type VariantController interface {
	FindByID(c *gin.Context)
	FindByIDProduct(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
}

func NewVariantController(
	vu inventoryusecase.VariantUsecase,
) VariantController {
	return &variantController{
		vu,
	}
}

func (vc *variantController) FindByID(c *gin.Context) {
	paramid, _ := strconv.Atoi(c.Query("id"))
	requestid, _ := c.Get("RequestID")

	result, err := vc.VariantUsecase.FindByID(paramid)
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


func (vc *variantController) FindByIDProduct(c *gin.Context) {
	paramidproduct,_ := strconv.Atoi(c.Query("id_product"))
	requestid, _ := c.Get("RequestID")

	result, err := vc.VariantUsecase.FindByIDProduct(paramidproduct)
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

func (vc *variantController) List(c *gin.Context) {
	requestid, _ := c.Get("RequestID")

	result, err := vc.VariantUsecase.List()
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

func (vc *variantController) Create(c *gin.Context) {
	var input inventoryrequest.VariantCreatedRequest
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

	result, err := vc.VariantUsecase.Create(input)

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

func (vc *variantController) DeleteByID(c *gin.Context) {
	paramid, _ := strconv.Atoi(c.Query("id"))
	requestid, _ := c.Get("RequestID")

	result, err := vc.VariantUsecase.DeleteByID(paramid)
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
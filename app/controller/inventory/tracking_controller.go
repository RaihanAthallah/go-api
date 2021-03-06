package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	response "github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	inventoryusecase "github.com/inventory-management-tokobejo/go-api/app/usecase/inventory"
)

type trackingController struct {
	TrackingUsecase inventoryusecase.TrackingUsecase
}

type TrackingController interface {
	FindByID(c *gin.Context)
	FindByIDProduct(c *gin.Context)
	List(c *gin.Context)
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
}

func NewTrackingController(
	tu inventoryusecase.TrackingUsecase,
) TrackingController {
	return &trackingController{
		tu,
	}
}

func (tc *trackingController) FindByID(c *gin.Context) {
	paramid, _ := strconv.Atoi(c.Query("id"))
	requestid, _ := c.Get("RequestID")

	result, err := tc.TrackingUsecase.FindByID(paramid)
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


func (tc *trackingController) FindByIDProduct(c *gin.Context) {
	paramidproduct,_ := strconv.Atoi(c.Query("id_product"))
	requestid, _ := c.Get("RequestID")

	result, err := tc.TrackingUsecase.FindByIDProduct(paramidproduct)
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

func (tc *trackingController) List(c *gin.Context) {
	requestid, _ := c.Get("RequestID")

	result, err := tc.TrackingUsecase.List()
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

func (tc *trackingController) Create(c *gin.Context) {
	var input inventoryrequest.RequestCreatedTracking
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

	result, err := tc.TrackingUsecase.Create(input)

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

func (tc *trackingController) DeleteByID(c *gin.Context) {
	paramid, _ := strconv.Atoi(c.Query("id_product"))
	requestid, _ := c.Get("RequestID")

	result, err := tc.TrackingUsecase.DeleteByID(paramid)
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
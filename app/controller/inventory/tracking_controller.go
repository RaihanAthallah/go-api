package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	response "github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	// inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	inventoryusecase "github.com/inventory-management-tokobejo/go-api/app/usecase/inventory"
)

type trackingController struct {
	TrackingUsecase inventoryusecase.TrackingUsecase
}

type TrackingController interface {
	FindByID(c *gin.Context)
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

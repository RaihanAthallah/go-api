package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hayvee-website-development/go-api-hayvee/app/usecase"
	response "github.com/hayvee-website-development/go-entites-hayvee/entities/response"
)

type doctorController struct {
	DoctorUsecase usecase.DoctorUsecase
}

type DoctorController interface {
	ClinicDetail(c *gin.Context)
	ClinicCity(c *gin.Context)
	ClinicList(c *gin.Context)
}

func NewDoctorController(d usecase.DoctorUsecase) DoctorController {
	return &doctorController{d}
}

func (d doctorController) ClinicDetail(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("id"))
	requestid, _ := c.Get("RequestID")
	result, err := d.DoctorUsecase.DetailClinic(param)
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
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}

func (d doctorController) ClinicCity(c *gin.Context) {
	param := c.Param("city")
	requestid, _ := c.Get("RequestID")
	result, err := d.DoctorUsecase.FindByCity(param)
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
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}

func (d doctorController) ClinicList(c *gin.Context) {
	requestid, _ := c.Get("RequestID")
	result, err := d.DoctorUsecase.List(c)
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
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(
		http.StatusOK,
		rsp,
	)
	return
}

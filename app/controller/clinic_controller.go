package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	req "github.com/hayvee-website-development/go-api-hayvee/app/model/request/user"
	"github.com/hayvee-website-development/go-api-hayvee/app/usecase"
	logger "github.com/hayvee-website-development/go-api-hayvee/infrastructure/io"
	response "github.com/hayvee-website-development/go-entites-hayvee/entities/response"
)

type doctorController struct {
	DoctorUsecase usecase.DoctorUsecase
}

type DoctorController interface {
	ClinicList(c *gin.Context)
	ClinicByID(c *gin.Context)
	ClinicByCity(c *gin.Context)
	Register(c *gin.Context)
}

func NewDoctorController(d usecase.DoctorUsecase) DoctorController {
	return &doctorController{d}
}

func (d doctorController) ClinicList(c *gin.Context) {
	requestid, _ := c.Get("RequestID")
	result, err := d.DoctorUsecase.List()
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

}

func (d doctorController) ClinicByID(c *gin.Context) {
	params, _ := strconv.Atoi(c.Param("id"))
	requestid, _ := c.Get("RequestID")
	result, err := d.DoctorUsecase.FindID(params)
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

}

func (d doctorController) ClinicByCity(c *gin.Context) {
	params := c.Param("city")
	requestid, _ := c.Get("RequestID")
	result, err := d.DoctorUsecase.FindCity(params)
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

}

func (uc doctorController) Register(c *gin.Context) {
	var input req.RegRegister
	var result interface{}
	var errRes error
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "Register"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)

	if err := c.ShouldBind(&input); err != nil {
		cl.Errorf("[ERROR] %v", err.Error())
		resp := response.Response{
			Meta: response.Meta{
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if errVal := input.Validate(); errVal != nil {
		cl.Errorf("[ERROR] %v", errVal.Error())
		resp := response.Response{
			Meta: response.Meta{
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	result, errRes = uc.DoctorUsecase.Create(input)
	if errRes != nil {
		cl.Errorf("[ERROR] %v", response.RespMeta.TelErrUserSave)
		resp := response.Response{
			Meta: response.Meta{
				Message:   response.RespMeta.TelErrEmailAlreadyUsed,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := response.Response{
		Meta: response.Meta{
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(http.StatusOK, resp)
}

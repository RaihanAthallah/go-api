package controller

import (
	"net/http"

	request "github.com/hayvee-website-development/go-api-hayvee/app/model/request/user"
	"github.com/hayvee-website-development/go-api-hayvee/app/usecase"
	logger "github.com/hayvee-website-development/go-api-hayvee/infrastructure/io"

	responsemsg "github.com/hayvee-website-development/go-entites-hayvee/entities/response"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/copier"
)

type screeningController struct {
	ScreeningUsecase usecase.ScreeningUsecase
}

type ScreeningController interface {
	ScreeningAnswer(c *gin.Context)
}

func NewScreeningController(
	du usecase.ScreeningUsecase,
) ScreeningController {
	return &screeningController{
		du,
	}
}

func (uc *screeningController) ScreeningAnswer(c *gin.Context) {
	var input request.RegScreening
	paramtoken := c.Query("token")
	var result interface{}
	var errRes error
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "Register"})
	cl.Infof("[INFO] Header values: %v", c.Request.Header)

	if err := c.ShouldBind(&input); err != nil {
		cl.Errorf("[ERROR] %v", err.Error())
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if errVal := input.Validate(); errVal != nil {
		cl.Errorf("[ERROR] %v", errVal.Error())
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result = uc.ScreeningUsecase.FindByID(paramtoken)
	if result == false {
		result, errRes = uc.ScreeningUsecase.Create(input, paramtoken)
		if errRes != nil {
			cl.Errorf("[ERROR] %v", responsemsg.RespMeta.TelErrUserSave)
			resp := responsemsg.Response{
				Meta: responsemsg.Meta{
					Message:   responsemsg.RespMeta.TelErrUserSave,
					RequestID: requestid.(string),
				},
			}
			c.JSON(http.StatusBadRequest, resp)
			return
		}
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				RequestID: requestid.(string),
			},
			Data: result,
		}
		c.JSON(http.StatusOK, resp)

	} else {
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				RequestID: requestid.(string),
			},
			Data: result,
		}
		c.JSON(http.StatusOK, resp)
	}
}

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

type userController struct {
	UserUsecase usecase.UserUsercase
}

type UserController interface {
	LoginEmail(c *gin.Context)
	Register(c *gin.Context)
	// LoginEmailHayvee(c *gin.Context)
	// RegisterHayvee(c *gin.Context)
	// GetAllUserBySuperAdmin(c *gin.Context)
}

func NewUserController(
	du usecase.UserUsercase,
) UserController {
	return &userController{
		du,
	}
}

// func (uc *userController) GetAllUserBySuperAdmin(c *gin.Context) {
// 	paramtoken := c.Query("token")
// 	paramname := c.Query("name")
// 	paramemail := c.Query("email")
// 	paramphone := c.Query("phone")
// 	paramidentifier, _ := strconv.Atoi(c.Query("id_identifier"))
// 	requestid, _ := c.Get("RequestID")
// 	result := respuser.AllUsers{}
// 	resultdoctor, err := uc.DoctorUsecase.FindAll(
// 		paramtoken,
// 		paramname,
// 		paramemail,
// 		paramphone,
// 		paramidentifier,
// 	)
// 	result.KDDokter = resultdoctor
// 	resultpatient, _ := uc.PatientUsecase.FindAll(
// 		paramtoken,
// 		paramname,
// 		paramemail,
// 		paramphone,
// 		paramidentifier,
// 	)
// 	result.KDPasien = resultpatient
// 	resultsuperadmin, _ := uc.SuperAdminUsecase.FindAll(
// 		paramtoken,
// 		paramname,
// 		paramemail,
// 		paramphone,
// 		paramidentifier)
// 	result.KDSuperAdmin = resultsuperadmin

// 	if err != nil {
// 		rsp := response.Response{
// 			Meta: response.Meta{
// 				Message:   response.RespMeta.TelErrUserNotFound,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, rsp)
// 		return
// 	}
// 	rsp := response.Response{
// 		Meta: response.Meta{
// 			RequestID: requestid.(string),
// 		},
// 		Data: result,
// 	}
// 	c.JSON(
// 		http.StatusOK,
// 		rsp,
// 	)
// }

func (uc *userController) LoginEmail(c *gin.Context) {
	var input request.RegLoginEmail
	requestid, _ := c.Get("RequestID")
	cl := logger.WithFields(logger.Fields{"UserController": "LoginEmail"})
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
	if err := input.Validate(); err != nil {
		cl.Errorf("[ERROR] Error on field: %v", err.Error())
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if uc.UserUsecase.IsDuplicateEmail(input.Email) {
		cl.Errorf("[ERROR] Error on field: %v", responsemsg.RespMeta.TelErrEmailAlreadyUsed)
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				Message:   responsemsg.RespMeta.TelErrEmailAlreadyUsed,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	result := uc.UserUsecase.VerifyEmail(input.Email, input.Password)
	if result == false {
		cl.Errorf("[ERROR] Error on field: %v", responsemsg.RespMeta.TelErrEmailAlreadyUsed)
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				Message:   responsemsg.RespMeta.TelErrUserNotFound,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := responsemsg.Response{
		Meta: responsemsg.Meta{
			Message:   responsemsg.RespMeta.TelErrUserIsActive,
			RequestID: requestid.(string),
		},
		Data: result,
	}
	c.JSON(http.StatusOK, resp)
}

func (uc *userController) Register(c *gin.Context) {
	var input request.RegRegisterUser
	var result interface{}
	//	var errTripleDes error
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

	if !uc.UserUsecase.IsDuplicateEmail(input.Email) {
		cl.Errorf("[ERROR] %v", responsemsg.RespMeta.TelErrEmailAlreadyUsed)
		resp := responsemsg.Response{
			Meta: responsemsg.Meta{
				Message:   responsemsg.RespMeta.TelErrEmailAlreadyUsed,
				RequestID: requestid.(string),
			},
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	result, errRes = uc.UserUsecase.Create(input)

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
}

// func (uc *userController) LoginEmailHayvee(c *gin.Context) {
// 	var input request.RegLoginEmail
// 	requestid, _ := c.Get("RequestID")
// 	cl := logger.WithFields(logger.Fields{"UserController": "LoginEmail"})
// 	cl.Infof("[INFO] Header values: %v", c.Request.Header)

// 	if err := c.ShouldBind(&input); err != nil {
// 		cl.Errorf("[ERROR] %v", err.Error())
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}
// 	if err := input.Validate(); err != nil {
// 		cl.Errorf("[ERROR] Error on field: %v", err.Error())
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}

// 	if uc.UserUsecase.IsDuplicateEmail(input.Email) {
// 		cl.Errorf("[ERROR] Error on field: %v", responsemsg.RespMeta.TelErrEmailAlreadyUsed)
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrEmailAlreadyUsed,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}
// 	result := uc.UserUsecase.VerifyEmail(input.Email, input.Password)
// 	if result == false {
// 		cl.Errorf("[ERROR] Error on field: %v", responsemsg.RespMeta.TelErrEmailAlreadyUsed)
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrUserNotFound,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}
// 	resp := responsemsg.Response{
// 		Meta: responsemsg.Meta{
// 			Message:   responsemsg.RespMeta.TelErrUserIsActive,
// 			RequestID: requestid.(string),
// 		},
// 		Data: result,
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

// func (uc *userController) RegisterHayvee(c *gin.Context) {
// 	var input request.RegRegisterUser
// 	var result interface{}
// 	//	var errTripleDes error
// 	var errRes error
// 	requestid, _ := c.Get("RequestID")
// 	cl := logger.WithFields(logger.Fields{"UserController": "Register"})
// 	cl.Infof("[INFO] Header values: %v", c.Request.Header)

// 	if err := c.ShouldBind(&input); err != nil {
// 		cl.Errorf("[ERROR] %v", err.Error())
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}
// 	if errVal := input.Validate(); errVal != nil {
// 		cl.Errorf("[ERROR] %v", errVal.Error())
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrCodeNotValid,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}

// 	if !uc.UserUsecase.IsDuplicateEmail(input.Email) {
// 		cl.Errorf("[ERROR] %v", responsemsg.RespMeta.TelErrEmailAlreadyUsed)
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrEmailAlreadyUsed,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}
// 	result, errRes = uc.UserUsecase.Create(input)

// 	if errRes != nil {
// 		cl.Errorf("[ERROR] %v", responsemsg.RespMeta.TelErrUserSave)
// 		resp := responsemsg.Response{
// 			Meta: responsemsg.Meta{
// 				Message:   responsemsg.RespMeta.TelErrEmailAlreadyUsed,
// 				RequestID: requestid.(string),
// 			},
// 		}
// 		c.JSON(http.StatusBadRequest, resp)
// 		return
// 	}
// 	resp := responsemsg.Response{
// 		Meta: responsemsg.Meta{
// 			RequestID: requestid.(string),
// 		},
// 		Data: result,
// 	}
// 	c.JSON(http.StatusOK, resp)
// }

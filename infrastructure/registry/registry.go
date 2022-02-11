package registry

import (
	"github.com/hayvee-website-development/go-api-hayvee/app/controller"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewDoctorController() controller.DoctorController
	NewUserController() controller.UserController
	NewScreeningController() controller.ScreeningController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

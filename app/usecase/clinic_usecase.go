package usecase

import (
	"github.com/gin-gonic/gin"
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/listclinics"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
)

type doctorUsecase struct {
	BaseRepository   repository.BaseRepository
	ClinicRepository repository.ClinicRepository
}

type DoctorUsecase interface {
	List(c *gin.Context) (list []list.HvClinic, err error)
}

func NewDoctorUsecase(
	br repository.BaseRepository,
	cr repository.ClinicRepository,
) DoctorUsecase {
	return &doctorUsecase{br, cr}
}

func (d *doctorUsecase) List(c *gin.Context) (list []list.HvClinic, err error) {
	return d.ClinicRepository.List()
}

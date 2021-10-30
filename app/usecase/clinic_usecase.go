package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/listclinics"
	rsp "github.com/hayvee-website-development/go-api-hayvee/app/model/response/doctor"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
	"github.com/hayvee-website-development/go-api-hayvee/infrastructure/io"
	"github.com/jinzhu/copier"
)

type doctorUsecase struct {
	BaseRepository   repository.BaseRepository
	ClinicRepository repository.ClinicRepository
}

type DoctorUsecase interface {
	DetailClinic(id int) (list *list.HvClinic, err error)
	FindByCity(city string) (list []list.HvClinic, err error)
	List(c *gin.Context) (list []list.HvClinic, err error)
}

func NewDoctorUsecase(
	br repository.BaseRepository,
	cr repository.ClinicRepository,
) DoctorUsecase {
	return &doctorUsecase{br, cr}
}

func (d *doctorUsecase) DetailClinic(id int) (list *list.HvClinic, err error) {
	return d.ClinicRepository.Find(id int)
}

func (d *doctorUsecase) FindByCity(city string) (list []list.HvClinic, err error) {
	return d.ClinicRepository.FindByCity(city string)
}

func (d *doctorUsecase) List(c *gin.Context) (list []list.HvClinic, err error) {
	requestid, _ := c.Get("RequestID")
	cl := io.WithFields(io.Fields{"RequestID": requestid})
	return d.ClinicRepository.List()
}

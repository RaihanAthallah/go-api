package usecase

import (
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/listclinics"
	req "github.com/hayvee-website-development/go-api-hayvee/app/model/request/user"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
)

type doctorUsecase struct {
	BaseRepository   repository.BaseRepository
	ClinicRepository repository.ClinicRepository
}

type DoctorUsecase interface {
	List() (list []list.HvClinic, err error)
	FindCity(city string) (list []list.HvClinic, err error)
	FindID(id int) (list *list.HvClinic, err error)
	Create(user req.RegRegister) (interface{}, error)
}

func NewDoctorUsecase(
	br repository.BaseRepository,
	cr repository.ClinicRepository,
) DoctorUsecase {
	return &doctorUsecase{br, cr}
}

func (d *doctorUsecase) List() (list []list.HvClinic, err error) {
	return d.ClinicRepository.List()
}

func (d *doctorUsecase) FindCity(city string) (list []list.HvClinic, err error) {
	return d.ClinicRepository.FindByCity(city)
}

func (d *doctorUsecase) FindID(id int) (list *list.HvClinic, err error) {
	return d.ClinicRepository.Find(id)
}

func (uu *doctorUsecase) Create(input req.RegRegister) (interface{}, error) {
	uu.BaseRepository.BeginTx()
	user := list.HvClinic{
		Name:        input.Name,
		Address:     input.Address,
		ServiceTime: input.ServiceTime,
		City:        input.City,
		Contact:     input.Contact,
		Province:    input.Province,
		PostalCode:  input.PostalCode,
	}
	rUser, err := uu.ClinicRepository.Create(user)
	if err != nil {
		uu.BaseRepository.RollbackTx()
		return false, err
	}
	return rUser, err
}

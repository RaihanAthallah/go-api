package usecase

import (
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/listclinics"
	req "github.com/hayvee-website-development/go-api-hayvee/app/model/request/user"
	rsp "github.com/hayvee-website-development/go-api-hayvee/app/model/response/doctor"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/copier"
)

type doctorUsecase struct {
	BaseRepository   repository.BaseRepository
	ClinicRepository repository.ClinicRepository
}

type DoctorUsecase interface {
	List() (result []interface{}, err error)
	FindCity(city string) (result []interface{}, err error)
	FindID(id int) (result interface{}, err error)
	Create(user req.RegRegister) (interface{}, error)
}

func NewDoctorUsecase(
	br repository.BaseRepository,
	cr repository.ClinicRepository,
) DoctorUsecase {
	return &doctorUsecase{br, cr}
}

func (d *doctorUsecase) List() (result []interface{}, err error) {
	faker := faker.New()
	p := faker.Person()
	image := p.Image()

	listalldataclinics, err := d.ClinicRepository.List()
	for _, dataclinics := range listalldataclinics {
		responselistclinics := rsp.ListDoctor{}
		copier.Copy(&responselistclinics, &dataclinics)
		responselistclinics.Avatar = image.Name()

		result = append(result, responselistclinics)
	}
	return result, err
}

func (d *doctorUsecase) FindCity(city string) (result []interface{}, err error) {
	faker := faker.New()
	p := faker.Person()
	image := p.Image()

	listalldataclinics, err := d.ClinicRepository.FindByCity(city)
	for _, dataclinics := range listalldataclinics {
		responselistclinics := rsp.ListDoctor{}
		copier.Copy(&responselistclinics, &dataclinics)
		responselistclinics.Avatar = image.Name()

		result = append(result, responselistclinics)
	}
	return result, err
}

func (d *doctorUsecase) FindID(id int) (result interface{}, err error) {
	faker := faker.New()
	p := faker.Person()
	image := p.Image()

	listalldataclinics, err := d.ClinicRepository.Find(id)
	responselistclinics := rsp.ListDoctor{}
	copier.Copy(&responselistclinics, &listalldataclinics)
	responselistclinics.Avatar = image.Name()

	return responselistclinics, err
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

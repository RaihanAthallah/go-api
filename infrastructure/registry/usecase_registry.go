package registry

import (
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
	"github.com/hayvee-website-development/go-api-hayvee/app/usecase"
)

func (r *registry) NewDoctorUsecase() usecase.DoctorUsecase {
	return usecase.NewDoctorUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewClinicRepository(repository.NewBaseRepository(r.db)),
	)
}

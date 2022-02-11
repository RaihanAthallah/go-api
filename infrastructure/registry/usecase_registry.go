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

func (r *registry) NewUserUsercase() usecase.UserUsercase {
	return usecase.NewUserUsercase(
		repository.NewBaseRepository(r.db),
		repository.NewUserRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserDataRepository(repository.NewBaseRepository(r.db)),
		usecase.NewUserAccessUsecase(
			repository.NewBaseRepository(r.db),
			repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
			repository.NewUserRepository(repository.NewBaseRepository(r.db)),
		),
	)
}

func (r *registry) NewUserAccessUsecase() usecase.UserAccessUsecase {
	return usecase.NewUserAccessUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserRepository(repository.NewBaseRepository(r.db)),
	)
}

func (r *registry) NewScreeningUsecase() usecase.ScreeningUsecase {
	return usecase.NewScreeningUsecase(
		repository.NewBaseRepository(r.db),
		repository.NewScreeningRepository(repository.NewBaseRepository(r.db)),
		repository.NewUserAccessRepository(repository.NewBaseRepository(r.db)),
	)
}

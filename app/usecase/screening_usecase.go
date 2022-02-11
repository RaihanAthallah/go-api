package usecase

import (
	entity "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
	req "github.com/hayvee-website-development/go-api-hayvee/app/model/request/user"
	"github.com/hayvee-website-development/go-api-hayvee/app/repository"
)

type screeningUsecase struct {
	BaseRepository       repository.BaseRepository
	ScreeningRepository  repository.ScreeningRepository
	UserAccessRepository repository.UserAccessRepository
}

type ScreeningUsecase interface {
	FindByID(token string) bool
	Create(input req.RegScreening, token string) (interface{}, error)
}

func NewScreeningUsecase(
	br repository.BaseRepository,
	cr repository.ScreeningRepository,
	uar repository.UserAccessRepository,
) ScreeningUsecase {
	return &screeningUsecase{br, cr, uar}
}

func (d *screeningUsecase) FindByID(token string) bool {
	findiduser, _ := d.UserAccessRepository.FindIDByToken(token)

	verifyiduser, _ := d.ScreeningRepository.FindByID(findiduser.IDUser)

	return verifyiduser != nil
	// if verifyiduser != nil {
	// 	return true
	// }
	// return false
}

func (d *screeningUsecase) Create(input req.RegScreening, token string) (interface{}, error) {
	findiduser, _ := d.UserAccessRepository.FindIDByToken(token)
	d.BaseRepository.BeginTx()
	screening := entity.HvScreening{
		IDUser:  findiduser.IDUser,
		Number1: input.Number1,
		Number2: input.Number2,
		Number3: input.Number3,
		Number4: input.Number4,
		Number5: input.Number5,
	}
	rUser, err := d.ScreeningRepository.CreateAnswer(screening)
	if err != nil {
		d.BaseRepository.RollbackTx()
		return false, err
	}
	return rUser, err
}

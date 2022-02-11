//Package repository provide all query that need for usecase
package repository

import (
	"github.com/golang-module/carbon/v2"
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
)

type screeningRepository struct {
	base BaseRepository
}

type ScreeningRepository interface {
	CreateAnswer(screening list.HvScreening) (list.HvScreening, error)
	FindByID(iduser int) (*list.HvScreening, error)
	VerifyIDUser(iduser int) bool
}

func NewScreeningRepository(ar BaseRepository) ScreeningRepository {
	return &screeningRepository{ar}
}

func (r *screeningRepository) CreateAnswer(screening list.HvScreening) (list.HvScreening, error) {
	var baseDb = r.base

	screening.LastUpdate = carbon.Now().ToDateTimeString()
	screening.IsDone = true
	query := baseDb.GetDB().
		Create(&screening)
	err := query.Error
	return screening, err
}

func (r *screeningRepository) FindByID(iduser int) (*list.HvScreening, error) {
	var user list.HvScreening
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&list.HvScreening{IDUser: iduser})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *screeningRepository) VerifyIDUser(iduser int) bool {
	var user list.HvScreening
	sql := r.base.GetDB().Where("iduser = ?", iduser).
		First(&user)
	if sql.Error != nil {
		return false
	}
	if sql.RowsAffected == 0 {
		return false
	}
	return true
}

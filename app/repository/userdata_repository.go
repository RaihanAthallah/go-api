package repository

import (
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
)

type userDataRepository struct {
	base BaseRepository
}

type UserDataRepository interface {
	FindByNIM(iduser int) (*list.HvUserData, error)
}

func NewUserDataRepository(ar BaseRepository) UserDataRepository {
	return &userDataRepository{ar}
}

func (r *userDataRepository) FindByNIM(iduser int) (*list.HvUserData, error) {
	var user list.HvUserData
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&list.HvUserData{IDUser: iduser})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

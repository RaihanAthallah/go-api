package repository

import (
	"github.com/golang-module/carbon/v2"
	entity "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
)

type userAccessRepository struct {
	base BaseRepository
}

type UserAccessRepository interface {
	Find(id int) (entity.HvUserAccess, error)
	Create(ud entity.HvUserAccess) error
	Update(id int, token string) error
	ValidToken(token string) bool
	ValidTokenWithID(iduser int, token string) bool
}

func NewUserAccessRepository(ar BaseRepository) UserAccessRepository {
	return &userAccessRepository{ar}
}

func (r *userAccessRepository) Find(iduser int) (entity.HvUserAccess, error) {
	var ua entity.HvUserAccess
	err := r.base.GetDB().Where("iduser = ?", iduser).
		First(&ua).Error
	if err != nil {
		return ua, err
	}
	return ua, nil
}
func (r *userAccessRepository) ValidToken(token string) bool {
	var ua entity.HvUserAccess
	sql := r.base.GetDB().Where("token = ?", token).
		First(&ua)
	if sql.Error != nil {
		return false
	}
	if sql.RowsAffected == 0 {
		return false
	}
	return true
}
func (r *userAccessRepository) ValidTokenWithID(iduser int, token string) bool {
	var ua entity.HvUserAccess
	sql := r.base.GetDB().Where("iduser = ? and token = ?", iduser, token).
		First(&ua)
	if sql.Error != nil {
		return false
	}
	if sql.RowsAffected == 0 {
		return false
	}
	return true
}
func (r *userAccessRepository) Create(ua entity.HvUserAccess) error {
	ua.LastUpdate = carbon.Now().ToDateTimeString()
	err := r.base.GetDB().Create(&ua).Error
	return err
}
func (r *userAccessRepository) Update(iduser int, token string) error {
	err := r.base.GetDB().Model(entity.HvUserAccess{}).
		Where("iduser = ?", iduser).
		Updates(entity.HvUserAccess{Token: token}).Error
	if err != nil {
		return err
	}
	return nil
}

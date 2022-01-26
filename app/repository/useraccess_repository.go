package repository

import entity "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"

type userAccessRepository struct {
	base BaseRepository
}

type UserAccessRepository interface {
	Find(id int) (entity.HvUserAccess, error)
	Create(ud entity.HvUserAccess) error
	Update(id int, token string) error
	ValidTokenWithID(id int, token string) bool
	ValidToken(token string) (*entity.HvUserAccess, error)
	DeleteToken(iduser int, token string) (*entity.HvUserAccess, error)
}

func NewUserAccessRepository(ar BaseRepository) UserAccessRepository {
	return &userAccessRepository{ar}
}

func (r *userAccessRepository) Find(iduser int) (entity.HvUserAccess, error) {
	var ua entity.HvUserAccess
	err := r.base.GetDB().Where(&entity.HvUserAccess{IDUser: iduser}).
		First(&ua).Error
	if err != nil {
		return ua, err
	}
	return ua, nil
}
func (r *userAccessRepository) ValidTokenWithID(iduser int, token string) bool {
	var ua entity.HvUserAccess
	sql := r.base.GetDB().Where("userid = ? and tokenlogin = ?", iduser, token).
		First(&ua)
	if sql.Error != nil {
		return false
	}
	if sql.RowsAffected == 0 {
		return false
	}
	return true
}
func (r *userAccessRepository) ValidToken(token string) (*entity.HvUserAccess, error) {
	var user entity.HvUserAccess
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.HvUserAccess{Token: token})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userAccessRepository) Create(ua entity.HvUserAccess) error {
	err := r.base.GetDB().Create(&ua).Error
	return err
}
func (r *userAccessRepository) Update(iduser int, token string) error {
	err := r.base.GetDB().Model(entity.HvUserAccess{}).
		Where("userid = ?", iduser).
		Updates(entity.HvUserAccess{Token: token}).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *userAccessRepository) DeleteToken(iduser int, token string) (*entity.HvUserAccess, error) {
	var user entity.HvUserAccess
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&entity.HvUserAccess{IDUser: iduser}).
		Model(&user).
		Updates(map[string]interface{}{"tokenlogin": nil})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

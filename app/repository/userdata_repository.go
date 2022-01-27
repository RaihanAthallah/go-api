package repository

import (
	entity "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
)

type userDataRepository struct {
	base BaseRepository
}

type UserDataRepository interface {
	FindByID(id int) (*entity.HvUserData, error)
	FindAll() ([]entity.HvUserData, error)
	Create(ua entity.HvUserData) (entity.HvUserData, error)
	List() ([]entity.HvUserData, error)
	Update(id int, input map[string]interface{}) error
	FindByParam(filter map[string]interface{}) ([]entity.HvUserData, error)
}

func NewUserDataRepository(ar BaseRepository) UserDataRepository {
	return &userDataRepository{ar}
}

func (r *userDataRepository) FindByID(id int) (*entity.HvUserData, error) {
	var tr entity.HvUserData
	err := r.base.GetDB().
		Where(entity.HvUserData{IDUser: id}).
		First(&tr).Error
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

func (r *userDataRepository) FindAll() ([]entity.HvUserData, error) {
	var tr []entity.HvUserData
	err := r.base.GetDB().
		Find(&tr).Error
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (r *userDataRepository) Create(ua entity.HvUserData) (entity.HvUserData, error) {
	err := r.base.GetDB().Create(&ua).Error
	return ua, err
}

func (r *userDataRepository) List() ([]entity.HvUserData, error) {
	var user []entity.HvUserData
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userDataRepository) Update(id int, input map[string]interface{}) error {
	var user entity.HvUserData
	err := r.base.GetDB().Model(&user).
		Where("id_superadmin = ?", id).
		Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userDataRepository) FindByParam(filter map[string]interface{}) ([]entity.HvUserData, error) {
	var superadmin []entity.HvUserData
	var baseDb = r.base

	query := baseDb.GetDB()

	if filter["name"] != nil {
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+filter["name"].(string)+"%", "%"+filter["name"].(string)+"%")
	}

	if filter["email"] != nil {
		query = query.Where("email ILIKE ?", "%"+filter["email"].(string)+"%")
	}

	if filter["phone"] != nil {
		query = query.Where("phone ILIKE ?", "%"+filter["phone"].(string)+"%")
	}

	if filter["id_identifier"] != 0 {
		query = query.Where("id_identifier = ?", filter["id_identifier"].(int))
	}

	err := query.Find(&superadmin).Error
	if err != nil {
		return nil, err
	}
	return superadmin, nil
}

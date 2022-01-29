//Package repository provide all query that need for usecase
package repository

import (
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/consultation"
)

type userRepository struct {
	base BaseRepository
}

type UserRepository interface {
	Find(id int) (*list.HvUser, error)
	FindByParam(filter map[string]interface{}) (*list.HvUser, error)
	Create(user list.HvUser) (list.HvUser, error)
	Update(id int, input map[string]interface{}) error
	List() ([]list.HvUser, error)
}

func NewUserRepository(ar BaseRepository) UserRepository {
	return &userRepository{ar}
}

func (r *userRepository) Find(id int) (*list.HvUser, error) {
	var user list.HvUser
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&list.HvUser{IDUser: id})

	err := query.First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByParam(filter map[string]interface{}) (*list.HvUser, error) {
	var user list.HvUser
	query := r.base.GetDB()

	if filter["email"] != nil {
		query = query.Where(&list.HvUser{Email: filter["email"].(string)})
	}
	err := query.First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}

func (r *userRepository) Create(user list.HvUser) (list.HvUser, error) {
	err := r.base.GetDB().Create(&user).Error
	return user, err
}

func (r *userRepository) Update(id int, input map[string]interface{}) error {
	var user list.HvUser
	err := r.base.GetDB().Model(&user).
		Where("iduser = ?", id).
		Updates(&input).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) List() ([]list.HvUser, error) {
	var user []list.HvUser
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

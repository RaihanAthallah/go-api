package repository

import (
	list "github.com/hayvee-website-development/go-api-hayvee/app/model/entity/listclinics"
)

type clinicRepository struct {
	base BaseRepository
}

type ClinicRepository interface {
	Find(id int) (*list.HvClinic, error)
	FindByCity(city string) ([]list.HvClinic, error)
	List() ([]list.HvClinic, error)
	Create(user list.HvClinic) (list.HvClinic, error)
}

func NewClinicRepository(ar BaseRepository) ClinicRepository {
	return &clinicRepository{ar}
}

func (r *clinicRepository) Find(id int) (*list.HvClinic, error) {
	var clinic *list.HvClinic
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&list.HvClinic{ID: id})

	err := query.First(&clinic).Error
	if err != nil {
		return nil, err
	}
	return clinic, nil
}

func (r *clinicRepository) FindByCity(city string) ([]list.HvClinic, error) {
	var clinic []list.HvClinic
	var baseDb = r.base

	query := baseDb.GetDB().
		Where(&list.HvClinic{City: city})

	err := query.Find(&clinic).Error
	if err != nil {
		return nil, err
	}
	return clinic, nil
}

func (r *clinicRepository) List() ([]list.HvClinic, error) {
	var clinic []list.HvClinic
	var baseDb = r.base

	query := baseDb.GetDB()

	err := query.Find(&clinic).Error
	if err != nil {
		return nil, err
	}
	return clinic, nil
}

func (r *clinicRepository) Create(user list.HvClinic) (list.HvClinic, error) {
	err := r.base.GetDB().Create(&user).Error
	return user, err
}

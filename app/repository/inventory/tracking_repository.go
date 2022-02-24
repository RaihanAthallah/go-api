package repository

import (
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
)

type trackingRepository struct {
	base baserepository.BaseRepository
}

type TrackingRepository interface {
	FindByID(id int) (*inventoryentity.Tracking, error)
	FindByIDProduct(id_product int) (*inventoryentity.Tracking, error)
	List() ([]inventoryentity.Tracking, error)
	Create(tracking inventoryentity.Tracking) (inventoryentity.Tracking, error)
	DeleteByID(id int) (bool, error)
}

func NewTrackingRepository(ar baserepository.BaseRepository) TrackingRepository {
	return &trackingRepository{ar}
}

func (tr *trackingRepository) FindByID(id int) (*inventoryentity.Tracking, error) {
	var trackingentity *inventoryentity.Tracking
	var baseDb = tr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Tracking{Id: id})

	err := query.First(&trackingentity).Error
	if err != nil {
		return nil, err
	}
	return trackingentity, nil
}

func (tr *trackingRepository) FindByIDProduct(id_product int) (*inventoryentity.Tracking, error) {
	var trackingentity *inventoryentity.Tracking
	var baseDb = tr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Tracking{Id_Product: id_product})

	err := query.First(&trackingentity).Error
	if err != nil {
		return nil, err
	}
	return trackingentity, nil
}

func (tr *trackingRepository) List() ([]inventoryentity.Tracking, error) {
	var trackingentity []inventoryentity.Tracking
	var baseDb = tr.base

	query := baseDb.GetDB()

	err := query.Find(&trackingentity).Error
	if err != nil {
		return nil, err
	}
	return trackingentity, nil
}

func (tr *trackingRepository) Create(tracking inventoryentity.Tracking) (inventoryentity.Tracking, error) {
	err := tr.base.GetDB().Create(&tracking).Error
	return tracking, err
}

func (tr *trackingRepository) DeleteByID(id int) (bool, error) {
	var trackingentity inventoryentity.Tracking
	var baseDb = tr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Tracking{Id: id}).
		Delete(&trackingentity)

	err := query.First(&trackingentity).Error

	if err != nil {
		return false, err
	}

	return true, err
}

package repository

import (
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
)

type variantRepository struct {
	base baserepository.BaseRepository
}

type VariantRepository interface {
	FindByID(id int) (*inventoryentity.Variant, error)
	FindByIDProduct(id_product int) (*inventoryentity.Variant, error)
	List() ([]inventoryentity.Variant, error)
	Create(variant inventoryentity.Variant) (inventoryentity.Variant, error)
	DeleteByID(id int) (bool, error)
}

func NewVariantRepository(vr baserepository.BaseRepository) VariantRepository {
	return &variantRepository{vr}
}

func (vr *variantRepository) FindByID(id int) (*inventoryentity.Variant, error) {
	var variantentity *inventoryentity.Variant
	var baseDb = vr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Variant{Id: id})

	err := query.First(&variantentity).Error
	if err != nil {
		return nil, err
	}
	return variantentity, nil
}

func (vr *variantRepository) FindByIDProduct(id_product int) (*inventoryentity.Variant, error) {
	var variantentity *inventoryentity.Variant
	var baseDb = vr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Variant{Id_Product: id_product})

	err := query.First(&variantentity).Error
	if err != nil {
		return nil, err
	}
	return variantentity, nil
}

func (vr *variantRepository) List() ([]inventoryentity.Variant, error) {
	var variantentity []inventoryentity.Variant
	var baseDb = vr.base

	query := baseDb.GetDB()

	err := query.Find(&variantentity).Error
	if err != nil {
		return nil, err
	}
	return variantentity, nil
}

func (vr *variantRepository) Create(variant inventoryentity.Variant) (inventoryentity.Variant, error) {
	err := vr.base.GetDB().Create(&variant).Error
	return variant, err
}

func (vr *variantRepository) DeleteByID(id int) (bool, error) {
	var variantentity inventoryentity.Variant
	var baseDb = vr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Variant{Id: id}).
		Delete(&variantentity)

	err := query.First(&variantentity).Error

	if err != nil {
		return false, err
	}

	return true, err
}

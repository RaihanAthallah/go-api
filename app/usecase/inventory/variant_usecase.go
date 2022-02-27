package usecase

import (
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
	inventoryrepository "github.com/inventory-management-tokobejo/go-api/app/repository/inventory"
)

type variantUsecase struct {
	Base               baserepository.BaseRepository
	VariantRepository inventoryrepository.VariantRepository
}

type VariantUsecase interface {
	FindByID(id int) (*inventoryentity.Variant, error)
	FindByIDProduct(id_product int) (*inventoryentity.Variant, error)
	List() ([]inventoryentity.Variant, error)
	Create(variant inventoryrequest.VariantCreatedRequest) (interface{}, error)
	DeleteByID(id int) (bool, error)
}

func NewVariantUsecase(
	br baserepository.BaseRepository,
	vr inventoryrepository.VariantRepository,
) VariantUsecase {
	return &variantUsecase{br, vr}
}

func (vu *variantUsecase) FindByID(id int) (*inventoryentity.Variant, error) {
	finddatavariant, err := vu.VariantRepository.FindByID(id)

	return finddatavariant, err
}

func (vu *variantUsecase) FindByIDProduct(id_product int) (*inventoryentity.Variant, error){
	finddatavariant, err := vu.VariantRepository.FindByIDProduct(id_product)

	return finddatavariant, err
}

func (vu *variantUsecase) List() ([]inventoryentity.Variant, error){
	findlistvariant, err := vu.VariantRepository.List()

	return findlistvariant, err
}

func (vu *variantUsecase) Create(variant inventoryrequest.VariantCreatedRequest) (interface{}, error){
	vu.Base.BeginTx()

	entitycreatedvariant := inventoryentity.Variant{
		Id_Product:				variant.Id_Product,
		Variant:				variant.Variant,
		Option_Values:		 	variant.Option_Values,
		
	}
	createdatavariant, err := vu.VariantRepository.Create(entitycreatedvariant)

	if err != nil {
		vu.Base.RollbackTx()
		return false, err
	}

	return createdatavariant, err
}

 func (vu *variantUsecase) DeleteByID(id int) (bool, error) {
	finddatavariant, err := vu.VariantRepository.DeleteByID(id)

	if finddatavariant == true {
		return true, err
	}

	return false, err
}


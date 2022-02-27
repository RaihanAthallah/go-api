package usecase

import (
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
	inventoryrepository "github.com/inventory-management-tokobejo/go-api/app/repository/inventory"
)

type trackingUsecase struct {
	Base               baserepository.BaseRepository
	TrackingRepository inventoryrepository.TrackingRepository
}

type TrackingUsecase interface {
	FindByID(id int) (*inventoryentity.Tracking, error)
	FindByIDProduct(id_product int) (*inventoryentity.Tracking, error)
	List() ([]inventoryentity.Tracking, error)
	Create(tracking inventoryrequest.RequestCreatedTracking) (interface{}, error)
	DeleteByID(id int) (bool, error)
}

func NewTrackingUsecase(
	br baserepository.BaseRepository,
	tr inventoryrepository.TrackingRepository,
) TrackingUsecase {
	return &trackingUsecase{br, tr}
}

func (tu *trackingUsecase) FindByID(id int) (*inventoryentity.Tracking, error) {
	finddatatracking, err := tu.TrackingRepository.FindByID(id)

	return finddatatracking, err
}

func (tu *trackingUsecase) FindByIDProduct(id_product int) (*inventoryentity.Tracking, error){
	finddatatracking, err := tu.TrackingRepository.FindByIDProduct(id_product)

	return finddatatracking, err
}

func (tu *trackingUsecase) List() ([]inventoryentity.Tracking, error){
	findlisttracking, err := tu.TrackingRepository.List()

	return findlisttracking, err
}

func (tu *trackingUsecase) Create(tracking inventoryrequest.RequestCreatedTracking) (interface{}, error){
	tu.Base.BeginTx()

	entitycreatedtracking := inventoryentity.Tracking{
		Id_Product:			tracking.Id_Product,
		Desc:				tracking.Desc,
		Id_Track:		 	tracking.Id_Track,
		
	}
	createdatatracking, err := tu.TrackingRepository.Create(entitycreatedtracking)

	if err != nil {
		tu.Base.RollbackTx()
		return false, err
	}

	return createdatatracking, err
}

func (tu *trackingUsecase) DeleteByID(id int) (bool, error) {
	finddatatracking, err := tu.TrackingRepository.DeleteByID(id)

	if finddatatracking == true {
		return true, err
	}

	return false, err
}


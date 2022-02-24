package registry

import (
	repository "github.com/inventory-management-tokobejo/go-api/app/repository"
	inventoryrepository "github.com/inventory-management-tokobejo/go-api/app/repository/inventory"
	inventoryusecase "github.com/inventory-management-tokobejo/go-api/app/usecase/inventory"
)

func (r *registry) NewProductUsecase() inventoryusecase.ProductUsecase {
	return inventoryusecase.NewProductUsecase(
		repository.NewBaseRepository(r.db),
		inventoryrepository.NewProductRepository(repository.NewBaseRepository(r.db)),
	)
}

func (r *registry) NewTrackingUsecase() inventoryusecase.TrackingUsecase {
	return inventoryusecase.NewTrackingUsecase(
		repository.NewBaseRepository(r.db),
		inventoryrepository.NewTrackingRepository(repository.NewBaseRepository(r.db)),
	)
}

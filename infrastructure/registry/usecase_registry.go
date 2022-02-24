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

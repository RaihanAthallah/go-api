package registry

import (
	inventorycontroller "github.com/inventory-management-tokobejo/go-api/app/controller/inventory"
)

func (r *registry) NewProductController() inventorycontroller.ProductController {
	return inventorycontroller.NewProductController(r.NewProductUsecase())
}

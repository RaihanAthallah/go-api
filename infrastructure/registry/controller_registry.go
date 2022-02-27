package registry

import (
	inventorycontroller "github.com/inventory-management-tokobejo/go-api/app/controller/inventory"
)

func (r *registry) NewProductController() inventorycontroller.ProductController {
	return inventorycontroller.NewProductController(r.NewProductUsecase())
}

func (r *registry) NewTrackingController() inventorycontroller.TrackingController {
	return inventorycontroller.NewTrackingController(r.NewTrackingUsecase())
}

func (r *registry) NewVariantController() inventorycontroller.VariantController {
	return inventorycontroller.NewVariantController(r.NewVariantUsecase())
}
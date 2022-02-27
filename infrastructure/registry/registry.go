package registry

import (
	inventorycontroller "github.com/inventory-management-tokobejo/go-api/app/controller/inventory"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewProductController() inventorycontroller.ProductController
	NewTrackingController() inventorycontroller.TrackingController
	NewVariantController() inventorycontroller.VariantController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

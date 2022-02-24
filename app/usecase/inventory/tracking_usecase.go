package usecase

import (
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	// inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
	inventoryrepository "github.com/inventory-management-tokobejo/go-api/app/repository/inventory"
)

type trackingUsecase struct {
	Base               baserepository.BaseRepository
	TrackingRepository inventoryrepository.TrackingRepository
}

type TrackingUsecase interface {
	FindByID(id int) (*inventoryentity.Tracking, error)
	// FindByIDProduct(id_product int) (*inventoryentity.Tracking, error)
	// List() ([]inventoryentity.Tracking, error)
	// Create(tracking inventoryentity.Tracking) (inventoryentity.Tracking, error)
	// DeleteByID(id int) (bool, error)
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

// func (pu *productUsecase) FindBySKU(sku string) (*inventoryentity.Product, error) {
// 	finddataproduct, err := pu.ProductRepository.FindBySKU(sku)

// 	return finddataproduct, err
// }

// func (pu *productUsecase) List() ([]inventoryentity.Product, error) {
// 	findlistproduct, err := pu.ProductRepository.List()

// 	return findlistproduct, err
// }

// func (pu *productUsecase) Create(product inventoryrequest.RequestCreatedProduct) (interface{}, error) {
// 	pu.Base.BeginTx()

// 	entitycreatedproduct := inventoryentity.Product{
// 		SKU:                 product.SKU,
// 		Id_Supplier:         product.Id_Supplier,
// 		Id_Product_Type:     product.Id_Product_Type,
// 		Id_Brand:            product.Id_Brand,
// 		Product_Description: product.Product_Description,
// 		Weight:              product.Weight,
// 		Id_Weight_Type:      product.Id_Weight_Type,
// 		Initial_Stock:       product.Initial_Stock,
// 		Id_Location:         product.Id_Location,
// 		Initial_Cost:        product.Initial_Cost,
// 		Buy_Price:           product.Buy_Price,
// 		Wholesale_Price:     product.Wholesale_Price,
// 		Retail_Price:        product.Retail_Price,
// 	}
// 	createdataproduct, err := pu.ProductRepository.Create(entitycreatedproduct)

// 	if err != nil {
// 		pu.Base.RollbackTx()
// 		return false, err
// 	}

// 	return createdataproduct, err
// }

// func (pu *productUsecase) DeleteByIDProduct(id int) (bool, error) {
// 	finddataproduct, err := pu.ProductRepository.DeleteByIDProduct(id)

// 	if finddataproduct == true {
// 		return true, err
// 	}

// 	return false, err
// }

// func (pu *productUsecase) DeleteBySKU(sku string) (bool, error) {
// 	finddataproduct, err := pu.ProductRepository.DeeleteBySKU(sku)

// 	if finddataproduct == true {
// 		return true, err
// 	}

// 	return false, err
// }

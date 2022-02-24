package usecase

import (
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	inventoryrequest "github.com/inventory-management-tokobejo/go-api/app/model/request/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
	inventoryrepository "github.com/inventory-management-tokobejo/go-api/app/repository/inventory"
)

type productUsecase struct {
	Base              baserepository.BaseRepository
	ProductRepository inventoryrepository.ProductRepository
}

type ProductUsecase interface {
	FindByIDProduct(id int) (*inventoryentity.Product, error)
	FindBySKU(sku string) (*inventoryentity.Product, error)
	List() ([]inventoryentity.Product, error)
	Create(product inventoryrequest.RequestCreatedProduct) (interface{}, error)
	DeleteByIDProduct(id int) (bool, error)
	DeleteBySKU(sku string) (bool, error)
}

func NewProductUsecase(
	br baserepository.BaseRepository,
	pr inventoryrepository.ProductRepository,
) ProductUsecase {
	return &productUsecase{br, pr}
}

func (pu *productUsecase) FindByIDProduct(id int) (*inventoryentity.Product, error) {
	finddataproduct, err := pu.ProductRepository.FindByIDProduct(id)

	return finddataproduct, err
}

func (pu *productUsecase) FindBySKU(sku string) (*inventoryentity.Product, error) {
	finddataproduct, err := pu.ProductRepository.FindBySKU(sku)

	return finddataproduct, err
}

func (pu *productUsecase) List() ([]inventoryentity.Product, error) {
	findlistproduct, err := pu.ProductRepository.List()

	return findlistproduct, err
}

func (pu *productUsecase) Create(product inventoryrequest.RequestCreatedProduct) (interface{}, error) {
	pu.Base.BeginTx()

	entitycreatedproduct := inventoryentity.Product{
		SKU:                 product.SKU,
		Id_Supplier:         product.Id_Supplier,
		Id_Product_Type:     product.Id_Product_Type,
		Id_Brand:            product.Id_Brand,
		Product_Description: product.Product_Description,
		Weight:              product.Weight,
		Id_Weight_Type:      product.Id_Weight_Type,
		Initial_Stock:       product.Initial_Stock,
		Id_Location:         product.Id_Location,
		Initial_Cost:        product.Initial_Cost,
		Buy_Price:           product.Buy_Price,
		Wholesale_Price:     product.Wholesale_Price,
		Retail_Price:        product.Retail_Price,
	}
	createdataproduct, err := pu.ProductRepository.Create(entitycreatedproduct)

	if err != nil {
		pu.Base.RollbackTx()
		return false, err
	}

	return createdataproduct, err
}

func (pu *productUsecase) DeleteByIDProduct(id int) (bool, error) {
	finddataproduct, err := pu.ProductRepository.DeleteByIDProduct(id)

	if finddataproduct == true {
		return true, err
	}

	return false, err
}

func (pu *productUsecase) DeleteBySKU(sku string) (bool, error) {
	finddataproduct, err := pu.ProductRepository.DeeleteBySKU(sku)

	if finddataproduct == true {
		return true, err
	}

	return false, err
}

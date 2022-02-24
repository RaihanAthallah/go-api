package repository

import (
	"github.com/golang-module/carbon/v2"
	inventoryentity "github.com/inventory-management-tokobejo/go-api/app/model/entity/inventory"
	baserepository "github.com/inventory-management-tokobejo/go-api/app/repository"
)

type productRepository struct {
	base baserepository.BaseRepository
}

type ProductRepository interface {
	FindByIDProduct(id int) (*inventoryentity.Product, error)
	FindBySKU(sku string) (*inventoryentity.Product, error)
	List() ([]inventoryentity.Product, error)
	Create(product inventoryentity.Product) (inventoryentity.Product, error)
	DeleteByIDProduct(id int) (bool, error)
	DeeleteBySKU(sku string) (bool, error)
}

func NewProductRepository(ar baserepository.BaseRepository) ProductRepository {
	return &productRepository{ar}
}

func (pr *productRepository) FindByIDProduct(id int) (*inventoryentity.Product, error) {
	var productentity *inventoryentity.Product
	var baseDb = pr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Product{Id_Product: id})

	err := query.First(&productentity).Error
	if err != nil {
		return nil, err
	}
	return productentity, nil
}

func (pr *productRepository) FindBySKU(sku string) (*inventoryentity.Product, error) {
	var productentity *inventoryentity.Product
	var baseDb = pr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Product{SKU: sku})

	err := query.First(&productentity).Error
	if err != nil {
		return nil, err
	}
	return productentity, nil
}

func (pr *productRepository) List() ([]inventoryentity.Product, error) {
	var productentity []inventoryentity.Product
	var baseDb = pr.base

	query := baseDb.GetDB()

	err := query.Find(&productentity).Error
	if err != nil {
		return nil, err
	}
	return productentity, nil
}

func (pr *productRepository) Create(product inventoryentity.Product) (inventoryentity.Product, error) {
	product.Created_Date_Product = carbon.Now().ToDateTimeString()
	product.Last_Updated_Product = carbon.Now().ToDateTimeString()
	err := pr.base.GetDB().Create(&product).Error
	return product, err
}

func (pr *productRepository) DeleteByIDProduct(id int) (bool, error) {
	var productentity inventoryentity.Product
	var baseDb = pr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Product{Id_Product: id}).
		Delete(&productentity)

	err := query.First(&productentity).Error

	if err != nil {
		return false, err
	}

	return true, err
}

func (pr *productRepository) DeeleteBySKU(sku string) (bool, error) {
	var productentity inventoryentity.Product
	var baseDb = pr.base

	query := baseDb.GetDB().
		Where(&inventoryentity.Product{SKU: sku}).
		Delete(&productentity)

	err := query.First(&productentity).Error

	if err != nil {
		return false, err
	}

	return true, err
}

package repository

import (
	"math"

	"github.com/hayvee-website-development/go-api-hayvee/config"
	"github.com/hayvee-website-development/go-entites-hayvee/entities/response"
	"gorm.io/gorm"
)

type baseRepository struct {
	db *gorm.DB
}

type BaseRepository interface {
	GetDB() *gorm.DB
	BeginTx()
	CommitTx()
	RollbackTx()
	Paginate(value interface{}, pagination *response.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB
	GetSchema() config.Schema
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (br *baseRepository) GetDB() *gorm.DB {
	return br.db
}

func (br *baseRepository) BeginTx() {
	br.db = br.GetDB().Begin()
}

func (br *baseRepository) CommitTx() {
	br.GetDB().Commit()
}

func (br *baseRepository) RollbackTx() {
	br.GetDB().Rollback()
}

func (br *baseRepository) Paginate(value interface{}, pagination *response.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRecords int64
	db.Model(value).Count(&totalRecords)

	pagination.TotalRecords = totalRecords
	pagination.TotalPage = int(math.Ceil(float64(totalRecords) / float64(pagination.GetLimit())))

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}

func (br *baseRepository) GetSchema() config.Schema {
	return config.C.Database.Schema
}

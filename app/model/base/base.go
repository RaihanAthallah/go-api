package base

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID string `json:"id" gorm:"primary_key"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDeleted bool `json:"is_deleted"`
}

//BeforeCreate  will set a UUID rather than numeric ID
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	scope.SetColumn("CreatedAt", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("UpdatedAt", time.Now().Format("2006-01-02 15:04:05"))
	scope.SetColumn("IsDeleted", false)
	return nil
}

//BeforeUpdate will update updatedDate
func (base *Base) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn(base.UpdatedAt, time.Now().Format("2006-01-02 15:04:05"))
	return nil
}
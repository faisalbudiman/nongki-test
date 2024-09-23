package model

import (
	"nongki-test/internal/abstraction"
	"nongki-test/pkg/util/date"

	"gorm.io/gorm"
)

type AddressEntity struct {
	Name   string `json:"name" validate:"required" gorm:"index:idx_address_name,unique"`
	Detail string `json:"detail" validate:"required" gorm:"index:idx_address_detail"`
	UserID uint   `json:"user_id" gorm:"index:idx_user_id"`
}

type AddressFilter struct {
	Code *string `query:"code" filter:"ILIKE"`
	Name *string `query:"name" filter:"ILIKE"`
}

type AddressEntityModel struct {
	// abstraction
	abstraction.Entity

	// entity
	AddressEntity

	// relations
	User *UserEntityModel `json:"province" gorm:"foreignKey:UserID"`

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type AddressFilterModel struct {
	// abstraction
	abstraction.Filter

	// filter
	AddressFilter
}

func (AddressEntityModel) TableName() string {
	return "address"
}

func (m *AddressEntityModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = m.Context.Auth.ID
	return
}

func (m *AddressEntityModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = date.DateTodayLocal()
	m.UpdatedBy = &m.Context.Auth.ID
	return
}

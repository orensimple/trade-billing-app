package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Tabler is interface of GORM table name
type Tabler interface {
	TableName() string
}

// Account is the model of Account
type Account struct {
	ID            uuid.UUID       `gorm:"type:uuid"`
	CurrencyCode  string          `gorm:"type:text;not null"`
	Name          string          `gorm:"type:text"`
	Balance       decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	BlockedAmount decimal.Decimal `gorm:"type:decimal(10,2);not null"`
}

// TableName gets table name of Account
func (Account) TableName() string {
	return "accounts"
}

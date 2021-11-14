package domain

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Account is account in app.
type Account struct {
	ID            uuid.UUID
	CurrencyCode  string
	Name          string
	Balance       decimal.Decimal
	BlockedAmount decimal.Decimal
}

// BlockedRequest struct for blocked money.
type BlockedRequest struct {
	BlockedAmount decimal.Decimal `form:"blocked_amount" json:"blocked_amount" binding:"required"`
}

// PayRequest struct for pay money.
type PayRequest struct {
	PayAmount decimal.Decimal `form:"pay_amount" json:"pay_amount" binding:"required"`
}

// CreateRequest struct for create new account.
type CreateRequest struct {
	CurrencyCode string `form:"currency_code" json:"currency_code"`
	Name         string `form:"name" json:"name"`
}

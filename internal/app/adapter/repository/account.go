package repository

import (
	"errors"

	"github.com/orensimple/trade-billing-app/internal/app/domain"
	"gorm.io/gorm"
)

// Account is the repository of domain.Account
type Account struct {
	repo *gorm.DB
}

func NewAccountRepo(db *gorm.DB) Account {
	return Account{repo: db.Debug()}
}

// Create new account
func (u Account) Create(a *domain.Account) error {
	return u.repo.Create(a).Error
}

// Get account by filter
func (u Account) Get(f *domain.Account) (*domain.Account, error) {
	out := new(domain.Account)

	err := u.repo.Where(f).Take(out).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("account not found")
		}

		return nil, errors.New("failed get account")
	}

	return out, nil
}

// Update account info by id
func (u Account) Update(a *domain.Account) error {
	return u.repo.Save(&a).Error
}

// Delete account by id
func (u Account) Delete(f *domain.Account) error {
	return u.repo.Delete(&f).Error
}

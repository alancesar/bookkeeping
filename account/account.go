package account

import (
	"errors"
	"fmt"
	"time"
)

const (
	SalaryIncoming    = "SALARY_INCOMING"
	RevenueIncoming   = "REVENUE_INCOMING"
	InvestmentAccount = "INVESTMENT_ACCOUNT"
	CheckingAccount   = "CHECKING_ACCOUNT"
	CreditCard        = "CREDIT_CARD"
	Loan              = "LOAN"
	FixedAsset        = "FIXED_ASSET"
	Service           = "SERVICE"
	Consumable        = "CONSUMABLE"
	Depreciation      = "DEPRECIATION"
	Fee               = "FEE"
)

var (
	ErrInvalidType        = errors.New("invalid type")
	ErrInvalidTransaction = errors.New("invalid transaction")
)

type (
	Type string

	Account struct {
		ID        uint
		Type      Type
		Name      string
		CreatedAt time.Time
	}
)

func (a Account) ValidTo() []Type {
	switch a.Type {
	case SalaryIncoming, InvestmentAccount, Loan:
		return []Type{CheckingAccount}
	case RevenueIncoming:
		return []Type{InvestmentAccount}
	case CheckingAccount:
		return []Type{CheckingAccount, InvestmentAccount, CreditCard, Loan, FixedAsset, Service, Consumable, Fee}
	case CreditCard:
		return []Type{FixedAsset, Service, Consumable}
	case FixedAsset:
		return []Type{Depreciation}
	case Service, Consumable, Fee:
		return []Type{}
	default:
		return []Type{}
	}
}

func (a Account) ValidateTo(to Account) error {
	if a.Type == "" {
		return ErrInvalidType
	}

	validTypes := a.ValidTo()
	for _, validType := range validTypes {
		if to.Type == validType {
			return nil
		}
	}

	return fmt.Errorf("%w: %s type cannot go to %s type", ErrInvalidTransaction, a.Type, to.Type)
}

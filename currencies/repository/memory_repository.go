package repository

import (
	"currency_converter/currencies"
	"errors"
)

type MemoryCurrenciesRepo struct {
	currencyList map[string]currencies.Currency
}

func NewMemoryCurrenciesRepo() *MemoryCurrenciesRepo {
	return &MemoryCurrenciesRepo{currencyList: map[string]currencies.Currency{
		"COP": {
			Name:        "Colombian Peso",
			ShortName:   "COP",
			USDRelation: 0.000268732,
		},
		"EUR": {
			Name:        "Euro",
			ShortName:   "EUR",
			USDRelation: 1.08020,
		},
		"USD": {
			Name:        "USD",
			ShortName:   "Dollar",
			USDRelation: 1.00,
		},
	}}
}

func (m MemoryCurrenciesRepo) GetCurrency(shortName string) (*currencies.Currency, error) {
	if val, ok := m.currencyList[shortName]; ok {
		return &val, nil
	}
	return nil, errors.New("currency not found")
}

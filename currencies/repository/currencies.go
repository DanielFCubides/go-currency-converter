package repository

import "currency_converter/currencies"

type CurrencyRepository interface {
	GetCurrency(shortName string) (*currencies.Currency, error)
}

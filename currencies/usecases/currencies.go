package usecases

import "currency_converter/currencies"

type CurrenciesHandler interface {
	GetAllConversion(shortName string)(*[]currencies.Currency, error)
	GetConversion( fromShortName string, sourceAmount float64, toShortName string)(*float64, error)
}

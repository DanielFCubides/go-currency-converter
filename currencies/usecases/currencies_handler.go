package usecases

import (
	"currency_converter/currencies"
	"currency_converter/currencies/repository"
)

const USD = "USD"

type CurrenciesHandlerImpl struct {
	repository repository.CurrencyRepository
}

func NewCurrenciesHandlerImpl(repository repository.CurrencyRepository) *CurrenciesHandlerImpl {
	return &CurrenciesHandlerImpl{repository: repository}
}

func (c CurrenciesHandlerImpl) GetAllConversion(shortName string) (*[]currencies.Currency, error) {
	currenciesList := make([]currencies.Currency, 0)
	currency,err := c.repository.GetCurrency(shortName)
	if err != nil {
		return nil, err
	}
	currenciesList = append(currenciesList, *currency)
	return &currenciesList, nil
}

func (c CurrenciesHandlerImpl) GetConversion(fromShortName string, sourceAmount float64, toShortName string) (*float64, error) {
	currencySource, err := c.repository.GetCurrency(fromShortName)
	if err != nil {
		return nil, err
	}
	currencyTarget, err := c.repository.GetCurrency(toShortName)
	if err != nil {
		return nil, err
	}
	result := 0.0

	sourceUSD := sourceAmount * currencySource.USDRelation
	if toShortName == USD {
		result = sourceUSD
	} else {
		result = sourceUSD / currencyTarget.USDRelation
	}
	return &result, nil
}

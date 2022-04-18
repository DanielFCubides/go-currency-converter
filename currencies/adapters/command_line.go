package adapters

import (
	"currency_converter/currencies/usecases"
	"flag"
	"fmt"
)

type CMDAdapter struct {
	usecase usecases.CurrenciesHandler
}

func NewCMDAdapter(usecase usecases.CurrenciesHandler) *CMDAdapter {
	return &CMDAdapter{usecase: usecase}
}

func (c *CMDAdapter) RunCMD() {
	sourceCurrency := flag.String("from", "COP", "currency for make conversion")
	targetCurrency := flag.String("to", "USD", "currency to make conversion")
	amount := flag.Float64("amount", 100000.0, "amount to convert to target currency")
	flag.Parse()

	conversion, err := c.usecase.GetConversion(*sourceCurrency, *amount, *targetCurrency)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%.3f %s is equivalent to %.3f %s \n", *amount, *sourceCurrency, *conversion, *targetCurrency)

}

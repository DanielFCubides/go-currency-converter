package main

import (
	"currency_converter/currencies/adapters"
	"currency_converter/currencies/repository"
	"currency_converter/currencies/usecases"
)

func main() {
	repo := repository.NewMemoryCurrenciesRepo()
	usecase := usecases.NewCurrenciesHandlerImpl(repo)
	adapter := adapters.NewCMDAdapter(usecase)
	adapter.RunCMD()

}

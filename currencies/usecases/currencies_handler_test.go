package usecases

import (
	"currency_converter/currencies/repository"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UseCaseSuite struct {
	suite.Suite
	handler CurrenciesHandler
}

func TestUseCaseInit(t *testing.T){
	suite.Run(t, new(UseCaseSuite))
}

func (s *UseCaseSuite) SetupSuite() {
	s.handler = NewCurrenciesHandlerImpl(repository.NewMemoryCurrenciesRepo())
}

func (s *UseCaseSuite) TestCurrenciesHandlerImpl_GetConversion_toUSD_success() {
	conversion, err := s.handler.GetConversion("COP", 50000.0, "USD")
	if err != nil {
		s.Fail("error at conversion")
	}
	s.Assert().Equal("13.44", fmt.Sprintf("%.2f",*conversion))
	s.Assert().NoError(err)

}

func (s *UseCaseSuite) TestCurrenciesHandlerImpl_GetConversion_toEUR_success() {
	conversion, err := s.handler.GetConversion("COP", 50000.0, "EUR")
	if err != nil {
		s.Fail("error at conversion")
	}
	s.Assert().Equal("12.44", fmt.Sprintf("%.2f",*conversion))
	s.Assert().NoError(err)
}

func (s *UseCaseSuite) TestCurrenciesHandlerImpl_GetConversion_toMXN_fail() {
	_, err := s.handler.GetConversion("COP", 50000.0, "MXN")
	s.Assert().Error(err)

}

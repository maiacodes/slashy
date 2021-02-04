package commands

import (
	"fmt"
	"github.com/maiacodes/slashy/interaction"
	"math"
	"strconv"
	"strings"
)

type exchangeRates struct {
	Rates map[string]float32 `json:"rates"`
}

func cmdCurrency(e *interaction.Event) *interaction.EventCallback {
	// Fetch latest rates
	var response exchangeRates
	err := fetch("https://api.exchangeratesapi.io/latest", "GET", nil, &response)
	if err != nil {
		return e.Error("Cannot fetch exchange rates! `" + err.Error() + "`")
	}

	// Validate FROM currency
	from := strings.ToUpper(e.Data.Options[0].Value)
	eurFromRate, ok := response.Rates[from]
	if from == "EUR" {
		ok = true
		eurFromRate = 1
	}
	if !ok {
		return e.Error("From currency not found!")
	}

	// Validate AMOUNT
	amountString := e.Data.Options[1].Value
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		return e.Error("Parameter error")
	}
	eur := float32(amount) / eurFromRate

	// Validate TO currency
	to := strings.ToUpper(e.Data.Options[2].Value)
	eurToRate, ok := response.Rates[to]
	// Override for EUR because it's the base converting currency
	if to == "EUR" {
		ok = true
		eurToRate = 1
	}
	if !ok {
		return e.Error("To currency not found!")
	}
	converted := eur * eurToRate

	return e.Reply(fmt.Sprintf("**💵  |  %v %v = %v %v**", amount, from, math.Floor(float64(converted)*100)/100, to))
}

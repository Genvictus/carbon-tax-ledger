package controllers

import (
	"carbon-tax-ledger/contract"
	"carbon-tax-ledger/pkg/repository"
	"math/rand"
	"strconv"
)

func MintCarbonToken() (int, error) {
	contract, err := contract.GetContract(repository.ChainCodeName["CT"])
	if err != nil {
		return 0, err
	}

	amount := rand.Intn(20)
	_, err = contract.SubmitTransaction("Mint", strconv.Itoa(amount), "")
	if err != nil {
		return 0, err
	}

	return amount, nil
}

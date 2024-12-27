/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"log"

	"github.com/Genvictus/carbon-tax-ledger/token-erc20/co2-token/chaincode-go/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

func main() {
	tokenChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating token-erc-20 chaincode: %v", err)
	}

	if err := tokenChaincode.Start(); err != nil {
		log.Panicf("Error starting token-erc-20 chaincode: %v", err)
	}
}

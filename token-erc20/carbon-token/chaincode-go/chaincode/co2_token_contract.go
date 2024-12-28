package chaincode

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// Define access control constants
const minterMSPID = "Org1MSP"

// Define key names for options
const co2TokenKey = "co2_token"

const nameKey = "co2_token_key"
const symbolKey = "co2_token_symbol"
const decimalsKey = "2"
const totalSupplyKey = "co2_token_total"

// Define objectType names for prefix
const allowancePrefix = co2TokenKey

// CO2Contract provides functions for transferring tokens between accounts
type CO2Contract struct {
	contractapi.Contract
}

// event provides an organized struct for emitting events
type event struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Value int    `json:"value"`
}

// Mint creates new tokens and adds them to minter's account balance
// This function triggers a Transfer event
func (s *CO2Contract) Mint(ctx contractapi.TransactionContextInterface, amount int, recipient string) error {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return errors.New("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to mint new tokens
	err = verifyMinter(ctx, "client is not authorized to mint new tokens")
	if err != nil {
		return err
	}

	// Get ID of submitting client identity if empty
	var account string
	if recipient == "" {
		account, err = getClientID(ctx)
		if err != nil {
			return err
		}
	} else {
		account = recipient
	}

	if amount <= 0 {
		return errors.New("mint amount must be a positive integer")
	}

	currentBalanceBytes, err := ctx.GetStub().GetState(account)
	if err != nil {
		return fmt.Errorf("failed to read account %s from world state: %v", account, err)
	}

	var currentBalance int

	// If minter current balance doesn't yet exist, we'll create it with a current balance of 0
	if currentBalanceBytes == nil {
		currentBalance = 0
	} else {
		currentBalance, _ = strconv.Atoi(string(currentBalanceBytes)) // Error handling not needed since Itoa() was used when setting the account balance, guaranteeing it was an integer.
	}

	updatedBalance, err := add(currentBalance, amount)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(account, []byte(strconv.Itoa(updatedBalance)))
	if err != nil {
		return err
	}

	// Update the totalSupply
	totalSupplyBytes, err := ctx.GetStub().GetState(totalSupplyKey)
	if err != nil {
		return fmt.Errorf("failed to retrieve total token supply: %v", err)
	}

	var totalSupply int

	// If no tokens have been minted, initialize the totalSupply
	if totalSupplyBytes == nil {
		totalSupply = 0
	} else {
		totalSupply, _ = strconv.Atoi(string(totalSupplyBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.
	}

	// Add the mint amount to the total supply and update the state
	totalSupply, err = add(totalSupply, amount)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(totalSupplyKey, []byte(strconv.Itoa(totalSupply)))
	if err != nil {
		return err
	}

	// Emit the Transfer event
	transferEvent := event{"0x0", account, amount}
	transferEventJSON, err := json.Marshal(transferEvent)
	if err != nil {
		return fmt.Errorf("failed to obtain JSON encoding: %v", err)
	}
	err = ctx.GetStub().SetEvent("Transfer", transferEventJSON)
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	log.Printf("minter account %s balance updated from %d to %d", account, currentBalance, updatedBalance)

	return nil
}

// Burn redeems tokens the minter's account balance
// This function triggers a Transfer event
// Recipient uses empty string ("") as a special value to denote minting to the caller's account
func (s *CO2Contract) Burn(ctx contractapi.TransactionContextInterface, amount int, recipient string) error {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return errors.New("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to burn new tokens
	err = verifyMinter(ctx, "client is not authorized to burn tokens")
	if err != nil {
		return err
	}

	// Get ID of submitting client identity if empty
	var account string
	if recipient == "" {
		account, err = getClientID(ctx)
		if err != nil {
			return err
		}
	} else {
		account = recipient
	}

	if amount <= 0 {
		return errors.New("burn amount must be a positive integer")
	}

	currentBalance, err := getAccountBalance(ctx, account)
	if err != nil {
		return err
	}

	updatedBalance, err := sub(currentBalance, amount)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(account, []byte(strconv.Itoa(updatedBalance)))
	if err != nil {
		return err
	}

	// Update the totalSupply
	totalSupplyBytes, err := ctx.GetStub().GetState(totalSupplyKey)
	if err != nil {
		return fmt.Errorf("failed to retrieve total token supply: %v", err)
	}

	// If no tokens have been minted, throw error
	if totalSupplyBytes == nil {
		return errors.New("totalSupply does not exist")
	}

	totalSupply, _ := strconv.Atoi(string(totalSupplyBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.

	// Subtract the burn amount to the total supply and update the state
	totalSupply, err = sub(totalSupply, amount)
	if err != nil {
		return err
	}

	err = ctx.GetStub().PutState(totalSupplyKey, []byte(strconv.Itoa(totalSupply)))
	if err != nil {
		return err
	}

	// Emit the Transfer event
	transferEvent := event{account, "0x0", amount}
	transferEventJSON, err := json.Marshal(transferEvent)
	if err != nil {
		return fmt.Errorf("failed to obtain JSON encoding: %v", err)
	}
	err = ctx.GetStub().SetEvent("Transfer", transferEventJSON)
	if err != nil {
		return fmt.Errorf("failed to set event: %v", err)
	}

	log.Printf("account %s balance updated from %d to %d", account, currentBalance, updatedBalance)

	return nil
}

func (s *CO2Contract) Pay(ctx contractapi.TransactionContextInterface, amount int) error {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return errors.New("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	clientID, err := getClientID(ctx)
	if err != nil {
		return err
	}

	// Check if the amount paid is more than the balance needed to be paid
	tokenAmount, err := getAccountBalance(ctx, clientID)
	if err != nil {
		return err
	}
	if tokenAmount < amount {
		return fmt.Errorf("CO2 token is less than amount paid")
	}

	BurnArg := [][]byte{[]byte("Burn"), []byte(strconv.Itoa(amount)), []byte("")}
	response := ctx.GetStub().InvokeChaincode("primary_wallet", BurnArg, ctx.GetStub().GetChannelID())

	if response.GetStatus() == 200 {
		s.Burn(ctx, amount, "")
	} else {
		return fmt.Errorf("failed to burn from primary wallet: %v", response.GetMessage())
	}

	return nil
}

// BalanceOf returns the balance of the given account
func (s *CO2Contract) BalanceOf(ctx contractapi.TransactionContextInterface, account string) (int, error) {

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege manage tokens
	// Only the minter is allowed to query the balance of other accounts
	err := verifyMinter(ctx, "Client is not allowed to query other account's balance")
	if err != nil {
		return 0, err
	}
	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return 0, fmt.Errorf("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	balance, _ := getAccountBalance(ctx, account)

	return balance, nil
}

// ClientAccountBalance returns the balance of the requesting client's account
func (s *CO2Contract) ClientAccountBalance(ctx contractapi.TransactionContextInterface) (int, error) {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return 0, fmt.Errorf("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Get ID of submitting client identity
	clientID, err := getClientID(ctx)
	if err != nil {
		return 0, err
	}

	balance, _ := getAccountBalance(ctx, clientID)

	return balance, nil
}

// ClientAccountID returns the id of the requesting client's account
// In this implementation, the client account ID is the clientId itself
// Users can use this function to get their own account id, which they can then give to others as the payment address
func (s *CO2Contract) ClientAccountID(ctx contractapi.TransactionContextInterface) (string, error) {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return "", fmt.Errorf("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Get ID of submitting client identity
	clientAccountID, err := getClientID(ctx)
	if err != nil {
		return "", err
	}

	return clientAccountID, nil
}

// TotalSupply returns the total token supply
func (s *CO2Contract) TotalSupply(ctx contractapi.TransactionContextInterface) (int, error) {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return 0, fmt.Errorf("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Retrieve total supply of tokens from state of smart contract
	totalSupplyBytes, err := ctx.GetStub().GetState(totalSupplyKey)
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve total token supply: %v", err)
	}

	var totalSupply int

	// If no tokens have been minted, return 0
	if totalSupplyBytes == nil {
		totalSupply = 0
	} else {
		totalSupply, _ = strconv.Atoi(string(totalSupplyBytes)) // Error handling not needed since Itoa() was used when setting the totalSupply, guaranteeing it was an integer.
	}

	log.Printf("TotalSupply: %d tokens", totalSupply)

	return totalSupply, nil
}

// Name returns a descriptive name for fungible tokens in this contract
// returns {String} Returns the name of the token

func (s *CO2Contract) Name(ctx contractapi.TransactionContextInterface) (string, error) {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return "", fmt.Errorf("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	bytes, err := ctx.GetStub().GetState(nameKey)
	if err != nil {
		return "", fmt.Errorf("failed to get Name bytes: %s", err)
	}

	return string(bytes), nil
}

// Symbol returns an abbreviated name for fungible tokens in this contract.
// returns {String} Returns the symbol of the token

func (s *CO2Contract) Symbol(ctx contractapi.TransactionContextInterface) (string, error) {

	// Check if contract has been intilized first
	initialized, err := checkInitialized(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return "", fmt.Errorf("Contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	bytes, err := ctx.GetStub().GetState(symbolKey)
	if err != nil {
		return "", fmt.Errorf("failed to get Symbol: %v", err)
	}

	return string(bytes), nil
}

// Set information for a token and intialize contract.
// param {String} name The name of the token
// param {String} symbol The symbol of the token
// param {String} decimals The decimals used for the token operations
func (s *CO2Contract) Initialize(ctx contractapi.TransactionContextInterface, name string, symbol string, decimals string) (bool, error) {

	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to intitialize contract
	err := verifyMinter(ctx, "client is not authorized to initialize contract")
	if err != nil {
		return false, err
	}

	// Check contract options are not already set, client is not authorized to change them once intitialized
	bytes, err := ctx.GetStub().GetState(nameKey)
	if err != nil {
		return false, fmt.Errorf("failed to get Name: %v", err)
	}
	if bytes != nil {
		return false, fmt.Errorf("contract options are already set, client is not authorized to change them")
	}

	err = ctx.GetStub().PutState(nameKey, []byte(name))
	if err != nil {
		return false, fmt.Errorf("failed to set token name: %v", err)
	}

	err = ctx.GetStub().PutState(symbolKey, []byte(symbol))
	if err != nil {
		return false, fmt.Errorf("failed to set symbol: %v", err)
	}

	err = ctx.GetStub().PutState(decimalsKey, []byte(decimals))
	if err != nil {
		return false, fmt.Errorf("failed to set token name: %v", err)
	}

	return true, nil
}

// Set information for a token and intialize contract using default values.
func (s *CO2Contract) Init(ctx contractapi.TransactionContextInterface) (bool, error) {
	return s.Initialize(ctx, "co2_token", "co2_token", "2")
}

// Helper Functions

// add two number checking for overflow
func add(b int, q int) (int, error) {

	// Check overflow
	var sum int
	sum = q + b

	if (sum < q || sum < b) == (b >= 0 && q >= 0) {
		return 0, fmt.Errorf("Math: addition overflow occurred %d + %d", b, q)
	}

	return sum, nil
}

// Checks that contract options have been already initialized
func checkInitialized(ctx contractapi.TransactionContextInterface) (bool, error) {
	tokenName, err := ctx.GetStub().GetState(nameKey)
	if err != nil {
		return false, fmt.Errorf("failed to get token name: %v", err)
	}

	if tokenName == nil {
		return false, nil
	}

	return true, nil
}

// sub two number checking for overflow
func sub(b int, q int) (int, error) {

	// sub two number checking
	if q <= 0 {
		return 0, fmt.Errorf("Error: the subtraction number is %d, it should be greater than 0", q)
	}
	if b < q {
		return 0, fmt.Errorf("Error: the number %d is not enough to be subtracted by %d", b, q)
	}
	var diff int
	diff = b - q

	return diff, nil
}

// Get the client ID for the CO2 token contract
func getClientID(ctx contractapi.TransactionContextInterface) (string, error) {
	// Get ID of submitting client identity
	clientID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return "", fmt.Errorf("failed to get client id: %v", err)
	}

	return (nameKey + clientID), nil
}

// Verify if caller is authorized minter, return error if any checks fail, return nil if authorized
func verifyMinter(ctx contractapi.TransactionContextInterface, unauthorizedMessage string) error {
	// Check minter authorization - this sample assumes Org1 is the central banker with privilege to mint new tokens
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed to get MSPID: %v", err)
	}
	if clientMSPID != minterMSPID {
		return errors.New(unauthorizedMessage)
	}

	return nil
}

func getAccountBalance(ctx contractapi.TransactionContextInterface, account string) (int, error) {
	currentBalanceBytes, err := ctx.GetStub().GetState(account)
	if err != nil {
		return 0, fmt.Errorf("failed to read account %s from world state: %v", account, err)
	}

	// Check if minter current balance exists
	if currentBalanceBytes == nil {
		return 0, errors.New("The balance does not exist")
	}

	currentBalance, _ := strconv.Atoi(string(currentBalanceBytes)) // Error handling not needed since Itoa() was used when setting the account balance, guaranteeing it was an integer.

	return currentBalance, nil
}

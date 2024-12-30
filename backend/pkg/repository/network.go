package repository

import (
	"os"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

const (
	ChannelName string = "mychannel"
	CryptoPath  string = "../test-network/organizations/peerOrganizations/org1.example.com"
	CertPath    string = CryptoPath + "/users/minter@org1.example.com/msp/signcerts"
	KeyPath     string = CryptoPath + "/users/minter@org1.example.com/msp/keystore"
	TLSCertPath string = CryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
)

var (
	PeerEndpoint map[string]string = map[string]string{
		"Org1MSP": os.Getenv("PEER_ENDPOINT_ORG1MSP"),
		"Org2MSP": os.Getenv("PEER_ENDPOINT_ORG2MSP"),
	}
	GatewayPeer map[string]string = map[string]string{
		"Org1MSP": "peer0.org1.example.com",
		"Org2MSP": "peer0.org2.example.com",
	}
	ChainCodeName map[string]string = map[string]string{
		"CT": "carbon_tax_tokens",
		"WT": "primary_wallet",
	}
)

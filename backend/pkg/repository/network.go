package repository

const (
	ChannelName string = "mychannel"
)

var (
	PeerEndpoint map[string]string = map[string]string{
		"Org1MSP": "localhost:7051",
		"Org2MSP": "localhost:9051",
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

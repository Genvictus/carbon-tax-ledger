package contract

import (
	"carbon-tax-ledger/pkg/repository"
	"crypto/x509"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/hash"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// GetContract returns a new contract for the specified chaincode.
func GetContract(c *fiber.Ctx, chaincodeName string) (*client.Contract, error) {
	channelName := repository.ChannelName

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection, err := newGrpcConnection(c)
	if err != nil {
		return nil, err
	}
	id, err := newIdentity(c)
	if err != nil {
		return nil, err
	}
	sign, err := newSign(c)
	if err != nil {
		return nil, err
	}

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Gateway connection: %w", err)
	}

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	return contract, nil
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection(c *fiber.Ctx) (*grpc.ClientConn, error) {
	tlsCertPath := filepath.Join(repository.SessionDir, c.Get("session-id"), "tlsCert.pem")
	gatewayPeer := repository.GatewayPeer[c.Get("msp-id")]
	peerEndpoint := repository.PeerEndpoint[c.Get("msp-id")]

	certificatePEM, err := os.ReadFile(tlsCertPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read TLS certifcate file: %w", err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate from PEM: %w", err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.NewClient(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	return connection, nil
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity(c *fiber.Ctx) (*identity.X509Identity, error) {
	certPath := filepath.Join(repository.SessionDir, c.Get("session-id"), "cert.pem")
	mspID := c.Get("msp-id")

	certificatePEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate from PEM: %w", err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
	if err != nil {
		return nil, fmt.Errorf("failed to create X.509 identity: %w", err)
	}

	return id, nil
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign(c *fiber.Ctx) (identity.Sign, error) {
	keyPath := filepath.Join(repository.SessionDir, c.Get("session-id"), "key.pem")

	privateKeyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("failed to create private key from PEM: %w", err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create private key sign: %w", err)
	}

	return sign, nil
}

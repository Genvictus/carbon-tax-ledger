cd ${PWD}/..

# Start the test network using default config
pushd ${PWD}/test-network
# Clean any previous network setup
./network.sh down
# Start the new test network
./network.sh up createChannel -ca -c mychannel
popd

# Install the chaincodes
# Primary wallet chaincode
pushd ${PWD}/token-erc20/primary-wallet
./deploy.sh
popd
# Carbon token chaincode
pushd ${PWD}/token-erc20/carbon-token
./deploy.sh
popd

# Pull the frontend, backend, and oracle services image for docker
docker pull docker.io/library/node:22
docker pull docker.io/library/golang:1.23-alpine
docker pull docker.io/library/alpine:3.20.3

# Start the frontend, backend, and oracle services
docker compose up -d

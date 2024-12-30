#!/bin/bash

pushd ${PWD}

if [ ! -d ./install-fabric.sh ]; then
    curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh
fi

# Install the client binary
if [[ ! -d ./bin || ! -d ./builders || ! -d ./config ]]; then
    ./install-fabric.sh --fabric-version 2.5.10 binary docker
fi

# Export the path
export WORKSHOP_PATH=$(pwd)
export PATH=${WORKSHOP_PATH}/bin:$PATH
export FABRIC_CFG_PATH=${WORKSHOP_PATH}/config

popd

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

pushd ${PWD}/test-network

export PATH=${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=$PWD/../config/

export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.example.com/

# You can register a new minter client identity using the fabric-ca-client tool:
fabric-ca-client register --caname ca-org1 --id.name minter --id.secret minterpw --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"

# You can now generate the identity certificates and MSP folder by providing the minter's enroll name and secret to the enroll command:
fabric-ca-client enroll -u https://minter:minterpw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.example.com/users/minter@org1.example.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"

# Run the command below to copy the Node OU configuration file into the minter identity MSP folder.
cp "${PWD}/organizations/peerOrganizations/org1.example.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.example.com/users/minter@org1.example.com/msp/config.yaml"

# Shift back to the Org1 terminal, we'll set the following environment variables to operate the peer CLI as the minter identity from Org1.
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/minter@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051
export TARGET_TLS_OPTIONS=(-o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt")

# Init the chaincodes
peer chaincode invoke "${TARGET_TLS_OPTIONS[@]}" -C mychannel -n primary_wallet -c '{"function":"Init","Args":[]}'
peer chaincode invoke "${TARGET_TLS_OPTIONS[@]}" -C mychannel -n carbon_tax_tokens -c '{"function":"Init","Args":[]}'

popd

docker compose up --build -d
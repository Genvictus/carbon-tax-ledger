#!/bin/bash

cd ../../../test-network/
./network.sh deployCC -ccn primary_wallet -ccp ../token-erc20/primary-wallet/chaincode-go/ -ccl go

peer chaincode invoke "${TARGET_TLS_OPTIONS[@]}" -C mychannel -n primary_wallet -c '{"function":"Init","Args":[]}'

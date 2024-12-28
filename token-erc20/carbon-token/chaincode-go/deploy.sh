#!/bin/bash

cd ../../../test-network/
./network.sh deployCC -ccn carbon_tax_tokens -ccp ../token-erc20/carbon-token/chaincode-go/ -ccl go

peer chaincode invoke "${TARGET_TLS_OPTIONS[@]}" -C mychannel -n primary_wallet -c '{"function":"Init","Args":[]}'

#!/bin/bash

cd ${PWD}/..

# Down the test network
pushd ${PWD}/test-network
# Clean any previous network setup
./network.sh down
popd

docker compose down

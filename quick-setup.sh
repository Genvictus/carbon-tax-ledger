#!/bin/bash

# Install the client binary
if [[ ! -d ./bin || ! -d ./builders || ! -d ./config ]]; then
    ./install-fabric.sh --fabric-version 2.5.10 binary
fi

# Export the path
export WORKSHOP_PATH=$(pwd)
export PATH=${WORKSHOP_PATH}/bin:$PATH
export FABRIC_CFG_PATH=${WORKSHOP_PATH}/config

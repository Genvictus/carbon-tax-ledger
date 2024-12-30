#!/usr/bin/env bash

SUCCESS="✅"
WARN="⚠️ "
EXIT=0

if ! command -v docker &> /tmp/cmdpath
then
    echo "${WARN} Please install Docker; suggested install commands:"
    EXIT=1
else
    echo -e "${SUCCESS} Docker found:\t$(cat /tmp/cmdpath)"
fi

# Install jq
if ! command -v jq &> /tmp/cmdpath
then
  echo "${WARN} Please install jq; suggested install commands:"
  echo "sudo apt-update && sudo apt-install -y jq"
  EXIT=1
else
  echo -e "${SUCCESS} jq found:\t\t$(cat /tmp/cmdpath)"
fi

if ! command -v peer &> /tmp/cmdpath
then
  echo "${WARN} Please install the peer; suggested install commands:"
  echo "source client-setup.sh"
  EXIT=1
else
  echo -e "${SUCCESS} peer found:\t\t$(cat /tmp/cmdpath)"

  # double-check that the peer binary is compiled for the correct arch.  This can occur when installing fabric
  # binaries into a multipass VM, then running the Linux binaries from a Mac or windows Host OS via the volume share.
  peer version &> /dev/null
  rc=$?
  if [ $rc -ne 0 ]; then
    echo -e "${WARN}  Could not execute peer.  Was it compiled for the correct architecture?"
    peer version
  fi
fi

# tests if varname is defined in the env AND it's an existing directory
function must_declare() {
  local varname=$1

  if [[ ! -d ${!varname} ]]; then
    echo "${WARN} ${varname} must be set to a directory"
    EXIT=1

  else
    echo -e "${SUCCESS} ${varname}:\t${!varname}"
  fi
}

must_declare "FABRIC_CFG_PATH"
must_declare "WORKSHOP_PATH"

rm /tmp/cmdpath &> /dev/null

exit $EXIT
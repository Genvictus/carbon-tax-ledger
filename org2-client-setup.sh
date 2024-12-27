export PATH=${PWD}/bin:${PWD}/test-network:$PATH
export FABRIC_CA_CLIENT_HOME=${PWD}/test-network/organizations/peerOrganizations/org2.example.com/

fabric-ca-client register --caname ca-org2 --id.name recipient --id.secret recipientpw --id.type client --tls.certfiles "${PWD}/test-network/organizations/fabric-ca/org2/tls-cert.pem"

fabric-ca-client enroll -u https://recipient:recipientpw@localhost:8054 --caname ca-org2 -M "${PWD}/test-network/organizations/peerOrganizations/org2.example.com/users/recipient@org2.example.com/msp" --tls.certfiles "${PWD}/test-network/organizations/fabric-ca/org2/tls-cert.pem"

cp "${PWD}/test-network/organizations/peerOrganizations/org2.example.com/msp/config.yaml" "${PWD}/test-network/organizations/peerOrganizations/org2.example.com/users/recipient@org2.example.com/msp/config.yaml"

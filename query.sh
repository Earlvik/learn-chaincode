#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
ORANGE='\033[0;33m'
NC='\033[0m'
while IFS=' ' read -ra line || [[ -n "$line" ]]; do
    if [ "${line[1]}" = "$1" ]; then
    	ID="${line[2]}"
    fi	
done < "deploylog.txt"

ARGS=''
for var in ${@:3}
do
	ARGS+="\"$var\", "
done
ARGS=${ARGS%,*}
RESULT=$(curl -s -X POST --header "Content-Type: application/json" --header "Accept: application/json" -d "{
     \"jsonrpc\": \"2.0\",
     \"method\": \"query\",
     \"params\": {
         \"type\": 1,
         \"chaincodeID\": {
             \"name\": \"$ID\"
         },
         \"ctorMsg\": {
             \"function\": \"$2\",
             \"args\": [$ARGS]
         },
         \"secureContext\": \"admin\"
     },
     \"id\": 2
 }" "https://1acda275b31041d89efd8a04b9bac2ea-vp0.us.blockchain.ibm.com:5004/chaincode")
 ST=$(echo $RESULT | jq -r '.result.status')
 MSG=$(echo $RESULT | jq -r '.result.message')
 echo -e "${ORANGE}Status: $ST ${NC}"
 echo -e "${GREEN}Result: $MSG ${NC}"
 
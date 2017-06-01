#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
ORANGE='\033[0;33m'
NC='\033[0m'
echo "-----------------------------------------------------------------------------------------------------------------------------------------------------------------" >> deploylog.txt
echo -e "${ORANGE}Building folder $1${NC}"
cd $1
go build ./
cd ..
echo -e "${ORANGE}Pushing to github${NC}"
git add -A
DATE=`date '+%d/%m/%Y_%H:%M:%S'`
git commit -m "Deploy $1 $DATE"
git push
echo -e "${ORANGE}Deploying to Bluemix${NC}"

ID=$(curl -s -X POST --header "Content-Type: application/json" --header "Accept: application/json" -d "{
     \"jsonrpc\": \"2.0\",
     \"method\": \"deploy\",
     \"params\": {
         \"type\": 1,
         \"chaincodeID\": {
             \"path\": \"https://github.com/Earlvik/learn-chaincode/$1\"
         },
         \"ctorMsg\": {
             \"function\": \"init\",
             \"args\": [
                 \"hi there\"
             ]
         },
         \"secureContext\": \"admin\"
     },
     \"id\": 1
 }" "https://1acda275b31041d89efd8a04b9bac2ea-vp0.us.blockchain.ibm.com:5004/chaincode" | jq -r '.result.message')
echo "$DATE $1 $ID" >> deploylog.txt
git add -A
git commit --amend -m "Deploy $1 $DATE"
git push -f
echo -e "chaincodeID: ${GREEN}${ID}${NC}"
#!/bin/bash
echo "DATE       TIME     NAME                            CHAINCODE ID" > deploylog.txt
./deploy.sh car
./deploy.sh wash
./deploy.sh uber
./deploy.sh toll
./deploy.sh park
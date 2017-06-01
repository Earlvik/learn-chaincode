#!/bin/bash
echo "DATE       TIME     NAME                            CHAINCODE ID" > deploylog.txt
./deploy.sh start
./deploy.sh car
./deploy.sh park
./deploy.sh wash
./deploy.sh uber
./deploy.sh toll
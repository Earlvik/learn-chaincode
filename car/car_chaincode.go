/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var account = "car"


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	var err error

	err = stub.PutState(account, []byte(strconv.Itoa(5000)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "spend" {
		var err error

		if len(args) < 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting 1 - payment sum")
		}

		amount, err := strconv.Atoi(args[0])

		accountValue, err := stub.GetState(account)
		accountInt, err := strconv.Atoi(string(accountValue[:]))

		if err != nil {
			return nil, err
		}

		err = stub.PutState(account, []byte(strconv.Itoa(accountInt - amount)))

		if err != nil {
			return nil, err
		}
		return nil, nil
	} else if function == "earn" {
		var err error

		if len(args) < 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting 1 - payment sum")
		}

		amount, err := strconv.Atoi(args[0])

		accountValue, err := stub.GetState(account)
		accountInt, err := strconv.Atoi(string(accountValue[:]))

		if err != nil {
			return nil, err
		}

		err = stub.PutState(account, []byte(strconv.Itoa(accountInt + amount)))

		if err != nil {
			return nil, err
		}
		return nil, nil
	} 
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	var jsonResp string

	// Handle different functions
	if function == "balance" {			
		valAsbytes, err := stub.GetState(account)
		if err != nil {
			jsonResp = "{\"Error\":\"Failed to get state for " + account + "\"}"
			return nil, errors.New(jsonResp)
		}
		return valAsbytes, nil
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

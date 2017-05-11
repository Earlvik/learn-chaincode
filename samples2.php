<?php
// QUERY

include('vendor/rmccue/requests/library/Requests.php');
Requests::register_autoloader();
$headers = array();
$data = '{
     "jsonrpc": "2.0",
     "method": "query",
     "params": {
         "type": 1,
         "chaincodeID": {
             "name": "61c985e701312b259d16240f82fc812fc90236b520e35705ba542ec04cc5d7e70076e83412f9e9bbba6a23d58584ee0098be8be7560127e3a21df84c88335468"
         },
         "ctorMsg": {
             "function": "read",
             "args": [
                 "viktor"
             ]
         },
         "secureContext": "admin"
     },
     "id": 2
 }';
$response = Requests::post('https://bf2ecd302ef6404abcf3ad797a0eefaa-vp0.us.blockchain.ibm.com:5002/chaincode', $headers, $data);

// INVOKE

include('vendor/rmccue/requests/library/Requests.php');
Requests::register_autoloader();
$headers = array();
$data = '{
     "jsonrpc": "2.0",
     "method": "invoke",
     "params": {
         "type": 1,
         "chaincodeID": {
             "name": "61c985e701312b259d16240f82fc812fc90236b520e35705ba542ec04cc5d7e70076e83412f9e9bbba6a23d58584ee0098be8be7560127e3a21df84c88335468"
         },
         "ctorMsg": {
             "function": "write",
             "args": [
                 "viktor",
                 "furqan",
                 "10"
             ]
         },
         "secureContext": "admin"
     },
     "id": 2
 }';
$response = Requests::post('https://bf2ecd302ef6404abcf3ad797a0eefaa-vp0.us.blockchain.ibm.com:5002/chaincode', $headers, $data);
# It mainly includes wallet creation and viewing functions

## Compile command
`
go mod tidy
`

`go build
`

## common

`--node  tcp://x.x.x.x:26657  default: tcp://127.0.0.1:26657   You can select another gateway node
`

`--rawUrl http://x.x.x.x:8545 default http://127.0.0.1:8545  Ethereum transaction sending address
`

## method

### 1.create wallet
`cli create`

### 2.View mnemonics
`cli import --seed "absurd mixture rural brick snow stumble gadget walk noodle pause like run"`


### 3.View privateKey
`cli import --pk 12fc22d015c4f7e015eb15668e4c35a6303a987b0cba53c00d5c2c357443c94f`

### 4.transfer
`cli transfer 0B9A99EDA4B4C12323C4C1EB1210E98E8FE22FE1AE3543CACD044825CC9A793B  dex10436g09cfa2nz2k6yphzj9qs9u93qtvvkkftur  dex1va8aaeystat4twpy70ns7235pxwczg0698twv4  100dst,100fm`
 send result : 
 `{"txhash":"E0171C03D0E598E1B0FD7A6562F388F1803E5DEABE35B312759BDE63B01900F1","raw_log":"[]","logs":[],"events":null}`

 transfer is an asynchronous method,transaction status needs to be queried using a return hash


### 5.query cosmos tx
`cli tx E0171C03D0E598E1B0FD7A6562F388F1803E5DEABE35B312759BDE63B01900F1 --node tcp://x.x.x.x:26657`

tx_result --> code  The value of code is used to determine 

success or failure  code 0:success other:fail  

tx query result :  

`{"hash": "E0171C03D0E598E1B0FD7A6562F388F1803E5DEABE35B312759BDE63B01900F1",
"height": 10768,
"index": 0,
"tx_result": {
"code": 0,
"data": "Ch4KHC9jb3Ntb3MuYmFuay52MWJldGExLk1zZ1NlbmQ=",
"log": "[]",
"info": "",
"gas_wanted": "200000",
"gas_used": "119519",
"events": [],
"codespace": ""
},
"tx": "CrcBCrQBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEpMBCipkZXgxOTN0anJjcHZndHNxdHgzbnowcm15OGEyejh1ZW0yZ2dxeWM0NTcSKmRleDEwNDM2ZzA5Y2ZhMm56Mms2eXBoemo5cXM5dTkzcXR2dmtrZnR1chocCgNhdHQSFTEwMDAwMDAwMDAwMDAwMDAwMDAwMBobCgJmbRIVMTAwMDAwMDAwMDAwMDAwMDAwMDAwEnwKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQMOk0G719PZjinNTYe7QQyS1ToPWfVoA1gPBDpSvNl4DBIECgIIARg4Eh8KGQoDYXR0EhIxMDAwMDAwMDAwMDAwMDAwMDAQwJoMGkHhYrHqaeYe2rLyzpmPQYSTDB5w+BNo9DmWXSEtsJ4Qy12kpy1xplS0s+A+yu8Ie06+e6OZKK0hcQnwGTKOXrxgAQ==",
"proof": {
"root_hash": "98C5770CE17469997DC14B5E9899668E087F8FFAB8A8ECFA2FB25A3E57894502",
"data": "CrcBCrQBChwvY29zbW9zLmJhbmsudjFiZXRhMS5Nc2dTZW5kEpMBCipkZXgxOTN0anJjcHZndHNxdHgzbnowcm15OGEyejh1ZW0yZ2dxeWM0NTcSKmRleDEwNDM2ZzA5Y2ZhMm56Mms2eXBoemo5cXM5dTkzcXR2dmtrZnR1chocCgNhdHQSFTEwMDAwMDAwMDAwMDAwMDAwMDAwMBobCgJmbRIVMTAwMDAwMDAwMDAwMDAwMDAwMDAwEnwKWQpPCigvZXRoZXJtaW50LmNyeXB0by52MS5ldGhzZWNwMjU2azEuUHViS2V5EiMKIQMOk0G719PZjinNTYe7QQyS1ToPWfVoA1gPBDpSvNl4DBIECgIIARg4Eh8KGQoDYXR0EhIxMDAwMDAwMDAwMDAwMDAwMDAQwJoMGkHhYrHqaeYe2rLyzpmPQYSTDB5w+BNo9DmWXSEtsJ4Qy12kpy1xplS0s+A+yu8Ie06+e6OZKK0hcQnwGTKOXrxgAQ==",
"proof": {
"total": 1,
"index": 0,
"leaf_hash": "mMV3DOF0aZl9wUtemJlmjgh/j/q4qOz6L7JaPleJRQI=",
"aunts": null
}
}
}`

### 6.query cosmos balance
`cli balance  dex10436g09cfa2nz2k6yphzj9qs9u93qtvvkkftur`

common : --denom string   Query specified currency balance 

query balance result : 
`[{"denom":"dst","amount":"100000000000000000000"},{"denom":"fm","amount":"100000000000000000000"}]`


### 7.eth transaction  Send Ethereum contract transactions
`cli eth transaction 0x2c5721e02c42e0059A3313c7b21FAA11f99DA908 0x0FBC4C943A4492f546d403FEf036D30B5331396d 6057361d0000000000000000000000000000000000000000000000000000000000000096 D326C9DFDF8D8BCF10FBCD9EB4B605C9D2AE1CBF59E5B613F3599758E7441959`

send result :

`{"hash":"0x31eeb9f79d49ee81763ac4e239aee7ed0c328145f4d09f9bd2906b4f23601e14"}`


### 8.eth hash  Ethereum contract transaction status query

`cli eth hash 0x31eeb9f79d49ee81763ac4e239aee7ed0c328145f4d09f9bd2906b4f23601e14`

the value of status is used to determine success or failure  code 1:success  0:fail  

query result eg: 

`{
"root": "0x",
"status": "0x1",
"cumulativeGasUsed": "0x19a28",
"logsBloom": "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
"logs": [],
"transactionHash": "0x31eeb9f79d49ee81763ac4e239aee7ed0c328145f4d09f9bd2906b4f23601e14",
"contractAddress": "0x0000000000000000000000000000000000000000",
"gasUsed": "0x19a28",
"blockHash": "0xbb6518512eb0182f49f03779f5d2d378717196fcbb79e49975e708ef05135695",
"blockNumber": "0x2f10",
"transactionIndex": "0x0"
}`

### 9.eth query  Ethereum contract query method

`cli eth query 0x2c5721e02c42e0059A3313c7b21FAA11f99DA908 0x0FBC4C943A4492f546d403FEf036D30B5331396d 2e64cec1`

The returned content needs to be self formatted

query restult eg: 

`0x0000000000000000000000000000000000000000000000000000000000000096`








